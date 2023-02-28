package main

import (
    "context"
    "log"

    "google.golang.org/grpc"

    pb "example.com/task-management-system/proto"
)

func main() {
    conn, err := grpc.Dial(":50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to dial server: %v", err)
    }
    defer conn.Close()

    c := pb.NewTaskServiceClient(conn)

    // AddTask
    task := &pb.Task{Id: "1", Name: "Task 1", Description: "Task 1 Description"}
    res, err := c.AddTask(context.Background(), task)
    if err != nil {
        log.Fatalf("Failed to add task: %v", err)
    }
    log.Printf("Add task response: %v", res)

    // GetTasks
    req := &pb.TaskRequest{Limit: 3}
    res2, err := c.GetTasks(context.Background(), req)
    if err != nil {
        log.Fatalf("Failed to get tasks: %v", err)
    }
    log.Printf("Get tasks response: %v", res2)
}

