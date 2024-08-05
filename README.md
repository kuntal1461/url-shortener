# URL Shortener

A URL Shortener application built with Go, PostgreSQL, and deployed on Vercel. This application allows users to shorten long URLs and retrieve the original URLs using short links.

## Features

- **Shorten Long URLs:** Convert long URLs into concise, easy-to-share links.
- **Retrieve Original URLs:** Use short links to redirect to the original URLs.
- **Health Check Endpoint:** Monitor the application status for uptime and responsiveness.
- **Configurable:** Tailored for both local development and production deployment on Vercel.
- **Static Frontend:** A simple, interactive interface for users.



## Setup Instructions

### Prerequisites

- **Go 1.16+**
- **PostgreSQL**
- **Git**
- **Vercel Account** for deployment

### Local Development

1. **Clone the Repository**
    ```bash
    git clone https://github.com/yourusername/url-shortener.git
    cd url-shortener
    ```

2. **Set Up PostgreSQL**
    - Create a database named `url-shortener`.
    - Use SQL queries from `db-document/create_tables.sql` to set up your database tables.

3. **Configure Environment Variables**
    - Refer to the `.env` file to set up necessary environment configurations.

4. **Install Dependencies**
    ```bash
    go mod tidy
    ```

5. **Run the Application**
    ```bash
    go run cmd/main/main.go
    ```
    - Access the application at `http://localhost:8080`.

### Vercel Deployment

1. **Create a Vercel Project**
    - Set up a new project on Vercel and link it to your GitHub repository.

2. **Configure Environment Variables on Vercel**
    - Set up the necessary environment variables in the Vercel dashboard.

3. **Deploy**
    - Automatically deploy your project when changes are pushed to the main branch.

## Usage

- **Shorten URL:** POST request to `/api/shorten` with the long URL in the body.
- **Retrieve Original URL:** Navigate using the short link provided by the application.
- **Health Check:** Visit `/health` to verify the application's status.
- **Interact Through Static Page:** Open `index.html` from the `static` directory.

## Code Overview

- **`cmd/main/main.go`**: Main entry point of the application.
- **`configs/config.go`**: Manages configuration settings from environment variables.
- **`controllers/page_controller.go`**: Manages rendering of static pages.
- **`restcontrollers/url_rest_controller.go`**: Handles RESTful API requests.
- **`services/url_service.go`, `servicesimpl/url_service_impl.go`**: Define and implement the business logic.
- **`models/url.go`**: Data model for the URL records.
- **`repositories/url_repository.go`**: Manages database operations.
- **`db/db.go`**: Sets up database connections.
- **`static/`**: Contains static files for the frontend interface.

## Contributing

Contributions are welcome! Please submit pull requests or open issues for any enhancements, bug fixes, or feature requests.

## License

This project is licensed under the MIT License. Refer to the LICENSE file for more details.

---
Ensure to replace placeholders like `yourusername` with your actual GitHub username and update configuration details as per your environment setup.



## Project Structure

```plaintext
url-shortener/
├── cmd/
│   └── main/
│       └── main.go                # Entry point for the application
├── configs/
│   └── config.go                  # Configuration loading logic
├── controllers/
│   └── page_controller.go         # Handles rendering of static pages
├── restcontrollers/
│   └── url_rest_controller.go     # Handles REST API requests for URL operations
├── services/
│   └── url_service.go             # Business logic interfaces
├── servicesimpl/
│   └── url_service_impl.go        # Implementation of business logic
├── models/
│   └── url.go                     # Data model definitions
├── repositories/
│   └── url_repository.go          # Database operations
├── db/
│   └── db.go                      # Database connection setup
├── db-document/                   # Contains database schema and queries
│   └── create_tables.sql          # SQL queries for creating tables
├── static/                        # Static assets for the frontend
│   ├── index.html
│   ├── script.js
│   └── style.css
├── go.mod                         # Go module dependencies
└── go.sum                         # Go module checksums 

