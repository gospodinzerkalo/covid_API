<h1>API covid19 scraping</h1> <hr>
<h2>Overwiev</h2><br>
<p>API information about Covid-19 current time. All information taken from www.worldometers.info/coronavirus/ and informburo.kz/</p>
<hr>

<h2>Dependencies</h2>
<p><a href="github.com/valyala/fasthttp">fasthttp</a> - http server, but 10x faster than net/http</p>
<p><a href="github.com/urfave/cli">cli</a> - cli command</p>
<p><a href="github.com/gocolly/colly">colly</a> - lib for scraping websites</p>
<p><a href="fasthttprouter">fasthttprouter</a> -  router</p>
<h2>Using</h2>
<p><b>[GET]</b> <a href="covidapigolang.herokuapp.com/allcases">covidapigolang.herokuapp.com/allcases</a> > total cases</p>
<p><b>[GET]</b> <a href="covidapigolang.herokuapp.com/country/{countryName}">covidapigolang.herokuapp.com/country/{countryName}</a> > country information</p>
<p><b>[GET]</b> <a href="covidapigolang.herokuapp.com/countries">covidapigolang.herokuapp.com/countries</a> > all countries information</p> 
<p><b>[GET]</b> <a href="covidapigolang.herokuapp.com/updates/today">covidapigolang.herokuapp.com/updates/today</a> > updates for current day</p> 
<p><b>[GET]</b> <a href="covidapigolang.herokuapp.com/updates/all">covidapigolang.herokuapp.com/updates/all</a> >all updates for week</p> 
<p><b>[GET]</b> <a href="covidapigolang.herokuapp.com/kz/allcases">covidapigolang.herokuapp.com/kz/allcases</a> > cases in Kazakhstan's cities/regions</p> <hr>

<h2>Installation of project (Ubuntu)</h2>
<h3>Clone the Project </h3>
<code>git clone github.com/gospodinzerkalo/covid_API</code>

<h3>Installation (with docker)</h3>
<code>docker-compose build</code><br>
<code>docker-compose up</code>

<h3>Install the dependencies (with Makefile)</h3>
<code>make depends</code>
<h3>Build and Run</h3>
<code>make build</code> <br>
<code>make run</code>
