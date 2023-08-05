# Golang Mobile OTP based Auth

## Tech stack
- golang
- gofiber
- mongodb
- twilio sdk

## Project structure
- Project Root
    - config (configuration file i.e env variable configuration)
    - database (mongodb driver setup)
    - handler (All routes handler i.e handle registration, login etc...)
    - middleware (All middlewares i.e route auth protection middleware)
    - model (Mongodb model i.e User model)
    - router (route initialization and all routes i.e. auth related routes)
    - schema (application schema for incomming request body)
    - util (database and application helper function )
    - .env (for environment variables i.e. twillio apikey etc...)
    - .env.example (copy all env variable from here and paste to .env)
    - main.go (project entry point)

## Routes

- /api/auth
    - /register (create new account)
    - /login (sent otp to registered mobile number)
    - /verify_otp
    - /resend_otp
    - /me (get current loggedin user)
