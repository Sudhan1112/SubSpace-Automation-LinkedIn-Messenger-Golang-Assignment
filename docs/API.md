# API Documentation

The backend exposes a simple REST API to control the automation and retrieve status.

## Endpoints

### 1. Get System Status
Returns the current running state and the latest logs.

- **Endpoint**: `/api/status`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "running": boolean,
    "logs": ["log line 1", "log line 2", ...]
  }
  ```

### 2. Start Automation
Triggers the automation engine to begin the task.

- **Endpoint**: `/api/start`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "query": "string (e.g., 'Software Engineer')",
    "email": "string (optional)",
    "password": "string (optional)"
  }
  ```
- **Response**: `200 OK` (if started), `400 Bad Request` (if already running or invalid input).

### 3. Stop Automation
Signals the automation engine to halt gracefully.

- **Endpoint**: `/api/stop`
- **Method**: `POST`
- **Response**: `200 OK`

### 4. Get Scraped Data
Retrieves the list of profiles scraped during the session.

- **Endpoint**: `/api/data`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "ID": 1,
      "ProfileURL": "https://linkedin.com/in/...",
      "Metadata": "Name - Title",
      "Action": "SEARCH_FOUND",
      "Timestamp": "2023-10-27T10:00:00Z"
    },
    ...
  ]
  ```
