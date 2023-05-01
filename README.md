## API Documentation
Swagger documentation can be accessed from
```
baseURL/swagger/index.html
```
(e.g. localhost:8080/swagger/index.html)
## Pre-requisites
This program were build on a machine with these specifications:
- Go 1.20
- MariaDB 10.4.6

## Data Initialization
Run this command to initialize database schema and initial data
```
go run cmd/init_data/main.go
```
## Running the Program
To install all the required packages
```
go get -d ./...
```
Run the program with this command
```
go run main.go
```
Build the program
```
go build -o executableName main.go
```
## .Env File (example)
```
SECRET_BYTES=golangstandardapi
ACCESS_TOKEN_TTL=4320m
PORT=7007
DATABASE_URI=root:@tcp(localhost:3306)/golang_test?multiStatements=true
PUBLIC_KEY=./keys/publickey.crt
PRIVATE_KEY=./keys/private.key
```
## Unit Tests
Run unit tests with this commmand
```
go test -v -race -short ./...

```