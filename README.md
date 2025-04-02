# Blogator

A CLI tool for aggregating and browsing RSS feeds.

## Prerequisites

- Go 1.22.6+ (for installation only)
- PostgreSQL

## Installation

Install the CLI directly using Go:

```bash
go install github.com/FT1006/blogator@latest
```

After installation, the `blogator` binary will be available in your PATH.

## Configuration

Create a configuration file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/blogator?sslmode=disable"
}
```

Make sure to replace the database credentials with your own.

## Database Setup

1. Create a PostgreSQL database
2. The application uses [Goose](https://github.com/pressly/goose) for migrations. The schema files are in `sql/schema/`

## Commands

- `blogator register [username]` - Create a new user account
- `blogator login [username]` - Log in to your account
- `blogator reset` - Reset database/config
- `blogator users` - List all users
- `blogator addfeed [name] [url]` - Add a new RSS feed (requires login)
- `blogator feeds` - List all available feeds
- `blogator follow [feed_id]` - Follow a feed (requires login)
- `blogator following` - List feeds you're following (requires login)
- `blogator unfollow [feed_id]` - Unfollow a feed (requires login)
- `blogator agg [time duration(optional)]` - Aggregate feeds, time duration is optional and defaults to "30m"
- `blogator browse [limit(optional)]` - Browse posts, limit is optional and defaults to 2

## Development

This project uses:
- [SQLC](https://github.com/kyleconroy/sqlc) for type-safe SQL queries
- [Goose](https://github.com/pressly/goose) for database migrations