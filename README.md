<p align="center">
    <h1 align="center">US Zipcodes API</h1>
    <p align="center">An API server for fetching US zipcodes.</p>
    <br>
    <p align="center">
        <img src="./demos/render1726338999898.gif" alt="API server with MySQL database" width="800" height="500">
    </p>
</p>

## Description
A simple API to fetch the US zipcodes(2000 census) and their associated information in differet ways.

The goal for creating this project is for practicing, and this API is not intended for commercial use, only for educational purposes.

## Dependencies

To execute this API you will need to have installed the following software:
* `Golang` version `1.22.7` or higher
* `Docker`
* `Docker compose`

The minimum requirement will be a `Golang` version `1.22.7` or higher installed to run the API server with the build-in database.

## Usage

The API server can be executed using `make`or compiling the server  and executed it.

### Up and running the Server with Make
You can see the available `make` commands by executing `make help`

To run the API server you can execute one of the following commands:

* `make compose-localdb start-localdb`
* `make compose-mysqldb start-mysqldb`

For cleaning up the reources created execute `make clean`

### Compiling the API server

To compile the API server, you can execute one of the following commands:

* `go build -o <binary-name> ./cmd/zipcodes`
* `make compile`

The first one is for compile and get a binary specific for you current machine, the second one is for creating binaries for different OS and architectures(arm,386, amd64) in the the local `bin/` directory.

The population of the database is automated, downloading the US zipcodes from the CSV file created by [scpike](https://github.com/scpike) in the following repository [link](https://github.com/scpike/us-state-county-zip).

After the APi server is up and running, the following endpoints will be available at port `:20790`:

* `/zipcodes/{zipcode}`
* `/counties/{county}`
* `/states/{state}/counties/{county}`
* `/states/{state}/cities/{city}`
* `/counties/{county}/cities/{city}`

You can see how to fetch data form the server in the following demo:

<p align="center">
    <img src="./demos/render1726343642493.gif" alt="Fetching data with the provided endpoints" width="800" height="500">
</p>

These are the `curl` examples shown in the gift demo:

* `curl -i http://localhost:20790/zipcodes/90001`
* `curl -i http://localhost:20790/counties/Sampson`
* `curl -i http://localhost:20790/states/North%20Carolina/counties/Mecklenburg`
* `curl -i http://localhost:20790/states/Virginia/cities/Ivor`
* `curl -i http://localhost:20790/counties/Cook/cities/Rosemont`

### TO-DO

* [ ] Add documentation (Structuctures, Functions, Interfaces, etc)
* [ ] Add automated tests
* [ ] Normalize the datatabase
* [ ] Automate the port to use configuration
* [ ] Add MongoDB database connection
* [ ] Add GitHub actions to automate tasks
* [ ] Deploy the app to a local Kubernetes environment
* [ ] Deploy the service to a Kubernentes cloud environment