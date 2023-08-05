# Golang Mobile OTP based Auth

## Tech stack
- golang
- gofiber
- mongodb
- twilio sdk

## Project structure

```md
Project
├── README.md
├── config
│   └── config.go
├── database
│   ├── connection.go
│   └── database.go
├── go.mod
├── go.sum
├── handler
│   └── auth.go
├── main.go
├── middleware
│   └── auth.go
├── model
│   └── user.go
├── router
│   └── router.go
├── schema
│   ├── auth.go
│   └── response.go
└── util
    ├── twilio.go
    └── user.go
```


## Routes

- /api/auth
    - /register (create new account)
    - /login (sent otp to registered mobile number)
    - /verify_otp
    - /resend_otp
    - /me (get current loggedin user)
