package main

import (
    "context"
    "log"

    "google.golang.org/grpc"

    pb "example.com/task-management-system/proto" // import the generated protocol buffer package
)

func main() {
    // Dial the server using the gRPC framework and port 50051, which is where the server is listening
    conn, err := grpc.Dial(":50051", grpc.WithInsecure()) // WithInsecure is used to skip transport security
    if err != nil {
        log.Fatalf("Failed to dial server: %v", err) // handle any errors that occur while dialing
    }
    defer conn.Close() // ensure that the connection is closed when the main function exits

    // Create a new client for the TaskService using the connection
    c := pb.NewTaskServiceClient(conn)

    // AddTask
    task := &pb.Task{Id: "1", Name: "Task 1", Description: "Task 1 Description"} // create a new task
    res, err := c.AddTask(context.Background(), task) // send the task to the server using the AddTask RPC
    if err != nil {
        log.Fatalf("Failed to add task: %v", err) // handle any errors that occur while sending the task
    }
    log.Printf("Add task response: %v", res) // log the response from the server

    // GetTasks
    req := &pb.TaskRequest{Limit: 3} // create a new task request with a limit of 3
    res2, err := c.GetTasks(context.Background(), req) // send the request to the server using the GetTasks RPC
    if err != nil {
        log.Fatalf("Failed to get tasks: %v", err) // handle any errors that occur while sending the request
    }
    log.Printf("Get tasks response: %v", res2) // log the response from the server
}
