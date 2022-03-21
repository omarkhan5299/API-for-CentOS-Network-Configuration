# REST API-for-CentOS-Network-Configuratio
A REST API developed using GoLang (gorilla/mux) that allows users to configure the network settings for CentOS.

By: **Mohammed Omar Khan**

Time spent: **10** hours spent in total

The following **required** functionality is complete:

* [x] User can display the existing network configuration.
* [x] User can create a new network configartion by providing the necessary fields.
* [x] User can delete specific fields of the configuration.
* [x] User can modify/update the exisiting configuration.

## Installation & Run

Gorilla/mux is required to run the API. Run the command below to install the package.
```
go get -u github.com/gorilla/mux
```
To run the code, navigate to go/src/first_tutorial and execute main.go
```
cd go/src/first_tutorial
```
```
go run main.go
```
## Structure 
```
├── go
│   ├── bin
│   ├── pkg          
│   ├── src          
│       └── first_tutorial
│            └── main.go  
│        └── github.com
│        └── mux
│        └── go.mod
│        └── go.sum

```

## API

### /read
- `GET` : Get the existing network configuration

### /create
- `POST` : Create a brand new network configuration

### /delete/:field
- `DELETE` : Delete a specific field of the existing configuration

### /update/:field
- `POST` : Update a specific field in the existing netowkr configuration
