package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	desc "github.com/xeeetu/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

// Get Получает пользователя по id
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	select {
	case <-ctx.Done():
		// Если контекст отменён, возвращаем ошибку
		return nil, ctx.Err()
	default:
		// Продолжаем выполнение
	}
	fmt.Println("Get user with id: ", req.GetId())

	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.TypeUser_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

// Create Создаёт пользователя
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	select {
	case <-ctx.Done():
		// Если контекст отменён, возвращаем ошибку
		return nil, ctx.Err()
	default:
		// Продолжаем выполнение
	}
	fmt.Printf("Create user: %v", req)

	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

// Update Обновляет данные пользователя (можем частично)
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	select {
	case <-ctx.Done():
		// Если контекст отменён, возвращаем ошибку
		return nil, ctx.Err()
	default:
		// Продолжаем выполнение
	}
	fmt.Printf("Update user: %v", req)
	return &emptypb.Empty{}, nil
}

// Delete Удаляет пользователя по id
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	select {
	case <-ctx.Done():
		// Если контекст отменён, возвращаем ошибку
		return nil, ctx.Err()
	default:
		// Продолжаем выполнение
	}
	fmt.Printf("Delete user with id: %v", req.GetId())
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
