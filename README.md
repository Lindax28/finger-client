
# Finger Client

This client uses the [Finger user information protocol](https://datatracker.ietf.org/doc/html/rfc1288) to open a TCP connection that makes a request to a Finger server.

## Build and Run

1. Run `go build -o finger`
2. Run `./finger <username>@hostname` (e.g., `./finger lindax@happynetbox.com`)

> To create your own account to test, make a free profile at [happynetbox.com](https://happynetbox.com/)