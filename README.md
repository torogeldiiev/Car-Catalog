# Car Catalog Application

## Overview
This is a simple car catalog application that allows users to manage cars and people information.

## Multithreading for External API Interactions

This application interacts with external APIs that may introduce latency, i consider implementing multithreading to improve performance and responsiveness later on.
Multithreading will allow this application to handle multiple tasks concurrently, such as making multiple API requests simultaneously, without blocking the main execution flow.

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


6. **API Endpoints**:
API Endpoints
The application exposes the following API endpoints:

Cars
Manage cars: Allows creation, retrieval, updating, and deletion of cars.
People
Manage people: Allows creation, retrieval, updating, and deletion of people.
Usage
Access the API using your preferred API client (e.g., Postman). Send requests to the appropriate endpoints to interact with the application. Refer to the API documentation or code comments for more details on request formats and responses.

Example Postman Requests
**GET Requests**
Retrieve cars filtered by any fields of cars:

Endpoint: http://localhost:8080/cars/get?criteria=mark='Lada'&limit=10&offset=0
Retrieve cars filtered by year:

Endpoint: http://localhost:8080/cars/get?criteria=year>2000&limit=5&offset=0
Retrieve cars by registration number:

Endpoint: http://localhost:8080/cars/get?criteria=reg_num='1010BM'&limit=5&offset=0
Retrieve cars by ID:

Endpoint: http://localhost:8080/cars/get?criteria=cars.id>1&limit=3&offset=0
Retrieve person by ID:

Endpoint: http://localhost:8080/people/get?id=2

**POST Requests**
Create Car:

Endpoint: POST http://localhost:8080/cars/create
Request Body: JSON
json  

{
    "regNums": ["IT2214", "KRUIZ44"]
}
Response: Status Code 201 Created  


Create Person:

Endpoint: POST http://localhost:8080/people/create
Request Body: JSON
json  

{
    "name": "Jon",
    "surname": "Doe",
    "patronymic": "Alice"
}
Response: Status Code 201 Created, Content: ID of the created person


**DELETE Requests**
Delete Car:

Endpoint: http://localhost:8080/cars/delete?id=10  

Delete Person:

Endpoint: http://localhost:8080/people/delete?id=3  

**UPDATE Requests**
Update Car:

Endpoint: http://localhost:8080/cars/update?id=7
Request Body: JSON
json  

{
    "Year": 2005
}

Update Person:  

Endpoint: http://localhost:8080/people/update?id=2
Request Body: JSON
json  

{
    "patronymic": "Alex"
}


## Contributing
Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or create a pull request.

## License
This project is licensed under the [MIT License](LICENSE).
