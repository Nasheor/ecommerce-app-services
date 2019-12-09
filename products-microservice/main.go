package main 

import(
	"net/http"
	"os"
	"log"
	"github.com/go-kit/kit/log"
	"github.com/rs/cors"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var ps ProductsService
	ps = productsService{}

	var ds DeleteService
	ds = deleteService{}

	mux := http.NewServeMux()

	mux.Handle("/products", httptransport.NewServer (
		makeProductsEndpoint(ps),
		decodeGetProductsRequest,
		encodeResponse,
	))

	mux.Handle("/delete", httptransport.NewServer (
		makeDeleteProductEndpoint(ds),
		decodeDeleteProductRequest,
		encodeResponse,
	))

	handler := cors.Default().Handler(mux)
	log.Println("Products Microservice Running")
	logger.Log("msg", "HTTP", "addr", ":8081")
	logger.Log("err", http.ListenAndServe(":8081", handler))
}