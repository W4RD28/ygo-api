## YGO-API

This is my Go side project for making a Yu-Gi-Oh! card API.

## How to Run

You need the following
1. Go 1.22
2. PostgreSQL
3. MinIO

First you should change the `.env.example` file to a `.env` file and change the variables. Then, you could run the app with this command
```
go run ./cmd/ygo-api/main.go
```

Or, if you installed Air
```
air
```
