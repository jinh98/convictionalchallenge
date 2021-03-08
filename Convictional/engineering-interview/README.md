
# Challenge

  

This project provides a API to fullfil get product by ID

  

## Requirements

This project runs on a Linux environment and uses go 1.15.

  

## Building the API

 `git clone` into this respository and do `cd convictional`

 To build the project, run:

```bash

$ go build -o sample main.go

```

  

This command will build the project and create an executable`sample` for the sample program in `sample.go` to demonstrate basic uses of the library. 

  
  

## Using the sample program

```bash

$ ./sample

```

This command will run the executable for the sample program and start the server.

Only implemented one of the requirement APIs due to time constraints, to test it do 

```bash

$ curl localhost:443/product/{id}

```

The other requirements would be very similar to this manner...