Task Management System


This is a sample Go gRPC application that allows you to manage tasks. The application consists of a server and a client that communicate using the gRPC protocol.  The main comopnents of this application include: client.go, server.go, and tasks.proto

**client.go:** This file defines the client code for the gRPC task management system. It imports the protocol buffer package and uses the grpc.Dial() function to connect to the server running on port 50051. It creates a client for the TaskService using the connection and sends a new task to the server using the AddTask RPC method. It also sends a request to the server using the GetTasks RPC method and logs the response from the server.

**server.go:** This file defines the server code for the gRPC task management system. It imports the protocol buffer package and defines a new struct for the task server that embeds the UnimplementedTaskServiceServer struct from the generated package. It defines two methods for the task server: AddTask and GetTasks. AddTask logs the task received from the client and returns a response indicating success. GetTasks logs the limit received from the client, creates a slice of task pointers, and returns the slice of tasks in a TaskList message. It also listens for incoming connections on port 50051, registers the task server with the gRPC server, and serves incoming requests.

**tasks.proto:** This file defines the message types and gRPC service for the task management system using Protocol Buffers version 3. It specifies the package name for the generated Go code and the Go package name for the generated code. It defines a new gRPC service called TaskService with two RPC methods: AddTask and GetTasks. It defines four message types: Task, TaskResponse, TaskRequest, and TaskList. Task has three string fields: id, name, and description. TaskResponse has one string field: message. TaskRequest has one integer field: limit. TaskList has one repeated field: tasks, which is a slice of Task messages.

Prerequisites
Before you can run the application, you will need to have the following installed on your system:

Go
protoc compiler
Installation
Clone this repository to your local machine:

bash

git clone https://github.com/your-username/task-management-system.git

Navigate to the project directory:


cd task-management-system/example.com/task-management-system/

**Generate the Go code for the Protocol Buffers file:**
protoc --go_out=plugins=grpc:. proto/tasks.proto

**Build the server binary:**
go build -o server server.go

**Build the client binary:**
go build -o client client.go


Usage

**Start the server:**
./server

This will start the server and listen for incoming gRPC requests on port 50051.

Open another terminal window and navigate to the same directory.

**Run the client:**
./client

This will run the client code and make gRPC requests to the server. The server will log the requests and respond accordingly.

**Contributing**
If you would like to contribute to this project, please fork the repository and submit a pull request.



