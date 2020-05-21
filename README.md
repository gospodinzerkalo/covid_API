<h1>API covid19 scraping</h1> <hr>
<h2>Overwiev</h2><br>
<p>API information about Covid-19 current time. All information taken from www.worldometers.info/coronavirus/</p>
<hr>
<h2>Using</h2>
<p><b>GET</b> "/allcases" > total cases</p>
<p><b>GET</b> "/country/{country_name}" > country information</p>
<p><b>GET</b> "/country/countries" >all countries information</p> 

<p><b>GET</b> "/country/updates/today" >updates for current day</p> 

<p><b>GET</b> "/updates/all" >updates for week</p> <hr>
<h2>Installation of project (Ubuntu)</h2>
<h3>Clone the Project </h3>
<code>git clone github.com/gospodinzerkalo/covid_API</code>
<h3>Install the dependencies</h3>
<code>make depends</code>
<h3>Build and Run</h3>
<code>make build</code> <br>
<code>make run</code>
