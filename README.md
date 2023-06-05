# navy [![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 

## Description

The Navy Management System is a comprehensive software solution designed to streamline and enhance the management 
processes within a navy organization. This project aims to provide an efficient and user-friendly platform for managing 
various aspects of navy operations, including personnel, vessels, missions, logistics, and more.

## Installation

### Prerequisites

- [golang >= 1.20.10](https://golang.org/doc/install)
- [docker >= 23.0.5](https://docs.docker.com/get-docker/)
- [docker-compose >= v2.17.3](https://docs.docker.com/compose/install/)

## Usage

### Building the application docker image

```bash
make docker-build
```
### Running the application

To run the application, simply run the following command from the root directory of the project:

```bash
make docker-up
```

This will start the application and all of its dependencies in docker containers. The application will be available at 
[http://localhost:8080](http://localhost:8080).

### Stopping the application

To stop the application, simply run the following command from the root directory of the project:

```bash
make docker-down
```

### Running the tests

To run the tests, simply run the following command from the root directory of the project:

```bash
make test
```

### Other useful commands

To find other useful commands, run the following command from the root directory of the project:

```bash
make help
```
