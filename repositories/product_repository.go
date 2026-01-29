package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ProductRepository interface {
	GetAll(limit, offset int) ([]models.Product, int, error)
	GetByID(id int) (*models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(id int, product models.Product) (*models.Product, error)
	Delete(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAll(limit, offset int) ([]models.Product, int, error) {
	var total int
	err := r.db.QueryRow("SELECT count(*) FROM products").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(`
		SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name
		FROM products p
		JOIN categories c ON p.category_id = c.id
		ORDER BY p.id LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.CategoryName); err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}
	return products, total, nil
}

func (r *productRepository) GetByID(id int) (*models.Product, error) {
	var p models.Product
	err := r.db.QueryRow(`
		SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1`, id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.CategoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id",
		product.Name, product.Price, product.Stock, product.CategoryID,
	).Scan(&id)
	if err != nil {
		return models.Product{}, err
	}

	createdProduct, err := r.GetByID(id)
	if err != nil {
		return models.Product{}, err
	}
	if createdProduct == nil {
		return models.Product{}, sql.ErrNoRows
	}

	return *createdProduct, nil
}

func (r *productRepository) Update(id int, product models.Product) (*models.Product, error) {
	res, err := r.db.Exec("UPDATE products SET name=$1, price=$2, stock=$3, category_id=$4 WHERE id=$5",
		product.Name, product.Price, product.Stock, product.CategoryID, id)
	if err != nil {
		return nil, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}

	return r.GetByID(id)
}

func (r *productRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}
