package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type CategoryRepository interface {
	GetAll(limit, offset int) ([]models.Category, int, error)
	GetByID(id int) (*models.Category, error)
	Create(category models.Category) (models.Category, error)
	Update(id int, category models.Category) (*models.Category, error)
	Delete(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAll(limit, offset int) ([]models.Category, int, error) {
	var total int
	err := r.db.QueryRow("SELECT count(*) FROM categories").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query("SELECT id, name FROM categories ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, 0, err
		}
		categories = append(categories, c)
	}
	return categories, total, nil
}

func (r *categoryRepository) GetByID(id int) (*models.Category, error) {
	var c models.Category
	err := r.db.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).
		Scan(&c.ID, &c.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *categoryRepository) Create(category models.Category) (models.Category, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO categories (name) VALUES ($1) RETURNING id",
		category.Name,
	).Scan(&id)
	if err != nil {
		return models.Category{}, err
	}
	category.ID = id
	return category, nil
}

func (r *categoryRepository) Update(id int, category models.Category) (*models.Category, error) {
	res, err := r.db.Exec("UPDATE categories SET name=$1 WHERE id=$2", category.Name, id)
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
	category.ID = id
	return &category, nil
}

func (r *categoryRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id=$1", id)
	return err
}
