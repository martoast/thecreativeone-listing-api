# MyCreativeFinancing API

## Introduction
This API serves as the backend for the MyCreativeFinancing application, designed to manage and manipulate property data efficiently. Built with Go and leveraging the Gorilla Mux router along with the GORM ORM for interacting with MySQL, this API provides a robust platform for real estate professionals to access and manage their property listings on their custom websites.


## Features
- CRUD operations on properties
- Sorting and filtering of property listings
- Authentication and environment variable management for secure database connections

## Installation

### Steps
1. Clone the repository:
   ```bash
   git clone [repository-url]
   ```
2. Set environment variables for MySQL in your .env file:
   ```makefile
   MYSQL_DATABASE=your_database
   MYSQL_ROOT_PASSWORD=your_password
   MYSQL_HOST=your_host
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```
   

### Prerequisites
- Go 1.15+
- MySQL 8.0+
- Docker (optional)


## Usage
Once the application is running, you can access the API at `http://localhost:8080`.

## API Endpoints
- **GET /properties/**  
  Returns a list of all properties.

- **GET /properties/{propertyId}**  
  Returns details of a specific property.

- **POST /properties/**  
  Creates a new property. Requires property details in the request body.

- **PUT /properties/{propertyId}**  
  Updates an existing property. Requires updated property details in the request body.

- **DELETE /properties/{propertyId}**  
  Deletes a specific property.
