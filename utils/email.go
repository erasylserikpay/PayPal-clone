package utils

import (
    "gopkg.in/gomail.v2"
    "paypal-clone/config"
)

func SendEmail(to string, subject string, body string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", config.GetEnv("EMAIL_FROM", "your-email@example.com"))
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer(
        config.GetEnv("SMTP_HOST", "smtp.example.com"),
        587,
        config.GetEnv("SMTP_USERNAME", "your-smtp-username"),
        config.GetEnv("SMTP_PASSWORD", "your-smtp-password"),
    )

    if err := d.DialAndSend(m); err != nil {
        return err
    }
    return nil
}
