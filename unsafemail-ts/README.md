# Unsafemail (TypeScript/Bun Project)

A simple and fast email sending service built with TypeScript, Express, and Nodemailer, running on Bun. It exposes a single API endpoint to send emails with both plain text and HTML content.

## Features

-   **Fast and Efficient**: Leverages Bun runtime for high performance.
-   **Flexible Email Content**: Supports sending emails with plain text and/or HTML bodies.
-   **CORS Enabled**: Configurable CORS for secure cross origin requests.
-   **Helmet Security**: Basic security headers provided by `helmet`.

## Project Layout

-   `index.ts`: Main application file, sets up the Express server, middleware, and routes.
-   `utils/Route.ts`: Defines the `/api/send-email` POST endpoint for sending emails.
-   `utils/Mailer.ts`: Contains the Nodemailer transporter and `sendEmail` function for SMTP communication.
-   `.env`: Environment variables for configuration (SMTP credentials, CORS origin, port).
-   `package.json`: Project metadata and dependencies.
-   `tsconfig.json`: TypeScript compiler configuration.

## Setup

### 1. Install Dependencies

Make sure you have [Bun](https://bun.sh/) installed. Then, install the project dependencies:

```bash
bun install
```

### 2. Environment Variables

Create a `.env` file in the project root and add the following environment variables:

```text
PORT=3000
CORS_ORIGIN=*
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_SECURE=false
SMTP_USER=your-email@example.com
SMTP_PASS=your-email-password
SMTP_FROM=your-email@example.com
```

-   `PORT`: The port the Express server will listen on (default: `3000`).
-   `CORS_ORIGIN`: The allowed origin for CORS requests (e.g., `http://localhost:3000` or `*` for all origins).
-   `SMTP_HOST`: Your SMTP server host.
-   `SMTP_PORT`: Your SMTP server port (e.g., `587` for TLS, `465` for SSL).
-   `SMTP_SECURE`: Set to `true` if your SMTP server uses SSL (port 465), `false` otherwise (port 587 with TLS).
-   `SMTP_USER`: Your SMTP username (usually your email address).
-   `SMTP_PASS`: Your SMTP password or app specific password.
-   `SMTP_FROM`: The email address from which emails will be sent.

## Running the Application

To start the development server:

```bash
bun run index.ts
```

The server will be running on the configured `PORT` (default: `3000`).

## API Endpoint

### `POST /api/send-email`

Sends an email using the configured SMTP server.

#### Request Body (JSON)

```json
{
    "to": "recipient@example.com",
    "subject": "Hello from Bun!",
    "html": "<b>This is the HTML body.</b>"
}
```

-   `to` (string, required): The recipient's email address.
-   `subject` (string, required): The subject of the email.
-   `html` (string, required): The HTML content of the email.

#### Success Response (HTTP 200 OK)

```json
{
    "message": "Email sent successfully",
    "info": "250 2.0.0 OK  1678888888 example.com"
}
```

-   `message`: A success message.
-   `info`: The response from the SMTP server.

#### Error Response (HTTP 500 Internal Server Error)

```json
{
    "message": "Error sending email",
    "error": { /* error details */ }
}
```

-   `message`: An error message.
-   `error`: Details of the error that occurred during email sending.

## Example `curl` Request

```bash
curl -X POST http://localhost:3000/api/send-email \
    -H "Content-Type: application/json" \
    -d '{
        "to": "recipient@example.com",
        "subject": "Test Email from Bun",
        "html": "<b>This is a test email sent from the Bun API.</b>"
    }'
```
