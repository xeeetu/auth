package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	desc "github.com/xeeetu/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if errCon := conn.Close(); errCon != nil {
			log.Fatal(err)
		}
	}()

	c := desc.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pass := gofakeit.Password(true, true, true, true, false, 16)

	r, err := c.Create(ctx, &desc.CreateRequest{
		Name:            gofakeit.Name(),
		Email:           gofakeit.Email(),
		Password:        pass,
		PasswordConfirm: pass,
		Role:            desc.TypeUser_USER,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("User created with id: ", r.GetId())

	rGet, err := c.Get(ctx, &desc.GetRequest{Id: gofakeit.Int64()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User retrieved : %v", rGet)

	_, err = c.Update(ctx, &desc.UpdateRequest{Id: gofakeit.Int64(), Name: wrapperspb.String(gofakeit.Name())})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Delete(ctx, &desc.DeleteRequest{Id: gofakeit.Int64()})
	if err != nil {
		log.Fatal(err)
	}
}
