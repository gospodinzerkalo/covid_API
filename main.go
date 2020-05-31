package main


import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"fmt"
	"./api"
)


func main() {

	app := cli.NewApp()
	app.Commands = cli.Commands{
		&cli.Command{
			Name:   "start",
			Usage:  "start the local server",
			Action: StartServer,
		},
	}
	app.Run(os.Args)

}

func StartServer(d *cli.Context) error {
	router := mux.NewRouter()

	//unit endpoints
	router.Methods("GET").Path("/allcases").HandlerFunc(api.GetAllCases())
	router.Methods("GET").Path("/country/{country}").HandlerFunc(api.GetByCountry("country"))
	router.Methods("GET").Path("/countries").HandlerFunc(api.GetCountries())
	router.Methods("GET").Path("/updates/today").HandlerFunc(api.GetUpdatesToday())
	router.Methods("GET").Path("/updates/all").HandlerFunc(api.GetUpdatesAll())

	//Kazakhstan endpoints
	router.Methods("GET").Path("/kz/allcases").HandlerFunc(api.GetAllCasesKazakhstan())

	http.ListenAndServe(GetPort(), router)
	return nil
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




