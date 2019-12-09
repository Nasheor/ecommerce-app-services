package main 

import(
	"net/http"
	"os"
	"log"
	logger "github.com/go-kit/kit/log"
	"github.com/rs/cors"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := logger.NewLogfmtLogger(os.Stderr)

	var cs CheckoutService
	cs = checkoutService{}

	mux := http.NewServeMux()

	mux.Handle("/checkout", httptransport.NewServer (
		makeCheckoutEndpoint(cs),
		decodeCheckoutRequest,
		encodeResponse,
	))

	handler := cors.Default().Handler(mux)
	log.Println("Checkout Microservice Running")
	logger.Log("msg", "HTTP", "addr", ":8082")
	logger.Log("err", http.ListenAndServe(":8082", handler))
}