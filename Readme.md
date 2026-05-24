# 🎮 Game App — Backend (Go)

Layered Go backend (`delivery` → `service` → `repository`) for a game platform.  
Current scope: JWT auth, user management, request validation, and **MySQL** persistence.

---

## 🏗 Architecture

```text
        🌐 Client
            │
            ▼
   ┌────────────────────┐
   │  Echo (delivery)   │  ← HTTP + middleware
   └─────────┬──────────┘
             │
             ▼
   ┌────────────────────┐
   │  User / Auth Svc   │  ← business logic
   └─────────┬──────────┘
             │
      ┌──────┴──────┐
      ▼             ▼
 ┌─────────┐   ┌──────────┐
 │  JWT    │   │  MySQL   │  ← tokens + data
 └─────────┘   └──────────┘
```

| Layer | Path | Role |
|-------|------|------|
| Delivery | `delivery/httpserver` | HTTP handlers, JSON responses |
| Service | `service/*` | Register, login, profile, JWT |
| Repository | `repository/mysql` | Queries and migrations |
| Entity / Param | `entity`, `param` | Domain models and DTOs |

---

## ⚙️ Configuration

Defaults live in `main.go`. Override via `config.yml` or env vars prefixed with `GAMEAPP_` (loaded with `koanf`).

| Setting | Default | Description |
|---------|---------|-------------|
| `http_server.port` | `8088` | HTTP server port |
| `mysql.host` | `127.0.0.1` | MySQL host |
| `mysql.port` | `3306` | MySQL port |
| `mysql.username` | `root` | Database user |
| `mysql.password` | *(empty)* | Database password |
| `mysql.dbname` | `gameapp_db` | Database name |
| `auth.sign_key` | `jwt_secret_key` | JWT signing key |
| `auth.access_subject` | `ac` | Access token subject |
| `auth.refresh_subject` | `rf` | Refresh token subject |
| Access token TTL | `24h` | Access token lifetime |
| Refresh token TTL | `7d` | Refresh token lifetime |

**`config.yml` example:**

```yaml
auth:
  sign_key: hello
mysql:
  host: localhost
  port: 3308
  dbname: gameapp_db
```

**Env example:** `GAMEAPP_MYSQL_PASSWORD=secret`

---

## 🗄️ Database

- **Engine:** MySQL 8
- **Migrations:** `repository/mysql/migrations` (`sql-migrate` or automatic via `migrator.Up()` in `main`)

| Table | Purpose |
|-------|---------|
| `users` | Users (name, phone, password, role) |
| `permissions` | Permission definitions |
| `access_controls` | Role/user ↔ permission mapping |

**Docker (MySQL):**

```bash
docker compose up -d
```

| Docker variable | Value |
|-----------------|-------|
| `MYSQL_DATABASE` | `gameapp_db` |
| `MYSQL_USER` | `gameapp` |
| Port | `3306` |

**Manual migration:**

```bash
go install github.com/rubenv/sql-migrate/...@latest
sql-migrate up -env="production" -config="repository/mysql/dbconfig.yml"
```

---

## 🚀 Quick start

```bash
git clone https://hamgit.ir/mm.gov.1381/game-app.git
cd game-app
go mod download
go run main.go
```

Requirements: **Go 1.25+** · **MySQL** (or `docker compose`)

```text
http://localhost:8088
```
