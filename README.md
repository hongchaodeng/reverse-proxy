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
export PROXY_ROUTES="/api:https://catfact.ninja/fact,/todo:https://jsonplaceholder.typicode.com/todos/1"
make run
```

Now you can open your browser and visit http://localhost:8080/api and http://localhost:8080/todo to see the result.

## What resources used to build this implementation?

- Golang
- Go standard library
- VSCode IDE

## What design decisions I made, including limitations of the system?

There are a couple of design decisions I made in this project and I will explain them in the following.

- **No support for WebSocket or other non-http protocol**: The program only supports http protocol. Usually a reverse
  proxy supports other protocols (e.g. WebSocket, gRPC, etc.) to provide more functionalities.
- **No security measures**: The program does not provide any security measures (e.g. TLS encryption, authentication,
  authorization, etc.). Usually a reverse proxy provides security measures to protect the backend.
- **No error handling**: The program does not handle errors. Usually a reverse proxy handles errors (e.g. retry, circuit
  breaker, etc.) to provide more robustness.
- **No middleware support**: The program does not support middleware. Usually a reverse proxy supports middleware
  (e.g. authentication, authorization, rate limiting, etc.) to provide more functionalities.

Note that the above limitations are not hard to overcome. We can easily add them to the program given more time and effort.

## How to scale this?

This program is stateless.
To scale it, we can deploy multiple instances of it behind a load balancer.

## How to make it more secure?

We can add the following security measures to make it more secure:

- TLS encryption: We can use TLS to encrypt the communication between the client and the reverse proxy.
- Authentication: We can use authentication to authenticate the client to avoid unauthorized access.
- Authorization: We can use policy-based authorization to decide which client can access which backends.
