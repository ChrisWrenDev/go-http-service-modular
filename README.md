# HTTP Service in Go

This project is a modular and maintainable HTTP service written in Go. It is inspired by best practices outlined by Mat Ryer, focusing on simplicity, scalability, and clarity.

[https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)

## Project Structure

The project is divided into several files, each serving a specific purpose:

### **1. `main.go`**

- The entry point of the application.
- Sets up the server configuration, initializes dependencies, and starts the HTTP server.
- Handles graceful shutdown of the server using signal handling.

### **2. `server.go`**

- Contains the `NewServer` function, which sets up the HTTP server.
- Combines middleware, routing, and dependency injection into a single cohesive server handler.

### **3. `routes.go`**

- Defines all the application routes.
- Associates URL paths with their respective HTTP handler functions.
- Ensures routes are logically grouped and easy to extend.

### **4. `handlers.go`**

- Implements HTTP handler functions for each route.
- Encapsulates business logic, making each handler small and focused.

### **5. `middleware.go`**

- Contains middleware functions that wrap HTTP handlers.
- Examples:
  - **Logging Middleware**: Logs each incoming request.
  - **Authentication Middleware**: Verifies user authentication (placeholder logic).

### **6. `logger.go`**

- Centralized logging setup.
- Uses Go’s standard `log` package to provide consistent log output.

### **7. `config.go`**

- Defines configuration options such as server host and port.
- Allows for easy configuration changes without modifying other files.

### **8. `stores.go`**

- Abstracts data storage logic (e.g., database or in-memory storage).
- Defines structs and interfaces for different storage components.
- Examples:
  - `CommentStore`: Handles comment-related data operations.
  - `AnotherStore`: Handles operations for other business logic.

---

## Key Patterns and Architecture

### **1. Dependency Injection**

- Dependencies like `logger`, `config`, and `stores` are passed into the `NewServer` function.
- Encourages loose coupling and makes testing easier.

### **2. Middleware Pattern**

- Middleware wraps HTTP handlers, providing reusable functionality like logging and authentication.
- Example:
  ```go
  func loggingMiddleware(logger *log.Logger) func(http.Handler) http.Handler
  ```
