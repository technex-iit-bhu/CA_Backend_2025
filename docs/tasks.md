# API Documentation

## Base Path: `/api`

---

## Task Routes

1. **Get All Tasks**
   - **Route**: `/tasks/`
   - **Method**: `GET`
   - **Description**: Retrieves a list of all tasks in the system.
   - **Response**: 
     - **200 OK**
       ```json
       [
         {
           "_id": "640e9fbc2f9c9f39a5401a5b",
           "title": "Complete the project documentation",
           "description": "Finish the detailed API documentation by the deadline.",
           "points": 50,
           "deadline": "2024-12-01T00:00:00Z",
           "image_url": "https://example.com/task-image.jpg"
         },
         ...
       ]
       ```
     - **500 Internal Server Error**
       ```json
       {
         "error": "Database connection error",
         "message": "Failed to fetch tasks"
       }
       ```

2. **Create Task**
   - **Route**: `/tasks/create`
   - **Method**: `POST`
   - **Description**: Creates a new task based on the provided task data.
   - **Request Body**:
     ```json
     {
       "title": "Complete the project documentation",
       "description": "Finish the detailed API documentation by the deadline.",
       "points": 50,
       "deadline": "2024-12-01T00:00:00Z",
       "image_url": "https://example.com/task-image.jpg"
     }
     ```
   - **Response**:
     - **201 Created**
       ```json
       {
         "id": "640e9fbc2f9c9f39a5401a5b",
         "message": "Task created successfully"
       }
       ```
     - **400 Bad Request**
       ```json
       {
         "error": "Failed to parse JSON Body",
         "message": "JSON parsing error"
       }
       ```
     - **500 Internal Server Error**
       ```json
       {
         "error": "Database connection error",
         "message": "Failed to create task"
       }
       ```

3. **Get Task by ID**
   - **Route**: `/tasks/task/:task_id`
   - **Method**: `GET`
   - **Description**: Retrieves details of a specific task using its `task_id`.
   - **Path Parameters**:
     - `task_id`: ID of the task to retrieve.
   - **Response**:
     - **200 OK**
       ```json
       {
         "_id": "640e9fbc2f9c9f39a5401a5b",
         "title": "Complete the project documentation",
         "description": "Finish the detailed API documentation by the deadline.",
         "points": 50,
         "deadline": "2024-12-01T00:00:00Z",
         "image_url": "https://example.com/task-image.jpg"
       }
       ```
     - **404 Not Found**
       ```json
       {
         "message": "Task does not exist"
       }
       ```
     - **500 Internal Server Error**
       ```json
       {
         "error": "Database connection error",
         "message": "Failed to retrieve task"
       }
       ```

4. **Update Task by ID**
   - **Route**: `/tasks/update/:task_id`
   - **Method**: `POST`
   - **Description**: Updates a task’s details based on the provided data.
   - **Path Parameters**:
     - `task_id`: ID of the task to update.
   - **Request Body**:
     ```json
     {
       "title": "Complete the project documentation",
       "description": "Updated description for the task",
       "points": 75,
       "deadline": "2024-12-05T00:00:00Z",
       "image_url": "https://example.com/updated-task-image.jpg"
     }
     ```
   - **Response**:
     - **200 OK**
       ```json
       {
         "message": "Task updated Successfully",
         "task": {
           "title": "Complete the project documentation",
           "description": "Updated description for the task",
           "points": 75,
           "deadline": "2024-12-05T00:00:00Z",
           "image_url": "https://example.com/updated-task-image.jpg"
         }
       }
       ```
     - **400 Bad Request**
       ```json
       {
         "error": "Failed to parse JSON Body",
         "message": "JSON parsing error"
       }
       ```
     - **404 Not Found**
       ```json
       {
         "message": "Task not found"
       }
       ```
     - **500 Internal Server Error**
       ```json
       {
         "error": "Database connection error",
         "message": "Failed to update task"
       }
       ```

---

## Submission Routes

1. **Submit Task**
   - **Route**: `/submissions/submit`
   - **Method**: `POST`
   - **Description**: Submits a task for review.
   - **Request Headers**:
     - `Authorization`: Bearer token for user authentication.
   - **Request Body**:
     ```json
     {
       "task": "640e9fbc2f9c9f39a5401a5b",
       "link": "https://submission-link.com",
       "image_url": "https://example.com/submission-image.jpg"
     }
     ```
   - **Response**:
     - **201 Created**
       ```json
       {
         "id": "641e0a1b2f9c9f39a5401c5d",
         "message": "Submitted successfully"
       }
       ```
     - **400 Bad Request**
       ```json
       {
         "message": "Authorization header missing or improperly formatted"
       }
       ```
     - **500 Internal Server Error**
       ```json
       {
         "error": "Database connection error",
         "message": "Failed to create task submission"
       }
       ```

2. **Get User Submissions**
   - **Route**: `/submissions/get_user_submissions`
   - **Method**: `GET`
   - **Description**: Retrieves all submissions made by the authenticated user.
   - **Request Headers**:
     - `Authorization`: Bearer token for user authentication.
   - **Response**:
     - **200 OK**
       ```json
       [
         {
           "_id": "641e0a1b2f9c9f39a5401c5d",
           "task": "640e9fbc2f9c9f39a5401a5b",
           "user": "username",
           "timestamp": "2024-11-10T10:00:00Z",
           "link": "https://submission-link.com",
           "image_url": "https://example.com/submission-image.jpg",
           "verified": false,
           "admin_comment": ""
         },
         ...
       ]
       ```
     - **400 Bad Request**
       ```json
       {
         "message": "Authorization header missing or improperly formatted"
       }
       ```
     - **500 Internal Server Error**
       ```json
       {
         "error": "Database connection error",
         "message": "Failed to fetch submissions"
       }
       ```

3. **Verify Submission by ID**
   - **Route**: `/submissions/verify/:submission_id`
   - **Method**: `GET`
   - **Description**: Verifies a submission and updates user points based on task points.
   - **Path Parameters**:
     - `submission_id`: ID of the submission to verify.
   - **Response**:
     - **200 OK**
       ```json
       {
         "message": "Task updated Successfully",
         "submission": {
           "_id": "641e0a1b2f9c9f39a5401c5d",
           "task": "640e9fbc2f9c9f39a5401a5b",
           "user": "username",
           "timestamp": "2024-11-10T10:00:00Z",
           "link": "https://submission-link.com",
           "image_url": "https://example.com/submission-image.jpg",
           "verified": true,
           "admin_comment": ""
         }
       }
       ```
     - **404 Not Found**
       ```json
       {
         "message": "Submission not found"
       }
       ```
     - **500 Internal Server Error**
       ```json
       {
         "error": "Database connection error",
         "message": "Failed to verify submission"
       }
       ```

---

## Data Models

- **Task**
  - Represents a task in the system with fields:
    ```json
    {
      "_id": "640e9fbc2f9c9f39a5401a5b",
      "title": "Task title",
      "description": "Task description",
      "points": 50,
      "deadline": "2024-12-01T00:00:00Z",
      "image_url": "https://example.com/image.jpg"
    }
    ```

- **TaskSubmission**
  - Represents a user’s task submission with fields:
    ```json
    {
      "_id": "641e0a1b2f9c9f39a5401c5d",
     

 "task": "640e9fbc2f9c9f39a5401a5b",
      "user": "username",
      "timestamp": "2024-11-10T10:00:00Z",
      "link": "https://submission-link.com",
      "image_url": "https://example.com/submission-image.jpg",
      "verified": false,
      "admin_comment": ""
    }
    ```

---

## Notes

- **Error Handling**: Standardized error messages help users understand issues, whether due to parsing errors, missing parameters, or database issues.
- **Authentication**: Some routes require `Authorization` headers for user authentication.
- **Database**: MongoDB is used, with error handling for potential database connection issues or query failures.