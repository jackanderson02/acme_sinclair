package service

import(
	"acme/model"
	"acme/repository/user"
	"fmt"
	"errors"
)

type ProductService struct {
	repository user.Repository
}

func NewProductService(repo user.Repository) *ProductService{
	return &ProductService{
		repository: repo,
	}
}


func (service *ProductService) GetProductsService() ([]model.Product, error){
	repo := service.repository
	products, err := repo.GetProducts()

	if err != nil {
		fmt.Println("Error getting products from DB:", err)
		return nil, errors.New("There was an error getting the produts from the database.")
	}

	return products, nil

}