# Symptom Tracker

A simple mobile app to help you track your health symptoms.

## Setup
### Postgres
Make sure you install the [Postgres driver](https://github.com/lib/pq):
```
go get github.com/lib/pq
```

## Using the App
### In Dev Mode
To run the app:
```
go run cmd/web/*
```

### Command Line Flags
The following command line flags are available for setting configurations at runtime:
- addr
- dsn (data source name for psql)
- html-dir
- static-dir

If not set, default values will be used. Use `-help` to get a list of all available commands.

Example of usage:
```
go run cmd/web/* -addr=":80" -static-dir="/public" -html-dir="/templates"
```