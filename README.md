# Rabbit Server


A dedicated server implementation for [Ring Around the Rabbit](https://insertapp.itch.io/ring-around-the-rabbit) written in go

# Environment

To set the sqlite database location use the env variable `DB_PATH`

# Building

To build the package either use `go build` (CGO required), the frontend/templates and assets folder must be included in the same directory as the compiled binary.

Or the provided dockerfile `docker build -t rabbitserver .`