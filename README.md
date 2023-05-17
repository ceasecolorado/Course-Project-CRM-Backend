# Udacity GoLang Course Project: CRM Backend

This repo contains my submission to the final Project of Udacity Go Nanodegree Program. A CRM backend where users can make their own HTTP requests to perform CRUD operations.

## Dependencies

- go version go1.20.4
- github.com/gorilla/mux
- golang.org/x/exp/slices

# Install

- Install Golang by [visitig the offical instructions](https://golang.org/doc/install#testing)

- Instal gorilla/mux

```sh
go get -u github.com/gorilla/mux
```

- Install golang experiments

```sh
go get -u golang.org/x/exp/slices
```

## Launch

Run the manin.go file

```sh
go run main.go
```

## Usage

The script launches a server on the localhost port 3000. The application handles the following API calls:

- Getting a single customer: GET /customers/{id}
- Getting all customers in the database: GET /customers
- Creating a New customer: POST /customers
- Updating a customer: PUT /customers/{id}
- Delete a customer: DELETE /customers/{id}

## Authors

Cesar Colorado

## License

MIT
