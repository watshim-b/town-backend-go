package main

import (
	"fmt"
	"log"
	"net"

	"github.com/watshim-b/town-backend-go/handler"
	hd "github.com/watshim-b/town-backend-go/handler_definition"
	"google.golang.org/grpc"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":% d", port))

	if err != nil {
		log.Fatalf(" failed to listen: %v", err)
	}

	s := handler.NewLoginServer()

	server := grpc.NewServer()

	hd.RegisterLoginServiceServer(server, &s)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
