# library-server
Rest API service for library

## Installation

### Step 1: Clone the Repo

```sh

$ git clone git@github.com:Mohitp98/library-server.git 

```

### Step 2: Install the dependencies

```sh
$ go mod install
```

**Note:- Default port number for the service: `5000`**

### Step 3:  Run application

```sh
$ go build . -o app
$ ./app
```

## Unit Testing

To run unit test cases for Service. Firstly clone this repo and run following commands:

```sh
# install dependencies
$ go test -v .
```