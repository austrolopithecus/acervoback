# Acervo Comics

**Acervo Comics** is a comic book collection and trading platform built using **Go**. Users can register, login, catalog their comic books using ISBN, trade comics with other users, and leave reviews. This project is designed following the **Domain-Driven Design (DDD)** architecture pattern.

## Features

- User registration and authentication using JWT
- Comic book cataloging with ISBN
- Trading comic books between users
- Review system for comics
- Comprehensive API documentation using Swagger

## Table of Contents

- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Dependencies](#dependencies)
- [License](#license)

## Installation

To get started with **Acervo Comics**, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/austrolopithecus/acervoback.git
    cd acervoback
    ```

2. Install Go dependencies:

    ```bash
    go mod tidy
    ```

3. Set up a PostgreSQL database and configure environment variables in a `.env` file at the root of the project:

    ```
    DATABASE_URL=postgres://user:password@localhost:5432/acervo
    JWT_SECRET=your-secret-key
    ```

4. Install PostgreSQL and create a database named `acervo`.

5. Ensure that PostgreSQL is running and properly configured to accept connections.

## Running the Application

Once the environment is set up, run the application:

```bash
go run main.go

Project Structure

The project follows the DDD (Domain-Driven Design) approach and is divided into the following layers:

    handlers/: Contains the HTTP handlers responsible for processing incoming requests and sending responses.

    middlewares/: Middleware components, such as JWT authentication, which are used to secure endpoints.

    services/: Contains the business logic of the application. It acts as an intermediary between the handlers and the repository layer.

    repositories/: Data access layer that interacts with the PostgreSQL database using GORM.

    models/: Defines the core entities and database models such as User, Comic, Exchange, and Review.

    requests/ & responses/: Defines the structure for incoming requests and outgoing responses to standardize communication between the API and the client.

API Endpoints
User Management

    Register a new user: POST /user/register
    Login: POST /user/login
    Get current user information: GET /user/me (requires JWT)

Comic Book Management

    Add a comic to the collection by ISBN: PUT /comic
    View all comics for the logged-in user: GET /comic

Comic Trading

    Request a comic trade: POST /exchange
    Accept a trade request: POST /exchange/:id/accept
    Complete a trade: POST /exchange/:id/complete

Reviews

    Add a review to a comic: POST /comic/:id/review
    Get all reviews for a comic: GET /comic/:id/reviews
