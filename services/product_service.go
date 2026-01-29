package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"kasir-api/utils"
)

type ProductService interface {
	GetAllProducts(page, pageSize int) ([]models.Product, *utils.PaginationMeta, error)
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(id int, product models.Product) (*models.Product, error)
	DeleteProduct(id int) error
}

type productService struct {
	repository repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repository: repo}
}

func (s *productService) GetAllProducts(page, pageSize int) ([]models.Product, *utils.PaginationMeta, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	products, total, err := s.repository.GetAll(pageSize, offset)
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

	return products, meta, nil
}

func (s *productService) GetProductByID(id int) (*models.Product, error) {
	return s.repository.GetByID(id)
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	return s.repository.Create(product)
}

func (s *productService) UpdateProduct(id int, product models.Product) (*models.Product, error) {
	return s.repository.Update(id, product)
}

func (s *productService) DeleteProduct(id int) error {
	return s.repository.Delete(id)
}
