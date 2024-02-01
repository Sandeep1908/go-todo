# Golang TODO App with ScyllaDB

This is a simple Golang application for managing TODO items, using ScyllaDB as the database.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Endpoints](#endpoints)
- [Configuration](#configuration)
- [Connecting to ScyllaDB via Docker](#connecting-to-scylladb-via-docker)
- [Design Decisions](#design-decisions)
 

## Prerequisites

Make sure you have the following installed:

- Golang: [Install Golang](https://golang.org/doc/install)
- Docker: [Install Docker](https://docs.docker.com/get-docker/)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/Sandeep1908/go-todo.git
   cd your-todo-app

- go mod tidy
- go run main.go

Open your browser and visit http://localhost:8080 to access the API.

## Endpoints
- POST /todos: Create a new TODO item.
- GET /todos/:user_id: Retrieve TODO items for a specific user.
- PUT /todos/:user_id/:post_id: Update a TODO item.
- DELETE /todos/:user_id/:post_id: Delete a TODO item.

## Configuration
The application uses the default configuration for connecting to ScyllaDB.
You can customize the database connection settings in the initializeDB function in main.go.

## Connecting to ScyllaDB via Docker


1. Pull the ScyllaDB Docker image:
   ```bash
   docker pull scylladb/scylla:5.2.0


2. Run ScyllaDB in a Docker container:
   ```bash
    docker run --name scylla-container -d -p 9042:9042 scylladb/scylla:5.2.0

 
 3. Verify that ScyllaDB is running:
   ```bash
    docker logs scylla-container

    

4. Update the initializeDB function in main.go with the appropriate ScyllaDB connection settings.


## Design Decisions

- Database Schema: The TODO items are stored in a ScyllaDB table with columns like id, user_id, title, description, status, created, and updated.
- Pagination: Implemented a paginated list endpoint to retrieve TODO items with optional filtering and sorting.
- ScyllaDB: Chosen for its scalability, high availability, and wide-column store capabilities.