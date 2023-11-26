# go-rest
One of the ways to create a REST server using the Golang standard library.

## API methods
### GET /users
Get up to 10 oldest users.

### GET /user?id=\<id\>
Get information about the user by his \<id\>.

### POST /users
Send a request to create a new user with information from the request body. The request body must be in the form:
```
{
    "name": "YOURNAME",
    "surname": "YOURSURNAME",
    "age": YOURAGE
}
```

## Technologies
* Golang 1.21.1
* MySQL + [Driver](https://github.com/go-sql-driver/mysql)

## ❓ How to start
To start the server enter the following command into the console:
```
go run ./main.go
```
and then go to the address __localhost:4000__ (or __127.0.0.1:4000__).
