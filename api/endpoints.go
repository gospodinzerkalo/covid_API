package api

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/valyala/fasthttp"
	"net/http"
	"strings"
)

const (
	baseURL = "https://www.worldometers.info/coronavirus/"
	kzURL	= "https://informburo.kz/novosti/koronavirus-v-kazahstane-situaciya-na-21-maya-live.html"
)

//get total cases
func GetAllCases() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
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
			writeResponse(ctx, http.StatusInternalServerError, []byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx, http.StatusOK, data)

	}
}
// get case in special country
func GetByCountry() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Println("ss")
		c := colly.NewCollector()

		vars := ctx.UserValue("country")
		resCountry := &Country{}
		c.OnHTML("#main_table_countries_today tbody", func(element *colly.HTMLElement) {
			element.ForEach("tr", func(_ int, el*colly.HTMLElement)  {
				s := el.ChildText(".mt_a")
				if s==vars{
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
			writeResponse(ctx, http.StatusInternalServerError, []byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx, http.StatusOK, data)

	}
}

// get all countries with information about cases
func GetCountries() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
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
			writeResponse(ctx, http.StatusInternalServerError, []byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx, http.StatusOK, data)
	}
}

// get news and updates current day
func GetUpdatesToday() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
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
			writeResponse(ctx, http.StatusInternalServerError, []byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx, http.StatusOK, data)
	}
}

// get news for the week
func GetUpdatesAll() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
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
			writeResponse(ctx,http.StatusInternalServerError,[]byte("Error!"))
		}
		writeResponse(ctx,http.StatusOK,data)
	}
}


// get cases in Kazakhstan's city/region
func GetAllCasesKazakhstan() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		c := colly.NewCollector()
		cases := []KazakhstanCases{}
		c.OnHTML(".table", func(element *colly.HTMLElement) {
			cities := element.ChildTexts("tbody tr")
			for _,city := range cities[1:]{
				split := strings.Split(city,"\n")
				cas := KazakhstanCases{
					Name:      split[0],
					Cases:     split[1],
					Recovered: split[2],
					Deaths:    split[3],
				}
				cases = append(cases,cas)
			}
		})
		c.Visit(kzURL)
		data,err := json.Marshal(&cases)
		if err !=nil{
			writeResponse(ctx,http.StatusInternalServerError,[]byte("Error!"))
			return
		}
		writeResponse(ctx,http.StatusOK,data)
	}
}


func writeResponse(ctx *fasthttp.RequestCtx, status int, msg []byte) {
	ctx.SetStatusCode(status)
	ctx.Write(msg)
}