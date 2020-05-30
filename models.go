package main

// AllCases ...
type AllCases struct {
	Cases 		string 	`json:"cases"`
	Deaths		string	`json:"deaths"`
	Recovered	string 	`json:"recovered"`
	ActiveCases	string	`json:"active_cases"`
	Critical	string	`json:"critical"`
}


// Country ...
type Country struct {
	Name		string	`json:"name"`
	Place		string 	`json:"place"`
	Cases		string 	`json:"cases"`
	NewCases	string	`json:"new_cases"`
	Deaths		string 	`json:"deaths"`
	NewDeaths	string	`json:"new_deaths"`
	Recovered	string	`json:"recovered"`
	ActiveCases	string	`json:"active_cases"`
	Critical	string	`json:"critical"`
}

type Updates struct {
	Results	[]string	`json:"results,omitempty"`
}

type UpdatesAll struct {
	Day		string		`json:"day"`
	Results	[]string 	`json:"results"`
}

type KazakhstanCases struct {
	Name 		string	`json:"name"`
	Cases		string	`json:"cases"`
	Recovered	string	`json:"recovered"`
	Deaths		string	`json:"deaths"`
}

