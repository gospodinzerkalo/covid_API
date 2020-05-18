package api

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"strings"
)

const (
	baseURL = "https://www.worldometers.info/coronavirus/"
)

func GetAllCases() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := colly.NewCollector()
		list := make([]string,0)

		c.OnHTML(".maincounter-number", func(element *colly.HTMLElement) {
			list = append(list,strings.TrimSpace(element.Text))
		})
		c.Visit(baseURL)
		fmt.Println()
		cases := AllCases{
			Cases:    	list[0],
			Deaths:    list[1],
			Recovered: list[2],
		}
		data, err := json.Marshal(cases)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, []byte("Error: "+err.Error()))
			return
		}
		writeResponse(w, http.StatusOK, data)

	}
}

func GetByCountry(countryParam string) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		c := colly.NewCollector()

		vars := mux.Vars(r)
		country, ok := vars[countryParam]
		if !ok {
			writeResponse(w, http.StatusBadRequest, []byte("Crime ID not found"))
			return
		}
		c.OnHTML("#main_table_countries_today", func(element *colly.HTMLElement) {
			fmt.Print(element.ChildText("tbody:contains(\"Kazakhstan\")"))
		})
		c.Visit(baseURL)
		fmt.Println(country)
	}
}





func writeResponse(w http.ResponseWriter, status int, msg []byte) {
	w.WriteHeader(status)
	w.Write(msg)
}
