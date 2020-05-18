build:
	go build main.go
run:
	./main start
depends:
	go get "github.com/gorilla/mux"
	go get "github.com/urfave/cli"
	go get "github.com/gocolly/colly"