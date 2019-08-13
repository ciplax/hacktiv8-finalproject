# hacktiv8-finalproject
Final project of hacktive8 golang class
This app is far from perfect.

## Requirements
- [GOlang](https://golang.org/dl/) (App were written and tested using go version go1.12.6 windows/amd64)
- [Mariadb/Mysql](https://mariadb.org/download/) 

## Setup
- Import database using [hacktiv8.sql](./hacktiv8.sql) to any database/schema
- Adjust the property of [config.ini](./files/etc/config.ini) in accordance to the database/schema you used in previous step 
```
DBMaster = "dbuser:dbpassword@/dbname?charset=utf8"
DBSlave = "dbuser:dbpassword@/dbname?charset=utf8"
```

## Running
```
go run app.go
```
