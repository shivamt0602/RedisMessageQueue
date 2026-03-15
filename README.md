RedisMessageQueue

A simple producer-consumer demo using Redis:

Spring Boot publishes messages to a Redis queue (messageQueue)

Go (Fiber) consumes messages with multiple workers

Features

Spring Boot API to push messages

Redis queue (FIFO)

Go workers using BRPOP for message consumption

Multiple workers for scalable processing

Project Structure
msg_queue/
├── main.go
├── routes/
├── handlers/
├── workers/
├── config/
├── go.mod
├── go.sum
├── spring-boot/
└── README.md
Prerequisites

Go 1.21+

Spring Boot 3+

Redis running on localhost:6379

Git

Setup
1️⃣ Clone repo
git clone git@github.com:shivamt0602/RedisMessageQueue.git
cd RedisMessageQueue
2️⃣ Start Redis
redis-server
3️⃣ Start Spring Boot publisher
cd spring-boot
./mvnw spring-boot:run

Publish messages:

POST http://localhost:8080/addData
{
  "message": "Hello from Spring Boot"
}
4️⃣ Start Go worker service
cd msg_queue
go run main.go
