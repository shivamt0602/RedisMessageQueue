RedisMessageQueue

A simple example demonstrating message queuing with Redis:

Spring Boot publishes messages to a Redis queue (messageQueue)

Go (Fiber framework) consumes messages from the Redis queue with multiple workers

This project demonstrates a full producer-consumer setup using Redis as a message broker.

Features

Spring Boot API to publish messages

Redis queue for message storage (FIFO)

Go workers using BRPOP for consuming messages

Multiple concurrent Go workers for scalable message processing

Project Structure
msg_queue/
├── go.mod              # Go module file
├── go.sum
├── main.go             # Go entry point
├── routes/             # Fiber routes
├── handlers/           # Route handlers
├── workers/            # Worker functions consuming Redis queue
├── config/             # Redis configuration for Go
├── spring-boot/        # Spring Boot project (publisher)
└── README.md
Prerequisites

Go 1.21+

Spring Boot 3+

Redis
 running locally on localhost:6379

Git

Setup Instructions
1️⃣ Clone the repository
git clone git@github.com:shivamt0602/RedisMessageQueue.git
cd RedisMessageQueue
2️⃣ Start Redis
redis-server
3️⃣ Start Spring Boot publisher
cd spring-boot
./mvnw spring-boot:run

Publishes messages via POST request:

POST http://localhost:8080/addData
Content-Type: application/json

{
  "message": "Hello from Spring Boot"
}

Messages are pushed to the Redis queue messageQueue.

4️⃣ Start Go worker service
cd msg_queue
go run main.go

Starts multiple workers consuming messages from Redis

Workers print logs like:

Worker 1 processing: Hello from Spring Boot
Worker 2 processing: Another message
How it Works

Spring Boot RedisTemplate pushes messages using LPUSH.

Go workers block on BRPOP to consume messages.

Multiple workers compete for messages — each message is processed only once.

Queue preserves FIFO order: the oldest message is processed first.

Example Flow
[Spring Boot] POST /addData → "msg1"
[Spring Boot] POST /addData → "msg2"

[Redis Queue] (left → right): [msg2, msg1]

[Go Worker Logs]
Worker 1 processing: msg1
Worker 2 processing: msg2
Notes

Old keys in Redis may appear in Java-serialized format if Spring Boot was used without StringRedisSerializer.

Only newly published messages after configuring StringRedisSerializer are human-readable.

Multiple workers improve throughput but exact assignment of messages to workers is non-deterministic.
