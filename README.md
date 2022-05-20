# REST-API-CRUD
The repository deals with the development of rest api using go with the integration of database using gorm and router using mux

## Installation process

### For mux router
```
go get -u github.com/gorilla/mux
```
### For GORM Integration tools
```
go get -u gorm.io/gorm
```
### Integrating the mysql server
```
go get -u gorm.io/driver/mysql
```
### Setting the environment port
```
PORT=8000
```
#### Note that the port has to be declared inside an .env file

### Setting up the router
```
router.HandleFunc()
http.ListenAndServe()
```
## How to run
- Open postman
- Type
```
http://localhost:3000/(your endpoint)
```
- Send the request 
- Before sending the request run the following command in the project
```
go run main.go
```
