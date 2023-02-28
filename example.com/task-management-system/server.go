package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"

    pb "example.com/task-management-system/proto" // import the generated protocol buffer package
)

type taskServer struct { // define a new struct for the task server
    pb.UnimplementedTaskServiceServer // embed the UnimplementedTaskServiceServer struct from the generated package
}

func (s *taskServer) AddTask(ctx context.Context, task *pb.Task) (*pb.TaskResponse, error) {
    log.Printf("Received task: %v", task) // log the task received from the client
    return &pb.TaskResponse{Message: "Task added successfully."}, nil // return a response indicating success
}

func (s *taskServer) GetTasks(ctx context.Context, req *pb.TaskRequest) (*pb.TaskList, error) {
    log.Printf("Received task request with limit: %v", req.Limit) // log the limit received from the client
    tasks := []*pb.Task{ // create a slice of task pointers
        {Id: "1", Name: "Task 1", Description: "Task 1 Description"}, // add task 1
        {Id: "2", Name: "Task 2", Description: "Task 2 Description"}, // add task 2
        {Id: "3", Name: "Task 3", Description: "Task 3 Description"}, // add task 3
    }
    return &pb.TaskList{Tasks: tasks}, nil // return the slice of tasks in a TaskList message
}

func main() {
    lis, err := net.Listen("tcp", ":50051") // listen for incoming connections on port 50051
    if err != nil {
        log.Fatalf("Failed to listen: %v", err) // handle any errors that occur while listening
    }
    s := grpc.NewServer() // create a new gRPC server
    pb.RegisterTaskServiceServer(s, &taskServer{}) // register the task server with the gRPC server
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err) // handle any errors that occur while serving
    }
}
