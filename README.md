# Go E-Commerce API

A backend e-commerce API built with Go, Gin, and MongoDB.

## Tech Stack

- Go
- Gin (`github.com/gin-gonic/gin`)
- MongoDB (`go.mongodb.org/mongo-driver`)
- JWT (`github.com/dgrijalva/jwt-go`)

## Features

- User signup and login
- JWT-based authentication middleware
- Product creation and search
- Cart operations (add, remove, checkout, instant buy)
- User address management

## Project Structure

- `main.go` - application entry point
- `routes/` - route registration
- `controllers/` - HTTP handlers
- `database/` - MongoDB setup and DB helpers
- `middleware/` - auth middleware
- `models/` - request/response and DB models
- `tokens/` - JWT generation and validation

## Prerequisites

- Go (1.26+ recommended by `go.mod`)
- MongoDB running on `mongodb://localhost:27017`

Optional:

- Docker and Docker Compose

## Environment Variables

Set these before running the app:

- `PORT` (optional, defaults to `8000`)
- `SECRET_KEY` (required for access token signing/validation)
- `REFRESH_SECRET` (required for refresh token signing)

PowerShell example:

```powershell
$env:PORT="8000"
$env:SECRET_KEY="your-secret-key"
$env:REFRESH_SECRET="your-refresh-secret"
```

## Running Locally

1. Install dependencies:

```bash
go mod tidy
```

2. Make sure MongoDB is running on `localhost:27017`.

3. Start the API:

```bash
go run .
```

Server starts on `http://localhost:8000` by default.

## Run with Docker Compose (Mongo + Mongo Express)

This project includes `docker-compose.yaml` for MongoDB and Mongo Express:

```bash
docker compose up -d
```

Mongo Express will be available at `http://localhost:8081`.

## API Endpoints

### Public

- `POST /users/signup`
- `POST /users/login`
- `POST /admin/addproduct`
- `GET /users/productview`
- `GET /users/search?name=<query>`

### Authenticated (requires token)

Send JWT using either:

- `Authorization: Bearer <token>`
- `token: <token>`

Routes:

- `GET /addtocart?id=<productObjectId>&userID=<userId>`
- `GET /removeitem?id=<productObjectId>&userID=<userId>`
- `GET /cartcheckout?id=<userId>`
- `GET /instantbuy?id=<productObjectId>&userID=<userId>`
- `POST /address/:id`
- `DELETE /address/:id`

## Notes

- The authenticated routes rely on token claims for user context.
- Ensure `SECRET_KEY` and `REFRESH_SECRET` are set, or login/token validation will fail.
- Product and user data are stored in the `Go-Ecommerce` MongoDB database.
