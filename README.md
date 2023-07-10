Here's a sample README file for my GitHub repository based on the main.go code:

# Task Management API

This is a simple Task Management API implemented in Go using the Gin framework and SQLite as the database.

## Requirements

- Go 1.16 or higher
- SQLite

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/task-management-api.git
   ```

2. Navigate to the project directory:

   ```shell
   cd task-management-api
   ```

3. Build and run the application:

   ```shell
   go run .
   ```

   The API server will start running on `http://localhost:8080`.

## Endpoints

### Create a Task

**URL:** `/tasks`

**Method:** `POST`

**Request Body:**

```json
{
  "title": "Task 1",
  "description": "This is task 1",
  "due_date": "2023-07-05",
  "status": "Pending"
}
```

**Response:**

```json
{
  "id": 1,
  "title": "Task 1",
  "description": "This is task 1",
  "due_date": "2023-07-05",
  "status": "Pending"
}
```

### Get a Task

**URL:** `/tasks/:id`

**Method:** `GET`

**Response:**

```json
{
  "id": 1,
  "title": "Task 1",
  "description": "This is task 1",
  "due_date": "2023-07-05",
  "status": "Pending"
}
```

### Update a Task

**URL:** `/tasks/:id`

**Method:** `PUT`

**Request Body:**

```json
{
  "title": "Updated Task",
  "description": "This is an updated task",
  "due_date": "2023-07-06",
  "status": "Completed"
}
```

**Response:**

```json
{
  "id": 1,
  "title": "Updated Task",
  "description": "This is an updated task",
  "due_date": "2023-07-06",
  "status": "Completed"
}
```

### Delete a Task

**URL:** `/tasks/:id`

**Method:** `DELETE`

**Response:**

```json
{
  "message": "Task deleted"
}
```

### List All Tasks

**URL:** `/tasks`

**Method:** `GET`

**Response:**

```json
[
  {
    "id": 1,
    "title": "Task 1",
    "description": "This is task 1",
    "due_date": "2023-07-05",
    "status": "Pending"
  },
  {
    "id": 2,
    "title": "Task 2",
    "description": "This is task 2",
    "due_date": "2023-07-06",
    "status": "In Progress"
  }
]
```

## Database Schema

The API uses an SQLite database with the following schema for the `tasks` table:

```sql
CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  description TEXT,
  due_date TEXT,
  status TEXT
);
```
## Workflow Architecture
![workf](https://github.com/Itspranav025/TaskManagementAPI/assets/88264892/d9fa752e-b841-4e4d-be2e-5f5d4f4f28b0)

## Application Architecture
![ApplicationArchitecture](https://github.com/Itspranav025/TaskManagementAPI/assets/88264892/826c779c-d17b-44ba-8f08-29e62549de7b)
