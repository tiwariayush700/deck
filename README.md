## Deck Backend

A Simple backend server for serving restful requests related to deck of playing cards which can be used for any type of
card games.

## How to run locally

### Dependencies

Data store which is required to be running locally is : Postgresql
> One can run the data store locally as container using the command `docker-compose up` or simply by running the makefile as `make data-store`<br/>

## How to run linter

- Install linter on system using `https://golangci-lint.run/usage/install/#local-installation`
- For running linter run : `make lint`

### Steps (run the commands from root folder)

1) Please make sure at least `go 1.16` is setup locally
2) Run `make set-local` to set local env values
3) Run `make build` to resolve lib dependencies
4) Run`go run main/*.go -file=local.json` and check the logs on console
5) Open the following url to check if the setup is up and running,
   `localhost:8000/ping`

### For running everything

Simply run `make` command. It exports env vars, sets up the data stores, runs linters and tests, and hosts your backend
server on `8000` port

N.B - `The server is also hosted on hiroku. One can test it out from the http files provided`