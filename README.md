# SampleCompany Computer Management System

This project provides a REST API for managing computers issued by SampleCompany. It allows the system administrator to store and retrieve details of computers, assign them to employees, and get notifications when an employee has more than three devices.

## Table of Contents

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Documentation](#documentation)
- [Notes](#notes)

## Getting Started

Follow these instructions to set up and run the project on your local machine.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/samplecompany.git
   cd samplecompany
   ```

2. Install Go dependencies:

   ```sh
   go mod tidy
   ```

3. Install PostgreSQL using Docker:

   Create a `docker-compose.yml` file in the project root directory with the following content:

   ```yaml
   version: "3.1"

   services:
     postgres:
       image: postgres
       restart: always
       environment:
        POSTGRES_USER: ${POSTGRES_USER}
        POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
        POSTGRES_DB: ${POSTGRES_DB}
       ports:
         - "5432:5432"
       volumes:
         - postgres-data:/var/lib/postgresql/data

   volumes:
     postgres-data:
   ```

   Start the PostgreSQL container:

   ```sh
   docker-compose up -d
   ```

4. Install Swaggo for generating Swagger documentation:

   ```sh
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

## Configuration

1. Create a `.env` file in the project root directory and add the following environment variables:

   ```env
   DATABASE_URL="postgres://postgres:yourpassword@localhost:5432/sampledb"
   PORT=8081
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=yourpassword
   POSTGRES_DB=sampledb
   ```

   Replace `postgres` and `yourpassword` with your PostgreSQL user and password.

## Usage

1. Run the server:

   ```sh
   make run
   ```

   The server will start on `http://localhost:8081`.

## Endpoints

- `POST /computers`: Add a new computer
- `GET /computers`: Get all computers
- `GET /computers/employee/:abbr`: Get all computers assigned to an employee
- `GET /computers/:id`: Get data of a single computer
- `PUT /computers/:id`: Update a computer
- `DELETE /computers/:id`: Delete a computer
- `PUT /computers/:id/assign`: Assign a computer to another employee

## Documentation

1. Generate API documentation using Swaggo:

   ```sh
   swag init
   ```

   This will create a `docs` directory with the Swagger documentation.

2. Serve the Swagger documentation:

   Add the following route to `main.go`:

   ```go
   import (
       "github.com/gin-gonic/gin"
       "github.com/swaggo/files"
       "github.com/swaggo/gin-swagger"
       _ "samplecompany/docs"
   )

   func main() {
       r := gin.Default()

       // Other routes

       r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

       r.Run()
   }
   ```

   Now you can access the API documentation at `http://localhost:8081/swagger/index.html`.

## Notes

- Ensure PostgreSQL is running and properly configured before starting the application.
- Set the `DATABASE_URL` environment variable correctly to avoid connection issues.
- Make sure to run `swag init` whenever you update the API documentation comments in the source code.
- To test the endpoints, use tools like Postman or cURL.

## Amendments

- **Database Configuration**: If you change the database configuration, update the `DATABASE_URL` in the `.env` file accordingly.
- **Port Configuration**: The server runs on port 8081 by default. You can change this by setting the `PORT` environment variable.
- **Swagger Documentation**: Keep the Swagger comments up-to-date with the latest changes in the API.

## Conclusion

This README provides detailed instructions for setting up, running, and documenting the SampleCompany Computer Management System. Following these steps ensures that the project is properly configured and that the documentation remains up-to-date.
