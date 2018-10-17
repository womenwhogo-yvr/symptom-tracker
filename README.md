# Symptom Tracker

This web app helps you track your health symptoms.

## Using the App
### Command Line Flags
The following command line flags are available for setting configurations at runtime:
- addr
- html-dir
- static-dir

If not set, default values will be used. Use `-help` to get a list of all available commands.

Example of usage:
```
go run cmd/web/* -addr=":80" -static-dir="/public" -html-dir="/templates"
```