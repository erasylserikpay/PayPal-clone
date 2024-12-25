package config

import (
    "os"
)

func GetEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

var (
    SMTPHost     = GetEnv("SMTP_HOST", "smtp.yourprovider.com")
    SMTPPort     = GetEnv("SMTP_PORT", "587")
    SMTPUsername = GetEnv("SMTP_USERNAME", "your-smtp-username")
    SMTPPassword = GetEnv("SMTP_PASSWORD", "your-smtp-password")
    EmailFrom    = GetEnv("EMAIL_FROM", "your-email@example.com")
)

