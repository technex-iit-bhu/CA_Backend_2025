## **API Route Documentation**

This document provides a structured overview of all API routes defined in the `router.go` file, detailing the HTTP methods, endpoint paths, and the purpose of each route.

---

### **Initialization and Middleware**

The router initializes the Fiber application and sets up middleware for cross-origin requests, logging, and route protection.
- **CORS**: Enables CORS for all routes.
- **Protected Middleware**: Adds a layer of protection for all `/api` routes.
- **Logger Middleware**: Logs requests under the `/api` group.

---

## **Route List**

### **Base API Route**

- **Method**: `GET`
- **Endpoint**: `/api/`
- **Handler**: `handler.Hello`
- **Description**: A base endpoint to confirm the API is running. It could be used as a health check.

---

## **User Routes**

Each user-related route is nested under `/api/user` and manages user operations, such as registration, login, profile retrieval, and more.

### **1. Get User Profile**

- **Method**: `GET`
- **Endpoint**: `/api/user/profile`
- **Handler**: `user_handler.GetUserProfile`
- **Description**: Fetches the profile details of the currently authenticated user.
- **Request Headers**: `Authorization` (Bearer token required).

### **2. Create User**

- **Method**: `POST`
- **Endpoint**: `/api/user/register`
- **Handler**: `user_handler.CreateUser`
- **Description**: Registers a new user in the system.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string",
    "phone": "string",
    "name": "string",
    "email": "string",
    "institute": "string",
    "city": "string",
    ...other fields
  }
  ```

### **3. Login User**

- **Method**: `PATCH`
- **Endpoint**: `/api/user/login`
- **Handler**: `user_handler.LoginUser`
- **Description**: Authenticates the user and returns a JWT token for session management.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

### **4. Set Referral**

- **Method**: `PATCH`
- **Endpoint**: `/api/user/Referral`
- **Handler**: `user_handler.Referral`
- **Description**: Increments the `referral_count` of a user based on a provided referral code while setting referral of user permanent.
- **Request Body**:
  ```json
  {
    "referral_code": "string"
  }
  ```

---

## **Password Management Routes**

These routes are under the `/api/user/password` path and handle password recovery and reset functionalities.

### **1. Password Recovery Request**

- **Method**: `POST`
- **Endpoint**: `/api/user/password/recovery`
- **Handler**: `user_handler.RequestPasswordRecovery`
- **Description**: Initiates the password recovery process, likely by sending an email with a reset link or token.
- **Request Body**:
  ```json
  {
    "email": "string"
  }
  ```

### **2. Reset Password**

- **Method**: `POST`
- **Endpoint**: `/api/user/password/reset`
- **Handler**: `user_handler.ResetPassword`
- **Description**: Resets the user’s password based on a provided token.
- **Request Body**:
  ```json
  {
    "token": "string",
    "new_password": "string"
  }
  ```

---

## **Handlers Reference**

Here’s a summary of the primary functionality of each handler for additional context:

- **`GetUserProfile`**: Validates JWT token, retrieves the user profile from the database.
- **`CreateUser`**: Parses user data, validates fields, checks for duplicates, hashes the password, and saves a new user in the database.
- **`LoginUser`**: Authenticates user credentials, returns JWT on successful authentication.
- **`IncrementReferral`**: Increments the `referral_count` field for a user identified by `referral_code`.
- **`RequestPasswordRecovery`**: Initiates password recovery by sending a recovery email or token.
- **`ResetPassword`**: Verifies a recovery token, hashes the new password, and updates the user’s password.

---

## **Usage Example**

To use these routes, send requests to the specified endpoints with appropriate HTTP methods, headers, and body data.

For instance:
- **Creating a User**: Send a `POST` request to `/api/user/register` with JSON body containing user details.
- **Incrementing a Referral**: Send a `PATCH` request to `/api/user/updateReferral` with the referral code in JSON format.

---

## **Notes**

- **Authentication**: Most routes require JWT tokens for authorization.
- **Error Handling**: Each route provides error messages for invalid requests or internal errors.
