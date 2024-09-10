
# Distributed Stock Trading Simulator

## Overview
This project is a backend implementation of a distributed stock trading simulator. It consists of several services that interact to simulate real-time stock trading, where users can place buy and sell orders and get updates on stock prices in real-time.

### Components:
1. **User Service**: Handles user registration and login.
2. **Trading Service**: Handles buying and selling of stocks, interacting with user portfolios, and ensuring real-time stock price updates.
3. **Stock Service**: Publishes random stock price updates to a Kafka topic.
4. **Middleware**: JWT-based authentication to secure routes in the trading service.

## Technologies Used
- **Golang**: For building the services.
- **Kafka**: For real-time message streaming of stock prices.
- **Redis**: For caching stock prices.
- **PostgreSQL**: For persisting user data, orders, and portfolios.
- **Docker**: To containerize the services.
- **JWT**: For user authentication.

## Getting Started

### Prerequisites
- Docker and Docker Compose
- Golang
- Kafka
- Redis
- PostgreSQL

### Setting up the Project

1. **Clone the repository**:
    ```bash
    git clone https://github.com/JorkDev/stock-trading-simulator
    cd stock_trading_simulator
    ```

2. **Run the services using Docker Compose**:
    Ensure that Kafka, Redis, and PostgreSQL are running in containers:
    ```bash
    docker-compose up -d
    ```

3. **Set up the environment**:
    In each service directory (`user-service`, `trading-service`, and `stock-service`), initialize Go modules:
    ```bash
    go mod init your_module_name
    go mod tidy
    ```

4. **Run the services**:

    - **User Service**:
        ```bash
        cd user-service
        go run main.go
        ```

    - **Trading Service**:
        ```bash
        cd trading-service
        go run main.go
        ```

    - **Stock Service**:
        ```bash
        cd stock-service
        go run main.go
        ```

## API Endpoints

### User Service
- **POST /register**: Register a new user.
    - Request:
        ```json
        {
            "username": "john_doe",
            "email": "john@example.com",
            "password": "securepassword"
        }
        ```

- **POST /login**: Log in and get a JWT token.
    - Request:
        ```json
        {
            "email": "john@example.com",
            "password": "securepassword"
        }
        ```

### Trading Service
- **POST /trade/buy**: Place a buy order (JWT required).
    - Request:
        ```json
        {
            "stock_symbol": "AAPL",
            "quantity": 10
        }
        ```

- **POST /trade/sell**: Place a sell order (JWT required).
    - Request:
        ```json
        {
            "stock_symbol": "AAPL",
            "quantity": 5
        }
        ```

### Stock Service
- **Stock prices** are published to the Kafka topic `stock-prices` in real-time.

## Future Improvements
- Add WebSocket support for real-time stock price and portfolio updates.
- Implement an order matching engine for market and limit orders.
- Build a front-end to interact with the services.

## License
This project is licensed under the MIT License.
