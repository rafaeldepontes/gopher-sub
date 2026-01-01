# Gopher-Sub

A minimal full-stack subscription app using Go for the backend API and Next.js for the UI. Users can register, authenticate, subscribe to plans, and receive server-sent email notifications after completion.

## Requirements

in order to run this application you need:

- Go 1.25.5
- Next 16
- Node.Js 20
- Docker

## How to Run

Each directory has their own way to start the application, so I recommend you to take a look at it first.

But if you don't want to, from the root, you can:

1. Database and Email:

```bash
docker compose up # if you don't have the docker desktop
docker compose up -d
```

2. Backend:

```bash
cd ./backend
go mod tidy
go run main.go
```

3. Frontend:

```bash
cd ./frontend
npm run dev
```

## Contact

If anything goes wrong, please contact: `rafael.cr.carneiro@gmail.com`
