# Gitline

Gitline takes a username and displays user github activity across time

### Major dependencies

- [Go](https://golang.org/) (>= 1.16)
- [docker-compose](https://docs.docker.com/compose/install/) (>= 1.27.4)

#### To install the dependencies locally, run the following commands:
- `yarn install`
- `go mod download`
-------------------------------------------------------------------

## Local development

The application ran through Docker can be accessed on `localhost:3001/`.

To enable debugging and hot-reloading of Go files:

`docker-compose -f docker/docker-compose.dev.yml up --build --force-recreate`

Hot reloading is provided by Air, so any changes to the Go code (including templates)
will rebuild and restart the application without requiring manually stopping and restarting the compose stack.

### Without docker

Alternatively to set it up not using Docker use below (make sure you are in the root of the directory). 

- `yarn install && yarn build `
- `go build main.go `
- `./main `

This hosts it on `localhost:1235`

  -------------------------------------------------------------------

## Run Cypress tests

`docker-compose -f docker/docker-compose.dev.yml up -d --build --force-recreate`

`yarn && yarn cypress`
    
-------------------------------------------------------------------

## Run Unit/Integration tests

To run unit/integration tests locally, run `npm run test-handlers` and `npm run test-controllers`.
 
-------------------------------------------------------------------