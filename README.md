# Golang

This is a simple golang rest API project with complete CRUD functionality.
Database will not be used in this project, however the CRUD operations will be performed on hard coded data using the local memory.

# Setup
Run the following command first incase you are not able to build your project for the first time, then later you can skip

```
# Allow the build to pass
go env -w GO111MODULE=auto
```

Run the following command to install the dependencies for the project

```
# Install mux router
go get -u github.com/gorilla/mux
```

To run the backend you are supposed to run the generated build exe file 
In this case since my build exe file is named as `golang-cc.exe`

Note that the exe file is generated with the name of the project folder name. 

Open the terminal and run the following command
```
./golang-cc.exe
```

# Endpoints
## BASE URL is at localhost:8000 
### Get all books
```
GET /books
```

### Get a book by id
```
GET /books/{id}
```

### Create a book
```
POST /books
```

### Update a book
```
PUT /books/{id}
```

### Delete a book
```
DELETE /books/{id}
```

## Sample Body for POST and PUT
```
{
  "isbn":"4545454",
  "title":"Updated Title",
  "author":{"firstname":"Harry",  "lastname":"White"}
}
```
