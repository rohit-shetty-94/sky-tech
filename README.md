# Metrics Ingestion and API Service

## Overview

This service ingests time-series metrics and provides a REST API to query them by time range.

## Features

- RESTful API using Echo
- PostgreSQL storage via go-pg
- Unit test coverage
- Middleware logging
- Configurable via environment variables

## Ingestion Script
- Each execution loads data for the last 5 minutes at 1-minute intervals.

- Before inserting, it checks the latest existing timestamp in the database to avoid duplicate entries.

```
Example behavior:

If the script runs at 00:45, it ingests data from 00:41 to 00:45.

If it runs again at 00:47, it detects existing data up to 00:45 and only ingests data for 00:46 and 00:47.
```
## REST API Endpoints
Exposes a GET endpoint to retrieve metrics based on a start and end epoch timestamp:

- `GET /metrics?start=<epoch>&end=<epoch>`  

    | Method | Endpoint      | Description                      |
    | ------ | ------------- | -------------------------------- |
    | GET    | `/metrics`     | Retrieve metrics with time range |

    `Example : http://localhost:8080/metrics?start=1747763667&end=1747767010`

## Setup & Installation

### Pre-requisite
- Create database in SQL
    ```bash
    CREATE DATABASE sky_metrics
    WITH
    OWNER = postgres 
    ```
- Clone the Repository
    ```bash
    git clone https://github.com/rohit-shetty-94/sky-tech.git
    cd sky-tech
    ```
- Create .env (with DB credential)
    ```bash
    export DB_USER=*****
    export DB_PASSWORD=*****
    export DB_NAME=*****
    ```
- Pull all dependence 
    ```bash
    go mod tidy
    ```
- Execute the code
    ```
    go run main.go
    ```
## Running Unit Tests

To run the test suite:

```bash
go test ./...
```

## Author

**Rohit Shetty**\
**GitHub**: [rohit-shetty-94](https://github.com/rohit-shetty-94)\
**Email**: [shettyrohit61@gamil.com](mailto\:shettyrohit61@gamil.com)
