package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"github.com/gospodinzerkalo/covid_API/api"
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
	router := fasthttprouter.New()

	//endpoints...
	router.GET("/allcases",api.GetAllCases())
	router.GET("/kz/allcases",api.GetAllCasesKazakhstan())
	router.GET("/country/:country",api.GetByCountry())
	router.GET("/countries",api.GetCountries())
	router.GET("/updates/today",api.GetUpdatesToday())
	router.GET("/updates/all",api.GetUpdatesAll())

	log.Fatal(fasthttp.ListenAndServe(GetPort(), router.Handler))
	return nil
}
func GetPort() string{
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}




