package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"kasir-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(name string) ([]models.Product, error) {
	query := `SELECT 
				p.id, p.name, p.price, p.stock, p.category_id, 
				c.id as cat_id, c.name as cat_name, c.description as cat_description
			FROM products p INNER JOIN
			categories c ON p.category_id = c.id `

	args := []interface{}{}
	if name != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+name+"%")
	}

	query += " ORDER BY p.id"

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		var catID sql.NullInt64
		var catName sql.NullString
		var catDesc sql.NullString

		err := rows.Scan(
			&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID,
			&catID, &catName, &catDesc,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error %w", err)
		}

		if catID.Valid {
			categoryID := int(catID.Int64)
			p.CategoryID = &categoryID
			p.Category = &models.Category{
				ID:          int(catID.Int64),
				Name:        catName.String,
				Description: catDesc.String,
			}
		} else {
			continue
		}

		products = append(products, p)

	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error %w", err)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	if product.CategoryID == nil {
		return errors.New("category_id is required")
	}

	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, *product.CategoryID).Scan(&product.ID)

	if err != nil {
		return fmt.Errorf("create error %w", err)
	}

	return nil

}

func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `SELECT
				p.id, p.name, p.price, p.stock, p.category_id,
				c.id as cat_id, c.name as cat_name, c.description as cat_description
			FROM products p INNER JOIN
			categories c ON p.category_id = c.id
			WHERE p.id = $1 `

	var p models.Product
	var catID sql.NullInt64
	var catName sql.NullString
	var catDesc sql.NullString
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID,
		&catID, &catName, &catDesc,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan atau tidak memiliki kategori")
	}

	if err != nil {
		return nil, fmt.Errorf("database error %w", err)
	}

	if !catID.Valid {
		return nil, errors.New("Kategori tidak valid")
	}

	categoryID := int(catID.Int64)
	p.CategoryID = &categoryID
	p.Category = &models.Category{
		ID:          int(catID.Int64),
		Name:        catName.String,
		Description: catDesc.String,
	}

	return &p, nil

}

func (repo *ProductRepository) Update(product *models.Product) error {
	if product.CategoryID == nil {
		return errors.New("category_id is required")
	}

	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, *product.CategoryID, product.ID)
	if err != nil {
		return fmt.Errorf("update error %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected error: %w", err)
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete error %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected error: %w", err)
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
