# Go REST API with GORM

A simple yet powerful REST API built with Go, using GORM as the ORM and MySQL as the database.

## 👨‍💻 Author

**Juan Jesus Martin Melero**

## 🚀 Technologies

-   **Go** - Programming language
-   **GORM** - Object Relational Mapper (ORM)
-   **MySQL** - Database
-   **Gorilla Mux** - HTTP router and URL matcher
-   **Godotenv** - Environment variable loader

## 📝 Project Structure

```
.
├── db/
│   └── database.go    # Database connection configuration
├── handlers/
│   ├── handlers.go    # HTTP request handlers
│   └── response.go    # Response utilities
├── models/
│   └── users.go      # User model definition
└── main.go           # Application entry point
```

## 🛣️ API Endpoints

| Method | Endpoint    | Description             |
| ------ | ----------- | ----------------------- |
| GET    | /users      | Get all users           |
| GET    | /users/{id} | Get a specific user     |
| POST   | /users      | Create a new user       |
| PUT    | /users/{id} | Update an existing user |
| DELETE | /users/{id} | Delete a user           |

## 🛠️ Setup

1. Clone the repository
2. Create a `.env` file in the root directory with the following variables:
    ```
    DB_USER=your_username
    DB_PASSWORD=your_password
    DB_HOST=your_host
    DB_PORT=your_port
    DB_NAME=your_database
    DB_CHARSET=utf8mb4
    ```
3. Install dependencies:
    ```bash
    go mod download
    ```
4. Run the application:
    ```bash
    go run main.go
    ```

## 📦 Dependencies

-   github.com/gorilla/mux
-   gorm.io/gorm
-   gorm.io/driver/mysql
-   github.com/joho/godotenv

## 🔒 Environment Variables

The application uses environment variables for database configuration. Make sure to set up your `.env` file with the correct values:

-   `DB_USER`: Database username
-   `DB_PASSWORD`: Database password
-   `DB_HOST`: Database host
-   `DB_PORT`: Database port
-   `DB_NAME`: Database name
-   `DB_CHARSET`: Database charset (default: utf8mb4)

## 📜 License

This project is open source and available under the [MIT License](LICENSE)
