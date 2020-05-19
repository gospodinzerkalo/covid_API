package main


import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
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

	router.Methods("GET").Path("/allcases").HandlerFunc(api.GetAllCases())
	router.Methods("GET").Path("/country/{country}").HandlerFunc(api.GetByCountry("country"))
	router.Methods("GET").Path("/countries").HandlerFunc(api.GetCountries())
	router.Methods("GET").Path("/updates").HandlerFunc(api.GetUpdates())

	http.ListenAndServe("0.0.0.0:8000", router)
	return nil
}



