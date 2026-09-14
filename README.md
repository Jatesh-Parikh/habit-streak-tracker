# Go Habit Streak Tracker GraphQL API

A secure, full-stack GraphQL API built with Go, Gin, SQLite, and JWT authentication. This project demonstrates building a habit tracking system with daily check-ins, automatic streak calculations, and user authentication using modern GraphQL architecture.

> This project is based on a tutorial by [Muslim Halelee](https://www.youtube.com/watch?v=rrY7tcDSGZ8&list=PLHsjm_W8kcWZLPDxUplr8yredk95F-KIx&index=3).

## Features

- **GraphQL API**: Modern query language for APIs with type-safe schema
- **User Authentication**: Secure registration and login with JWT tokens
- **Password Security**: bcrypt hashing for secure password storage
- **Habit Management**: Create, read, update, and delete personal habits
- **Daily Check-ins**: Mark habits as completed for any date
- **Automatic Streak Tracking**: Real-time calculation of current and longest streaks
- **User-Specific Data**: Each user has their own private habit collection
- **Date-Based Logging**: Track habit completion history with timestamps
- **SQLite Database**: Lightweight, embedded database with zero configuration
- **Database Migrations**: Version-controlled database schema changes
- **Hot Reloading**: Air integration for development
- **GraphQL Playground**: Interactive API explorer included

## Technologies Used

- **Go**: Backend programming language
- **Gin**: HTTP web framework
- **gqlgen**: GraphQL server library for Go
- **SQLite**: Embedded relational database
- **database/sql**: Go standard library database interface
- **modernc.org/sqlite**: Pure-Go SQLite driver (no CGO/GCC required)
- **JWT**: JSON Web Tokens for authentication
- **bcrypt**: Password hashing
- **golang-migrate**: Database migrations
- **Air**: Hot reloading for development
- **godotenv**: Environment variable management

## Project Structure

```
habit-streak-tracker/
├── cmd/
│   └── api/
│       └── main.go                      # Application entry point
├── data/
│   └── habits.db                        # SQLite database (auto-created)
├── guides/
│   ├── Habit-Streak-Tracker-Setup-Guide.md
│   ├── Install-Sqlite3-Windows.md
│   └── Migrate-Sqlite-Driver.md
├── internal/
│   ├── database/
│   │   └── sqlite.go                    # SQLite connection
│   ├── graph/
│   │   ├── generated/
│   │   │   └── generated.go             # gqlgen generated code (do not edit)
│   │   ├── model/
│   │   │   └── models_gen.go            # gqlgen generated models
│   │   ├── resolvers/
│   │   │   ├── habit_log.go             # HabitLog field resolvers
│   │   │   ├── habit.go                 # Habit field resolvers
│   │   │   ├── mutation.go              # Mutation resolvers
│   │   │   ├── query.go                 # Query resolvers
│   │   │   ├── resolver.go              # Root resolver (dependency injection)
│   │   │   ├── schema.resolvers.go      # gqlgen resolver wiring
│   │   │   └── user.go                  # User field resolvers
│   │   └── schema.graphqls              # GraphQL schema definition
│   ├── middleware/
│   │   └── auth_middleware.go           # JWT authentication middleware
│   ├── models/
│   │   ├── habit_log.go                 # HabitLog model
│   │   ├── habit.go                     # Habit model
│   │   └── user.go                      # User model
│   ├── repository/
│   │   ├── habit_log_repository.go      # HabitLog database operations
│   │   ├── habit_repository.go          # Habit database operations
│   │   └── user_repository.go           # User database operations
│   └── utils/
│       ├── description.go               # Description validation
│       ├── email.go                     # Email validation
│       ├── jwt.go                       # JWT generation & validation
│       ├── name.go                      # Name validation
│       ├── password.go                  # bcrypt hashing & password strength
│       └── streak.go                    # Streak calculation logic
├── migrations/
│   ├── 000001_create_users_table.down.sql
│   ├── 000001_create_users_table.up.sql
│   ├── 000002_create_habits_table.down.sql
│   ├── 000002_create_habits_table.up.sql
│   ├── 000003_create_habit_logs_table.down.sql
│   └── 000003_create_habit_logs_table.up.sql
├── scripts/
│   └── migrate.ps1                      # Migration helper script
├── tmp/
│   └── main.exe                         # Air build output
├── .air.toml                            # Air configuration
├── .env                                 # Environment variables (create this)
├── .gitignore
├── go.mod                               # Go module definition
├── go.sum                               # Go dependencies checksum
├── gqlgen.yml                           # gqlgen configuration
├── README.md
└── tools.go                             # gqlgen tools build tag
```
