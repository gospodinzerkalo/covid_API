package main


import (
	"net/http"
	"os"
	//"github.com/gocolly/colly"
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

	http.ListenAndServe("0.0.0.0:8000", router)
	return nil
}





//config := &tls.Config{
//	InsecureSkipVerify: true,
//}
//transport := &http.Transport{
//	TLSClientConfig:config,
//}
//client := &http.Client{
//	Transport: transport,
//}
//
//response,err := client.Get(baseURL)
//if err!=nil {
//	fmt.Println(err.Error())
//	os.Exit(1)
//}
//body,err := ioutil.ReadAll(response.Body)
//
//if err!=nil{
//	fmt.Println(err.Error())
//	os.Exit(1)
//}
//fmt.Println(string(body))

//doc, err := goquery.NewDocument(baseURL)
//if err != nil {
//	log.Fatal(err)
//}
//
//doc.Find(".maincounter-number").Each(func(i int, s *goquery.Selection) {
//	ss,_ := s.Html()
//	fmt.Printf("Review %d: %s\n", i, ss)
//})