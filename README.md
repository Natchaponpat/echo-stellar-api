# echo-stellar-api
Exercise for How-to-use [Echo framework](https://echo.labstack.com/) and [Stellar Go SDK](https://www.stellar.org/developers/go/reference/)
## Installation

1. Clone the project to your computer.
```
mkdir -p $(go env GOPATH)/src/github.com/Natchaponpat && cd $_
git clone https://github.com/Natchaponpat/echo-stellar-api.git
```
2. Download package dependencies. Go to project directory and run command depend on your Go version
```
go get github.com/labstack/echo
go get github.com/globalsign/mgo
```
3. Test the project
```
go build ./...
```
