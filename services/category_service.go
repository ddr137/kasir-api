package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"kasir-api/utils"
)

type CategoryService interface {
	GetAllCategories(page, pageSize int) ([]models.Category, *utils.PaginationMeta, error)
	GetCategoryByID(id int) (*models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(id int, category models.Category) (*models.Category, error)
	DeleteCategory(id int) error
}

type categoryService struct {
	repository repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repository: repo}
}

func (s *categoryService) GetAllCategories(page, pageSize int) ([]models.Category, *utils.PaginationMeta, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	categories, total, err := s.repository.GetAll(pageSize, offset)
	if err != nil {
		return nil, nil, err
	}

	totalPage := 0
	if pageSize > 0 {
		totalPage = (total + pageSize - 1) / pageSize
	}

	meta := &utils.PaginationMeta{
		Page:      page,
		Total:     total,
		TotalPage: totalPage,
	}

	return categories, meta, nil
}

func (s *categoryService) GetCategoryByID(id int) (*models.Category, error) {
	return s.repository.GetByID(id)
}

func (s *categoryService) CreateCategory(category models.Category) (models.Category, error) {
	return s.repository.Create(category)
}

func (s *categoryService) UpdateCategory(id int, category models.Category) (*models.Category, error) {
	return s.repository.Update(id, category)
}

func (s *categoryService) DeleteCategory(id int) error {
	return s.repository.Delete(id)
}
