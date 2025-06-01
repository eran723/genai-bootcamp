# Language Learning Portal Backend

This is the backend service for the Language Learning Portal, built with Go, Gin, and SQLite.

## Features

- Vocabulary management with Japanese, Romaji, and English translations
- Study session tracking
- Learning progress monitoring
- Word grouping and organization
- Review history tracking

## Prerequisites

- Go 1.21 or later
- SQLite3

## Project Structure

```
backend/
├── cmd/                    # Application entrypoints
│   └── server/            # Main server application
├── internal/              # Private application code
│   ├── api/              # API handlers and routes
│   ├── models/           # Database models
│   ├── database/         # Database connection and queries
│   └── service/          # Business logic
├── db/ 
│   ├── migrations/           # SQL migration files
│   ├──  seeds/               # Seed data files
├── magefiles/           # Mage task definitions
├── go.mod              # Go module file
└── go.sum              # Go dependencies checksum
```

## Setup

1. Clone the repository
2. Navigate to the backend directory
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the migrations:
   ```bash
   # TODO: Add migration command
   ```
5. Start the server:
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### Words

- `GET /api/words` - List words (paginated)
- `GET /api/words/:id` - Get a specific word
- `POST /api/words` - Create a new word
- `PUT /api/words/:id` - Update a word
- `DELETE /api/words/:id` - Delete a word

### Dashboard

- `GET /api/dashboard/last_session` - Get last study session details
- `GET /api/dashboard/stats` - Get study statistics
- `GET /api/dashboard/progress` - Get learning progress

## Development

To run the server in development mode:

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

## Testing

To run tests:

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details. 