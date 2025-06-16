package products

import (
	"github.com/dmytrovorona/monolith-shop/controllers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", controllers.ListProducts)
	mux.HandleFunc("/products/", controllers.CreateProduct)

	log.Println("Product service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
