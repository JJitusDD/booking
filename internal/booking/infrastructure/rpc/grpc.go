//go:generate protoc -I. -I../../domain --go-grpc_out=Minternal/book/domain/book.proto=.:. --go_out=Minternal/book/domain/book.proto=.:. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative book_rpc.proto

package booking_rpc

import (
	"context"

	booking "microservice-template-ddd/internal/booking/application"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"microservice-template-ddd/internal/user/infrastructure/rpc"
	"microservice-template-ddd/pkg/rpc"
)

func Use(_ context.Context, rpcClient *grpc.ClientConn) (BookRPCClient, error) {
	// Register clients
	client := NewBookRPCClient(rpcClient)

	return client, nil
}

type BookingServer struct {
	log *zap.Logger

	UnimplementedBookRPCServer

	// Application
	service *booking.Service

	// ServiceClients
	UserService    user_rpc.UserRPCClient
}

func New(runRPCServer *rpc.RPCServer, log *zap.Logger, bookService *booking.Service) (*BookingServer, error) {
	server := &BookingServer{
		log: log,

		service: bookService,
	}

	// Register services
	RegisterBookRPCServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}
