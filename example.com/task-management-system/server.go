package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"

    pb "example.com/task-management-system/proto"
)

type taskServer struct {
    pb.UnimplementedTaskServiceServer
}

func (s *taskServer) AddTask(ctx context.Context, task *pb.Task) (*pb.TaskResponse, error) {
    log.Printf("Received task: %v", task)
    return &pb.TaskResponse{Message: "Task added successfully."}, nil
}

func (s *taskServer) GetTasks(ctx context.Context, req *pb.TaskRequest) (*pb.TaskList, error) {
    log.Printf("Received task request with limit: %v", req.Limit)
    tasks := []*pb.Task{
        {Id: "1", Name: "Task 1", Description: "Task 1 Description"},
        {Id: "2", Name: "Task 2", Description: "Task 2 Description"},
        {Id: "3", Name: "Task 3", Description: "Task 3 Description"},
    }
    return &pb.TaskList{Tasks: tasks}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterTaskServiceServer(s, &taskServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

