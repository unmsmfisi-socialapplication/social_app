# Social App

This is a backend written in go. It uses `go-chi` as a HTTP router.
We use a standard go layout. [See](https://github.com/golang-standards/project-layout)

## Table of Contents

- [Social App](#social-app)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Usage](#usage)
  - [Docker](#docker)
    - [Prerequisites](#docker-prerequisites)
    - [Getting Started](#docker-getting-started)
    - [Stopping the application](#stopping-the-application)
    - [Troubleshooting](#troubleshooting)

## Getting Started

### Prerequisites

- Visual Studio Code or other IDE
- Go  1.20

### Use Enviroment Variables
In order to use the environment variables you must duplicate in an `.env` file the content of the `.env.example` file, this file will contain all the necessary variables for the application to work.

### Use DataBase Migrations
1. Properly configure environment variables.
2. Execute the command `go run cmd/migrations/migrations.go` in order to perform the migrations correctly.

### Installation

to install use 

```bash
go mod vendor
```

## Usage

for run tests use 

```bash
go test ./...
```

## Docker

### Docker Prerequisites

- Docker
- Docker compose plugin

### Docker Getting started

1. Build the docker image
```bash
docker-compose build
```
2. Start the containers
```bash
docker-compose up -d
```
3. Access to the application in a web browser or using a tool like curl:
```bash
curl http://localhost:3333
```
4. Access Adminer to manage your PostgreSQL database: <br>
Open a web browser and go to `http://localhost:8080`. Use the following credentials:
- System: PostgreSQL
- Server: db
- Username: user
- Password: password
- Database: social_app

### Stopping the application
To stop the running containers, use the following command
```bash
docker-compose up -d
```

### Troubleshooting
If you encounter issues or errors while running the Go application with Docker Compose, here are some common troubleshooting steps:
Check the logs of the containers for error messages:
```bash
docker-compose logs
```
Ensure that application is configured correctly to work with the database or any other services it depends on.
To initialize the application (in the backend directory)
