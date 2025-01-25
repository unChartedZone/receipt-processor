# Receipt Processor

Webservice that can consume a receipt and calculate points for a given receipt.

## Starting up the API

If you have `go` installed on your machine, please do the following steps: 
1. Run `go mod download` to install all dependencies. 
2. Run `go run .`, the api should be listening on `http://localhost:8080`.

## Running API with Docker

To run the API with docker, first please build the docker image with the following command.

```sh
docker build -t receipt-processor .
```

Then to run the image in a container, use the following command, the API should be available on `http://localhost:8080`

```sh
docker run -p 8080:8080 receipt-processor
```

I have also included a docker compose file to easily startup the API with the following command

```
docker compose up --build
```
