# Gator - Blog Aggregator CLI

A command-line RSS feed aggregator built with Go and PostgreSQL.

## Prerequisites

- [Go](https://golang.org/dl/) (1.23+)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

```bash
go install github.com/jpanderson91/blog-aggregator@latest
```

This installs the `blog-aggregator` binary to your `$GOPATH/bin`.

## Configuration

Create a `.gatorconfig.json` file in your home directory (`~/.gatorconfig.json`):

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

Adjust the connection string to match your PostgreSQL setup.

### Database Setup

Create the database:

```bash
createdb gator
```

Run migrations (requires [goose](https://github.com/pressly/goose)):

```bash
goose -dir sql/schema postgres "your_connection_string" up
```

## Usage

### Register a user

```bash
blog-aggregator register <name>
```

### Login

```bash
blog-aggregator login <name>
```

### Add a feed

```bash
blog-aggregator addfeed <name> <url>
```

### List all feeds

```bash
blog-aggregator feeds
```

### Follow / Unfollow a feed

```bash
blog-aggregator follow <url>
blog-aggregator unfollow <url>
```

### See feeds you're following

```bash
blog-aggregator following
```

### Start the aggregator

```bash
blog-aggregator agg <duration>
```

Example: `blog-aggregator agg 1m` fetches feeds every minute.

### Browse posts

```bash
blog-aggregator browse [limit]
```

Default limit is 2 if not specified.

### Reset the database

```bash
blog-aggregator reset
```

### List all users

```bash
blog-aggregator users
```
