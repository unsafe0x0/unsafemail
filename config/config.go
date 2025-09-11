package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var (
    From     string
    Password string
    SmtpHost string
    SmtpPort string
)

func Init() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment")
    }

    From = os.Getenv("EMAIL_FROM")
    Password = os.Getenv("EMAIL_PASSWORD")
    SmtpHost = os.Getenv("SMTP_HOST")
    SmtpPort = os.Getenv("SMTP_PORT")

    if From == "" || Password == "" || SmtpHost == "" || SmtpPort == "" {
        log.Fatal("SMTP credentials are not properly set in .env or environment")
    }
}
