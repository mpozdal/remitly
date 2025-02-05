# Bank Swift Codes Application 🏦💳

This application is built using **Golang** and **MySQL** with Docker. It allows you to manage bank data by loading it from a CSV file, and then interact with it via various endpoints.

## Features 🚀

-   **Load bank data from a CSV file** containing information about banks.
-   **Display bank details** using the SWIFT code.
-   **Filter banks by country**.
-   **Add new bank entries**.
-   **Delete a bank entry** by SWIFT code.

## Endpoints 📡

-   `GET /v1/swift-codes/{swift-code}`: Display information about a bank by its SWIFT code.
-   `GET /v1/swift-codes/country/{countryISO2code}`: Show all banks from a specific country.
-   `POST /v1/swift-codes`: Add a new bank entry.
-   `DELETE /v1/swift-codes/{swift-code}`: Delete a bank entry by SWIFT code.

## Running the Application with Docker 🐳

To get the application up and running using Docker, follow these steps:

1. Clone the repository to your local machine:
    ```bash
    git clone https://github.com/mpozdal/remitly.git
    cd remitly
    ```
2. Build the application using docker-compose:
    ```bash
    docker-compose build --no-cache
    ```
3. Run tests:
    ```bash
    docker-compose up test
    ```
4. Start the application:
    ```bash
    docker-compose up
    ```
5. The application will now be accessible at http://localhost:8080


## Requirements ⚙️

-   **Docker**
-   **Docker Compose**
