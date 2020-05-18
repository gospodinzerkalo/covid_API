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
		resCountry := &Country{}
		c.OnHTML("#main_table_countries_today tbody", func(element *colly.HTMLElement) {
			element.ForEach("tr", func(_ int, el*colly.HTMLElement)  {
				s := el.ChildText(".mt_a")
				if s==country{
					c := el.ChildTexts("td")
					resCountry.Name = s
					resCountry.Place = c[0]
					resCountry.Cases = c[2]
					resCountry.NewCases = c[3]
					resCountry.Deaths = c[4]
					resCountry.NewDeaths = c[5]
					resCountry.Recovered =c[6]
					resCountry.ActiveCases = c[7]
					resCountry.Critical = c[8]
				}

			})
		})
		c.Visit(baseURL)
		data, err := json.Marshal(resCountry)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, []byte("Error: "+err.Error()))
			return
		}
		writeResponse(w, http.StatusOK, data)

	}
}
func GetCountries() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		countries := []*Country{}
		c := colly.NewCollector()
		c.OnHTML("#main_table_countries_today tbody", func(element *colly.HTMLElement) {
			element.ForEach("tr", func(_ int, el*colly.HTMLElement)  {

				resCountry := &Country{}
				c := el.ChildTexts("td")
				resCountry.Name = c[1]
				resCountry.Place = c[0]
				resCountry.Cases = c[2]
				resCountry.NewCases = c[3]
				resCountry.Deaths = c[4]
				resCountry.NewDeaths = c[5]
				resCountry.Recovered =c[6]
				resCountry.ActiveCases = c[7]
				resCountry.Critical = c[8]
				countries = append(countries,resCountry)

			})
		})
		c.Visit(baseURL)
		data, err := json.Marshal(countries)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, []byte("Error: "+err.Error()))
			return
		}
		writeResponse(w, http.StatusOK, data)
	}
}


func writeResponse(w http.ResponseWriter, status int, msg []byte) {
	w.WriteHeader(status)
	w.Write(msg)
}
