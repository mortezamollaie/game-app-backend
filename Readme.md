# Quiz Game Backend (Go)

Simple backend for a quiz competition game written in Go using only the standard library (no external frameworks).

In this game, each user joins a match and receives a set of questions. The user answers them and the system calculates the result. At the end of a match the player can win, lose, or draw. Players also earn points and are ranked in a leaderboard.

## Features

- Quiz matches for users
- Question answering system
- Match result calculation (Win / Lose / Draw)
- Player scoring
- Leaderboard ranking
- Built with Go standard library (net/http)

## Requirements

- Go 1.21 or newer

## Run the Project

Clone the repository:

```
https://hamgit.ir/mm.gov.1381/game-app.git
cd game-app
```

Run the server:

```
go run main.go
```

The server will start on:

```
http://localhost:8080
```

## Build Binary

```
go build -o server
./server
```

## Migrations
```bash
go install github.com/rubenv/sql-migrate/...@latest
sql-migrate up -env="production" -config="repository/mysql/dbconfig.yml"
sql-migrate down -env="production" -config="repository/mysql/dbconfig.yml" -limit=1
sql-migrate status -env="production" -config="repository/mysql/dbconfig.yml"
```