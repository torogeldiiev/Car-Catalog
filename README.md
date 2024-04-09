# Car Catalog Application

## Overview
This is a simple car catalog application that allows users to manage cars and people information.

## Setup
1. **Database Configuration**:
   - Ensure you have PostgreSQL installed and running.
   - Create a new PostgreSQL database for the application with this name 'car_catalog'
   - Update the database connection configuration in the `.env` file or directly in the `database.go` file.
2. **Environment Variables**:
   
   -Change .env file in config package(set your variables) or edit them in run configs
3. **Install Dependencies**:
   - Run the following command to install project dependencies:
     ```
     go mod tidy
     ```
   This command will automatically download and install the required dependencies listed in the `go.mod` file.

4. **Run Migrations**:
   - Ensure you have Goose installed (`go get -u github.com/pressly/goose/cmd/goose`).
   - Run the migrations to set up the database schema:
     ```
     goose -dir database postgres "postgres://your_username:your_password@localhost:5432/car_catalog" up
     ```
    
5. **Build and Run**:
go build -o car_catalog
./car_catalog

markdown
Copy code

6. **API Endpoints**:
- The application exposes the following API endpoints:
  - `/cars`: Manage cars (create, read, update, delete).
  - `/people`: Manage people (create, read, update, delete).

## Usage
- Access the API using your preferred API client (e.g., Postman).
- Send requests to the appropriate endpoints to interact with the application.
- Refer to the API documentation or code comments for more details on request formats and responses.

## Example Postman Requests

### GET
-- http://localhost:8080/cars/get?criteria=mark='Lada'&limit=10&offset=0  

-- http://localhost:8080/cars/get?criteria=year>2000&limit=5&offset=0  

-- http://localhost:8080/cars/get?criteria=reg_num='1010BM'&limit=5&offset=0  

-- http://localhost:8080/cars/get?criteria=cars.id>1&limit=3&offset=0  

--http://localhost:8080/people/get?id=2

### CREATE
http://localhost:8080/cars/create

request body - raw , json format

{
    "regNums" : ["IT2214" , "KRUIZ44"]
}


http://localhost:8080/people/create
{
    "name": "Jon",
    "surname": "Doe",
    "patronymic":"Alice"

}

### DELETE
http://localhost:8080/cars/delete?id=10
http://localhost:8080/people/delete?id=3

### UPDATE
http://localhost:8080/cars/update?id=7
request body - raw , json format
{
    "Year": 2005
}

http://localhost:8080/people/update?id=2
{
    "patronymic":"Alex"
}

## Multithreading for External API Interactions

If your application interacts with external APIs that may introduce latency, you might consider implementing multithreading to improve performance and responsiveness.
Multithreading allows your application to handle multiple tasks concurrently, such as making multiple API requests simultaneously, without blocking the main execution flow.

This can be implemented later if service requires it.

## Contributing
Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or create a pull request.

## License
This project is licensed under the [MIT License](LICENSE).
