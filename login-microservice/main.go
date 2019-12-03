package main
import(
	"net/http"
	"os"
	"github.com/go-kit/kit/log"
	"github.com/rs/cors"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var lvs LoginService
	lvs = loginService{}

	// Creating a http route multiplexer for handling cors
	mux := http.NewServeMux()
	
	// Wrapping the multiplexer aroung transport/http to expose the service to the outside world
	mux.Handle("/login", httptransport.NewServer ( 
		makeLoginEndpoint(lvs),
		decodeLoginRequest,
		encodeResponse,
	))
	
	// Enabling CORS with default configuration POST and GET
	handler := cors.Default().Handler(mux)
	
	// Starting up our service 
	logger.Log("msg", "HTTP", "addr", ":8080" )
	logger.Log("err", http.ListenAndServe(":8080", handler))
}