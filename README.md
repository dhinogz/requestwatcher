# RequestWatcher
A tool to inspect incoming HTTP requests, like webhooks, quickly and easily. 

![](https://github.com/dhinogz/requestwatcher/blob/main/docs/demo.gif)

## Features
- Responsive web interface
- Create a new HTTP watcher and a link would be generated
- Inspect incoming requests
- SPA like beheviour

## Tech
- Go
  - Chi
  - Templ Components
- PostgreSQL
- SQLC
- HTMX
- Server-Sent Events

## What's left to improve
- Install tailwind instead of CDN use
- Fix table overflow for smaller screens
- Make tests 
  - Use [testcontainers](https://golang.testcontainers.org/) for DB tests
  - Read up on how to test channels for Manager package
- Move packages to internal/ directory
- Consider changing to NoSQL solution for storage
  - We don't have many tables 
  - the Headers type is a map[string][]string
  - We only query all requests from one ID and add a request to one ID
  - Filestorage could be cool. Could then be shipped as a standalone Go binary
- Make body formatting pretty
- Develop a CLI using Cobra and Charm


## Getting Started
### Prerequisites
- [Go](https://go.dev/)
- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [air](https://github.com/cosmtrek/air) (for hot reload)
- [sqlc](https://sqlc.dev/)

Copy .env.example file (modify to your own liking)
```bash
$ cp .env.example .env
```

### Hot Reload
Hot reloads Go and Templ files using air
```bash
$ make dev
```

### Build
Build Go application
```bash
$ make build
```

Start PostgreSQL database
```bash
$ make db/up
```

Migrate SQL schemas
```bash
$ make db/migrate
```

Run
```bash
$ make run
```

