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
		cases := AllCases{}
		c.OnHTML(".content-inner", func(element *colly.HTMLElement) {
			list = append(list,strings.TrimSpace(element.Text))
			s1 := element.ChildTexts(".maincounter-number")
			s2 := element.ChildTexts(".number-table-main")
			cases.Cases = s1[0]
			cases.Deaths = s1[1]
			cases.Recovered = s1[2]
			cases.ActiveCases = s2[0]
			cases.Critical = s2[1]
		})

		c.Visit(baseURL)
		fmt.Println()
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

func GetUpdatesToday() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		c:= colly.NewCollector()
		list := make([]string,0)
		c.OnHTML("#news_block", func(element *colly.HTMLElement) {
			list =strings.Split(strings.TrimSpace(element.ChildTexts("div")[1]),"\n")


		})
		c.Visit(baseURL)
		updates := Updates{}
		for _,v := range list[1:] {
			if v!=""{
				updates.Results= append( updates.Results,strings.Replace(v, "[source]", "", 1))
			}
		}
		data, err := json.Marshal(updates)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, []byte("Error: "+err.Error()))
			return
		}
		writeResponse(w, http.StatusOK, data)
	}
}

func GetUpdatesAll() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		c := colly.NewCollector()
		updates := []UpdatesAll{}
		c.OnHTML(".newsdate_div", func(element *colly.HTMLElement) {
			updates = append(updates,UpdatesAll{Results: element.ChildTexts(".news_post")})

		})

		indD := 0
		c.OnHTML("[class='btn btn-light date-btn']", func(element *colly.HTMLElement) {
			updates[indD].Day = element.Text
			indD++
		})
		c.Visit(baseURL)
		data, err := json.Marshal(&updates)
		if err!=nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error!"))
		}
		writeResponse(w,http.StatusOK,data)
	}
}

func writeResponse(w http.ResponseWriter, status int, msg []byte) {
	w.WriteHeader(status)
	w.Write(msg)
}
