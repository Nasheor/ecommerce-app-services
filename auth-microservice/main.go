package main

import (
   "encoding/json"
   "github.com/google/uuid"
   "gopkg.in/oauth2.v3/models"
   "log"
   "net/http"

   "gopkg.in/oauth2.v3/errors"
   "gopkg.in/oauth2.v3/manage"
   "gopkg.in/oauth2.v3/server"
   "gopkg.in/oauth2.v3/store"
)
type Product struct {
	I string `json:"I"`
	Q int64 `json:"Q"`
   P float64 `json:"P"`
   N string `json:"name"`
}

func main() {
   manager := manage.NewDefaultManager()
   manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

   // Creating OAUTH2.v3 memory store for access tokens
   manager.MustTokenStorage(store.NewMemoryTokenStore())

   // Creating OAUTH2.v3 memory store for client credentials
   clientStore := store.NewClientStore()
   
   manager.MapClientStorage(clientStore)

   srv := server.NewDefaultServer(manager)
   srv.SetAllowGetAccessRequest(true)
   srv.SetClientInfoHandler(server.ClientFormHandler)
   manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

   srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
      log.Println("Internal Error:", err.Error())
      return
   })

   srv.SetResponseErrorHandler(func(re *errors.Response) {
      log.Println("Response Error:", re.Error.Error())
   })

   http.HandleFunc("/generate_token", func(w http.ResponseWriter, r *http.Request) {
      srv.HandleTokenRequest(w, r)
   })

   // Generates client credentials for the user's 
   // request from a third party app
   http.HandleFunc("/get_client_credentials", func(w http.ResponseWriter, r *http.Request) {
      clientId := uuid.New().String()[:8]
      clientSecret := uuid.New().String()[:8]
      err := clientStore.Set(clientId, &models.Client{
         ID:     clientId,
         Secret: clientSecret,
         Domain: "http://localhost",
      })
      if err != nil {
         log.Println(err.Error())
      }

      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
   })

   // The address of the resource requested from the domain with 
   // access token 
   http.HandleFunc("/products", validateToken(func(w http.ResponseWriter, r *http.Request) {
         res, err := http.Get("http://localhost:8081/products")
         if err != nil {
            log.Println(err)
         } else {
            products := make([]*Product, 0)
            if err := json.NewDecoder(res.Body).Decode(&products); err != nil {
               log.Println(err)
            }
            w.Header().Set("Content-Type", "application/json") 
            json.NewEncoder(w).Encode(&products)
      }
   }, srv))
   log.Println("Authorization Server Running.")
   log.Fatal(http.ListenAndServe(":9096", nil))
   log.Println("Authorization Server Running.")
}


//  Validates the access token previously generated for the user from the
//  requesting domain using the client credentials
func validateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      _, err := srv.ValidationBearerToken(r)
      if err != nil {
         http.Error(w, err.Error(), http.StatusBadRequest)
         return
      }

      f.ServeHTTP(w, r)
   })
}