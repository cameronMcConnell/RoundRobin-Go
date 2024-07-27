# RoundRobin-Go

RoundRobin-Go is a lightweight Go middleware that implements the round-robin algorithm for load balancing HTTP requests across multiple backend servers.

## Features

- Simple and efficient round-robin load balancing
- Concurrent request handling
- Easy configuration via included JSON and environment files
- Docker support for easy deployment

## Getting Started

### Prerequisites

- Go 1.22 or later
- Git

### Clone the Repository

To get started with RoundRobin-Go, clone the repository to your local machine:

```bash
git clone https://github.com/CameronMcConnell/RoundRobin-Go.git
cd RoundRobin-Go
```

### Configuration
The project includes two pre-configured files:

1. `servers.json`: Contains the list of backend servers for load balancing.
2. `.env`: Specifies the server address for the middleware.

You can modify these files to suit your needs:

Edit `servers.json` to add or remove backend servers:

```json
[
  "http://backend1:8081",
  "http://backend2:8082",
  "http://backend3:8083"
]
```

Adjust the `.env` file to change the server address if needed:

```bash
SERVER_ADDRESS=:8080
```
## Running the Application

Run the application:

```bash
go run main.go
```

The server will start on the address specified in the `.env` file.

### Usage
Once the server is running, it will distribute incoming requests to the `/roundrobin-go` endpoint across the configured backend servers in a round-robin fashion.

### Docker Support
To build and run RoundRobin-Go using Docker:

1. Build the Docker image:

```bash
docker build -t roundrobin-go .
```

2. Run the container:
```bash
docker run -p 8080:8080 roundrobin-go
```
Note: The `servers.json` and `.env` files are included in the Docker image during the build process.

## Project Structure

* `main.go`: Entry point of the application
* `lib/`: Contains the core logic for the round-robin server
* `servers.json`: Configuration file for backend servers (included)
* `.env`: Environment file for server address configuration (included)

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.