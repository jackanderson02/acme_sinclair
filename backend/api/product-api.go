package api



import (
	"acme/service"
	// "acme/model"
	"encoding/json"
	// "fmt"
	"net/http"
	// "strconv"
	// "io"
)

type ProductAPI struct {
    productService *service.ProductService
}

func NewProductAPI(productService *service.ProductService) *ProductAPI{
    return &ProductAPI{
        productService: productService,
    }
}

func (api *ProductAPI) GetProducts(writer http.ResponseWriter, _ *http.Request) {
	service := api.productService
	products, err := service.GetProductsService()
	json.NewEncoder(writer).Encode(products)

	if (err != nil){
		http.Error(writer, "Failed to retrieve products.", http.StatusInternalServerError)
	}
}