package main


import (
	"net/http"
	"os"
	"github.com/gorilla/mux"

	"fmt"
)


func main() {

	router := mux.NewRouter()

	//unit endpoints
	router.Methods("GET").Path("/allcases").HandlerFunc(GetAllCases())
	router.Methods("GET").Path("/country/{country}").HandlerFunc(GetByCountry("country"))
	router.Methods("GET").Path("/countries").HandlerFunc(GetCountries())
	router.Methods("GET").Path("/updates/today").HandlerFunc(GetUpdatesToday())
	router.Methods("GET").Path("/updates/all").HandlerFunc(GetUpdatesAll())

	//Kazakhstan endpoints
	router.Methods("GET").Path("/kz/allcases").HandlerFunc(GetAllCasesKazakhstan())

	http.ListenAndServe(GetPort(), router)


}



func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "5000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}




