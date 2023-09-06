# Benchmarking I/O-Bound Servers: Node.js vs. Go in Docker

Benchmarking is a crucial step in ensuring the performance and scalability of your server applications. In this blog post, we will explore how to set up and benchmark I/O-bound servers in Node.js and Go within Docker containers. We'll compare the performance of these two popular languages for I/O-bound use cases.

## Prerequisites

Before we begin, make sure you have the following prerequisites installed:

- **Docker**: Download and install Docker from [here](https://www.docker.com/get-started).
- **Node.js and npm**: Install Node.js and npm from [here](https://nodejs.org/).
- **Go**: Install Go from [here](https://golang.org/dl/).
- **Apache Benchmark (ab)**: Install ab using your system's package manager or download it from [here](https://httpd.apache.org/docs/2.4/programs/ab.html).

## Node.js Setup

### I/O-Bound Use Case (HTTP Server)

We'll create HTTP server implementation using Express.js:

#### Express.js Implementation (app1.js):

```javascript
const express = require('express');
const app = express();

app.get('/', (req, res) => {
  // Simulate I/O-bound operation (e.g., reading from a file or database)
  setTimeout(() => {
    res.send('Express Server Response');
  }, 100);
});

app.listen(3000, () => {
  console.log('Express Server listening on port 3000');
});
```

## Go Setup

### I/O-Bound Use Case (HTTP Server)

We'll create HTTP server implementations using the Go standard library

#### Standard Go Library Implementation (main.go):

```go
package main

import (
  "fmt"
  "net/http"
  "time"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // Simulate I/O-bound operation (e.g., reading from a file or database)
    time.Sleep(100 * time.Millisecond)
    fmt.Fprintf(w, "Go Standard Server Response")
  })

  fmt.Println("Go Standard Server listening on :3000")
  http.ListenAndServe(":3000", nil)
}
```

### Docker Containerization

We'll use Docker to encapsulate these server implementations, ensuring a controlled environment for benchmarking. Follow these steps for both Node.js and Go servers

Create docker files for both - Refer in codebase.

#### Build the Docker image for Node server
```
cd node
docker build -t node-server -f Dockerfile .
```

#### Build the Docker image for Standard Go server
```
cd go
docker build -t go-server -f Dockerfile .
```


#### Create and run Docker containers for both servers:
```
docker run -d -p 3000:3000 --name go-container node-server

docker run -d -p 3001:3001 --name node-container go-server
```

### Benchmarking
With the Docker containers up and running, we can now benchmark the server implementations. Use Apache Benchmark (ab) or your preferred benchmarking tool:

```
ab -n 1000 -c 10 http://localhost:3000/
ab -n 1000 -c 10 http://localhost:3001/
```