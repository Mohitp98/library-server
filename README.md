# library-server
Rest API service for library

## Installation

### Step 1: Clone the Repo

```sh
$ git clone git@github.com:Mohitp98/library-server.git 
$ cd library-server
```

### Step 2: Install MongoDB using Docker

```sh
# default database port is 27017
$ docker run --name mongo-database -p 27017:27017 -d mongo:latest
```

**Note:- Default port number for the service: `5000`**

### Step 3: Run application

```sh
# build application binary and execute.
$ go build -o app .
$ ./app
```

## API Documentation
```sh
# Add new book: (/books)
curl --location --request POST 'localhost:5000/books' \
--header 'Content-Type: application/json' \
--data-raw '{
	"name": "leaders eat last",
	"author": "simon sinek",
	"pages": 450,
	"price": 599,
	"domain": "personal-growth",
	"language": "english"
}'

# Get All Books: (/books)
$ curl --location --request GET 'localhost:5000/books'

# Get Book Details: (/book/{book_id})
$ curl --location --request GET 'localhost:5000/book/620d3374006fa347af37db7b'

```

## Unit Testing

To run unit test cases for Service. Firstly clone this repo and run following commands:

```sh
# install dependencies
$ go test -v .
```