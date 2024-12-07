# 📰 Blog Aggregator

Welcome to the Blog Aggregator! This is a 🔧 CLI tool that allows 👥 users to:

- **➕ Add RSS feeds** 🔄 from across the 🌎 internet to be collected.
- **📂 Store collected posts** in a 💻 PostgreSQL database.
- **🔄 Follow and unfollow RSS feeds** that other 👥 users have added.
- **🔍 View summaries** of aggregated posts in the 📝 terminal, with a 🔗 link to the full post.

---

## 🔄 Features and 🔧 Technologies Used

### 📊 Goose Migrations
**Goose** is a 🔐 database migration 🔧 tool written in Go. It manages database schema changes using SQL files, which keeps us close to ✈ raw SQL while providing 🔄 flexibility and 🔒 control.

#### 🔄 What is a Migration?
A migration is a set of ✍️ changes to your 🔐 database schema. Examples include:
- ✅ Creating a new table.
- ❌ Deleting a column.
- ➕ Adding new columns.

Migrations consist of:
- **"⬆️ Up" migrations**: Move the database schema forward.
- **"⬇️ Down" migrations**: Revert the schema to a previous state.

#### 🔧 Installing Goose
Install Goose using:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
Verify installation:
```bash
goose -version
```

#### 🔧 Creating a Migration
Create a `users` migration in the `sql/schema` 🗁 directory. Each migration file follows this naming format:
```
<number>_<name>.sql
```
Example: `001_users.sql`:
```sql
-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE users;
```
> Note: `-- +goose Up` and `-- +goose Down` comments are **required** and case-sensitive.

#### 🔄 Running Migrations
1. Get your 🆔 connection string:
   ```
   protocol://username:password@host:port/database
   ```
   Examples:
   - 💺 MacOS: `postgres://artiomstartev:@localhost:5432/rssagg`
   - 🔧 Linux: `postgres://postgres:postgres@localhost:5432/rssagg`

2. Run the migration:
   ```bash
   goose postgres <connection_string> up
   ```
3. Verify:
   - Use `psql` to 🔍 check if the `users` 📂 table is created.
   - Test the down migration by running:

     ```bash
     goose postgres <connection_string> down
     ```
4. ♻️ Recreate the table with an up migration.

---

### 📊 SQLC
**SQLC** is a Go 🔧 tool that generates Go 📄 code from SQL 🔄 queries. It simplifies working with raw SQL by making it type-safe and easier to manage.

#### 🔧 Installing SQLC
Install SQLC using:
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

#### 🔧 Configuring SQLC
Create a `sqlc.yaml` 🗁 file in the root of your 🔄 project:
```yaml
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```
This configuration:
- Points SQLC to the `sql/schema` 🗁 directory for the schema structure (used by Goose).
- Points SQLC to the `sql/queries` 🗁 directory for SQL queries.
- Specifies that generated Go 📄 code will be placed in the `internal/database` directory.

#### 🔧 Writing Queries
Example query in `sql/queries/users.sql`:
```sql
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;
```
> `$1`, `$2`, `$3`, and `$4` are parameters that will be passed from Go 📄 code. `:one` specifies that the query returns a single row.

#### 🔧 Generating Go Code
Run the following 🔧 command from the root of your 🔄 project:
```bash
sqlc generate
```
SQLC will generate Go 📄 code in the `internal/database` directory.

---

## 🔍 Notes
- This project uses raw SQL for 🎡 better control and ⚖️ performance.
- Goose ensures smooth 🄐 database migrations, while SQLC integrates type-safe SQL into the 🔧 application.

---

## 🎡 Contributing
Feel free to ✉ open 🔒 issues or submit ➕ pull requests to improve the 🎡 project.

---

✨ Happy hacking! 🚀
