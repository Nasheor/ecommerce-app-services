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
	mux := http.NewServeMux()
	
	mux.Handle("/login", httptransport.NewServer ( 
		makeLoginEndpoint(lvs),
		decodeLoginRequest,
		encodeResponse,
	))

	handler := cors.Default().Handler(mux)
	logger.Log("msg", "HTTP", "addr", ":8080" )
	logger.Log("err", http.ListenAndServe(":8080", handler))
}