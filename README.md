# Backend Challenge - FinTech Solutions Inc.

[![Stori](https://img.shields.io/badge/Stori%20Card-002C30)](https://www.storicard.com/)
[![Golang](https://img.shields.io/badge/Golang-00ADD8)](https://go.dev/)

This project involves creating two essential API services to migrate historical transaction records from CSV files to a robust database, as well as providing summarized balance information.

## Table of Contents

- [Project Description](#project-description)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API](#api)
  - [Migration Service](#migration-service)
  - [Balance Service](#balance-service)
- [Tests](#running-tests)
- [Documentation](#documentation)
- [Technologies Used](#technologies-used)
- [Author](#author)
- [Links](#links)

## Project Description

FinTech Solutions Inc. aims to improve data accessibility and transparency through the migration of transaction records to a scalable database system. This project includes two API services with hexagonal architecture to improve the scalability and maintainability of the code:

1. **Migration Service**: Processes transaction records from CSV files and stores them in a database.
2. **Balance Service**: Returns summarized information about balances.

## Features

- Processing of CSV files and migration to a database.
- User transactions balance summary via an API service.
- Monitoring migration status with email notifications.
- API documentation with Swagger.
- Unit tests for both services.
- Docker Compose implementation for easy setup and deployment.
- Health Check endpoints.
- Hexagonal architecture.

## Prerequisites

- Docker
- Docker Compose

## Installation

1.  Clone the repository:

    ```bash
    git clone https://github.com/AkselRivera/stori-challenge.git
    cd stori-challenge
    ```

2.  Rename `.env.template` file to `.env` and add productions variables

    ```bash
    docker compose up
    ```

    This will bring up the database and services on the specified ports.

## Usage

- Migration Service: Access the API at [http://localhost:8080/](http://localhost:8080/docs/) to upload a CSV file.

- Balance Service: Access the API at [http://localhost:8081/](http://localhost:8081/docs/) to get balance information.

## API

Each service implements an endpoint to check if they are online, this is the `Health` endpoint.

- GET /health
  - Check if the service is online
  - Response: Return a JSON with `status: ok`

### Migration Service

- POST /migrate
  - Uploads a `CSV` file and processes the transactions.
  - Response: If the `CSV` file is valid response `OK` Response otherwise it will return a `Bad Request` Response
  - Creates a goroutine and sends a summary of the results

### Balance Service

- GET /user/{user_id}/balance
  - Returns a summary of user balance.
  - Response: Balance information in JSON format.

## Running Tests

To run the unit tests, use the following command:

```bash
  // For Balance Service
  go test ./balance-service/...
  // or
  // For Migration service
  go test ./migration-service/...
```

> **_NOTE:_** There are only tests for services.

## Documentation

API documentation is available at `/docs/` endpoint once the services are running.

## Technologies Used

- Golang
- Gorm
- Docker
- Resend (for email sending)
- Swagger
- mockgen (for tests)

## Author

- [@AkselRivera](https://github.com/akselRivera)

## Links

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://mx.linkedin.com/in/aksel-morales-8b7844209)
