# Entain Task

This project is a Go-based application that provides services for managing user balances and transactions. It includes APIs for retrieving user balances and posting new transactions.

## Prerequisites

- Docker compose

## Running the application

1. Clone the repository:
   ```sh
    git clone https://github.com/Weeping-Willow/entain-task.git
    cd entain-task
   ```
2. Run the application:
   ```sh
    docker-compose up -d
   ```
3. Access the application at `http://localhost:9000`   

## Project Structure
- internal/api: Contains the API handlers and server setup.
- internal/service: Contains the business logic and service implementations.
- internal/repository: Contains the repository interfaces and implementations for database interactions.
- pkg/oapi/api: Contains the OpenAPI generated code. 
- api: Contains the OpenAPI specification file.
- migrations: Contains the database migration files.

## API Endpoints
Described in the Task.MD file

## Performance
Using postman to test the application with the following settings:
- Load profile: Peak
- Virtual users: 100
- Test duration: 1 minute
- Base Load: 50 users

the route in question that was tested was the `POST /user/{userId}/transaction` route.
for a singular user it just increased the balance by 1.00 for each request.

The application got these results:
- Total Requests: 4245
- Requests per second: 63.35
- Response time avg: 5ms
- Response time max: 93ms
- Response time min: 2ms
- Response time 90th percentile: 6ms
- Error rate: 0%

The application was able to handle the load without any issues.

## Testing
By default the application comes with 5 predefined users with ids 1, 2, 3, 4, 5.

### Unit Tests
To run the unit tests:
```sh
go test ./...
```
### Individual routes
- `POST /user/{userId}/transaction`
  ```sh
    curl --location 'localhost:9000/user/3/transaction' \
    --header 'Source-Type: game' \
    --header 'Content-Type: application/json' \
    --data '{"state": "win", "amount": "1", "transactionId": "random_id"}'
  ```
Note: The `transactionId` field should be a new value for each request

- `GET /user/{userId}/balance`
  ```sh
    curl --location 'localhost:9000/user/3/balance'
  ```

### Postman collection
A postman collection is provided in the root directory of the project. It contains the requests for the two routes mentioned above.
the file is called `entail.postman_collection.json`