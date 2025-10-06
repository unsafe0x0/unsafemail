# Unsafemail (Go Project)

A tiny Go HTTP service that exposes a single endpoint to send an email using SMTP.

This project is intentionally minimal and demonstrates how to wire a small HTTP handler to Go's standard library SMTP client.

## Project layout

- `main.go` - sets up an HTTP server and registers the `/send-email` route.
- `api/handler.go` - contains the HTTP handler which accepts a JSON payload and calls the email sender.
- `email/email.go` - contains the `Send` function which performs SMTP delivery using configuration from the `config` package.
- `config/config.go` - reads SMTP-related settings from environment variables (and provides sane defaults for host/port).
- `config/config.go` - loads environment variables (optionally from a `.env` file) and validates required SMTP settings.
- `go.mod` - Go module file.

## Contract

- Input: POST /send-email with JSON body: `{ "to": "recipient@example.com", "subject": "...", "body": "<b>HTML content here</b>" }`
	- The `body` field should contain the HTML content you want to send in the email. It will be sent as HTML (MIME type `text/html`).
- Output: 200 OK with body `Email sent successfully` on success. Appropriate HTTP error codes on failure.
	- The response is always HTTP 200 with a plain text message on success.
- Error modes: malformed JSON -> 400; non-POST -> 405; SMTP failure -> 500.

## Environment and configuration

The application loads environment variables at startup using the `config.Init()` function. `config.Init()` attempts to load a local `.env` file (via `github.com/joho/godotenv`) and then reads values from the process environment. If required values are missing it will exit the process.

Example `.env` file (place at the repository root):

```text
EMAIL_FROM=your-email@example.com
EMAIL_PASSWORD=your-email-password-or-app-password
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

## Build and run

1. Run locally:

```bash
go run main.go
```

The server listens on `:8080` and exposes the `POST /send-email` endpoint. `config.Init()` is executed on startup and will log an error and exit if required environment values are missing.


## Example request

Use curl to send an example request to the running server (with HTML body):

```bash
curl -X POST http://localhost:8080/send-email \
	-H "Content-Type: application/json" \
	-d '{"to":"recipient@example.com","subject":"Hello","body":"<b>This is a test.</b>"}'
```

If successful you should receive `Email sent successfully` as the response body (HTTP 200).


## License

No license file included. Add an appropriate LICENSE if you plan to share the repository publicly.
