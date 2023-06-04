# reverse-proxy

This project provides a simple reverse proxy implementation in Go.

## How to get started with the codebase?

![](doc/demo.gif)

Clone the repository to your local machine:

```bash
git clone https://github.com/hongchaodeng/reverse-proxy.git
cd reverse-proxy
```

Make sure you have Go 1.19+ installed on your machine. If not, please follow the [official installation guide](https://go.dev/doc/install).

Build and run the project:

```bash
export BACKEND_ADDR="http://localhost:8000"
make run
```

The above is assuming you have a web server running on `http://localhost:8000`.
In the following we use [docker](https://docs.docker.com/engine/install/) to run a simple web server.
Note that you can use other methods (e.g. `python -m http.server 8000`) to achieve the same goal.

Open a new terminal, run the following command:

```bash
docker run -it -p 8000:80 docker/getting-started
```

Now you can open your browser and visit `http://localhost:8080` to see the result.

## What resources used to build this implementation?

- Golang
- Go standard library
- Makefile
- Docker
- VSCode IDE

## What design decisions I made, including limitations of the system?

There are a couple of design decisions I made in this project and
I will explain them in the following with the corresponding limitations.

- **Single backend only**: The program takes only a single backend. Usually a reverse proxy takes multiple backends and
  routes the requests to the backend based on some routing logic. However, in this project, I only take a single backend
  and route all the requests to it.
- **Limited routing logic**: The program routes all requests to a single backend without any routing configuration.
  Usually a reverse proxy routes requests to different backends based on some routing logic. However, in this project, I
  implement the simplest version without custom routing.
- **No middleware support**: The program does not support middleware. Usually a reverse proxy supports middleware
  (e.g. authentication, authorization, rate limiting, etc.) to provide more functionalities.
- **No error handling**: The program does not handle errors. Usually a reverse proxy handles errors (e.g. retry, circuit
  breaker, etc.) to provide more robustness.
- **No security measures**: The program does not provide any security measures (e.g. TLS encryption, authentication,
  authorization, etc.). Usually a reverse proxy provides security measures to protect the backend.
- **No support for websockets or other non-http protocol**: The program only supports http protocol. Usually a reverse
  proxy supports other protocols (e.g. websockets, gRPC, etc.) to provide more functionalities.

## How to scale this?

This program is stateless. To scale it, we can deploy multiple instances of it behind a load balancer.

## How to make it more secure?

We can add the following security measures to make it more secure:

- TLS encryption: We can use TLS to encrypt the communication between the client and the reverse proxy.
- Authentication: We can use authentication to authenticate the client to avoid unauthorized access.
- Authorization: We can use policy-based authorization to decide which client can access which backends.
