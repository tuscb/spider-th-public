# spider-th-public
Spider in Thailand

This github repository provides files and folders for Spider in Thailand project.

The following guide is for quicky testing insert `.json` file storing data of spider into MongoDB database and query the ith data spider item.


## Requirement
The `Go` programming language is installed on your system, already.
The host machine must be a running `Docker` engine. `Git` is also installed.
You need to ensure that Docker have been installed on your machine before using this configuration.
This project has been cerrently tested on Docker Desktop 4.26.1. 

## Overview
This `docker-compose.yml` file is designed to facilitate the deployment of the MongoDB database. 
You can edit services based on your specific requirements. 

For development purposes, ensure you have the necessary tools and dependencies installed. 
Refer to the `go.mod` file for Go-related dependencies. 

`main.go` file is this application entry point.

Example of spider's data is stored `data_demoDBspider.json ` for demostration purpose.

## Deployment
To deploy this database as demo, you can use the provided `docker-compose.yml` file. 
You can customize it according to your necessary Docker commands.


## Usage
Open a terminal and run the following commands
 
1. Clone the repository.
```bash
git clone https://github.com/tuscb/spider-th-public.git
```

2. Change to the project directory.
```bash
cd spider-th-public
```

3. Initialize `go.mod` with module name go-spiderth and run `go mod tidy` to add missing and remove unused modules.
```bash
go mod init go-spiderth
```
```bash
go mod tidy
```

4. Review the Docker Compose configuration in `docker-compose.yml`, (optional) modify code such as username, password and others. 

Note that username is `user` and password is `1234`.  

5. Run the Docker Compose command to start the services
```bash
docker-compose up -d
```
Note that `docker-compose down`command is for stoping it.

6. Check whether Docker is sucessfully running
```bash
docker ps
```

7. If  username and password are changed on step 4, ensure to change them at line 66 in `main.go` file, too.
```
// Set your MongoDB connection string
uri := "mongodb://user:1234@localhost:27017"
```

8. Now run `main.go` file to import `data_demoDBspider.json` file which spider's data items are stored as the proposed design data structure. 
```bash
go run main.go -i data_demoDBspider.json 
```

9. Example to query spider data item 10th from the database after `data_demoDBspider.json` file already added into the database on the previous step.
```bash
go run main.go -q 10
```
Note that you can change 10 to the desired number of record.

## Summary
This is not a production MongoDB configuration. It simply iillustrates how to add and pull data from MongDB database including how data are stored in json file the most efficient method for successfully importing MongoDB code.

## Contact
If you have any suggestion or find a bug, please contact us by email.