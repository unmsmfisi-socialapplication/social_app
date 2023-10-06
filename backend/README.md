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


## Getting Started


### Prerequisites

- Visual Studio Code or other IDE
- Go  1.20

### Use Enviroment Variables
In order to use the environment variables you must duplicate in an `.env` file the content of the `.env.example` file, this file will contain all the necessary variables for the application to work.

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
