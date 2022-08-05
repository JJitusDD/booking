//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	book "microservice-template-ddd/internal/booking/application"
	book_rpc "microservice-template-ddd/internal/booking/infrastructure/rpc"
	"microservice-template-ddd/internal/booking/infrastructure/store"
	"microservice-template-ddd/internal/db"
	user_rpc "microservice-template-ddd/internal/user/infrastructure/rpc"
	"microservice-template-ddd/pkg/rpc"
)

type BookingService struct {
	Log *zap.Logger

	bookingRPCServer *book_rpc.BookingServer
}

// BookingService =========================================================================================================
var BookingSet = wire.NewSet(
	// log, tracer
	DefaultSet,

	// gRPC server
	runGRPCServer,
	NewBookingRPCServer,

	// gRPC client
	runGRPCClient,
	NewUserRPCClient,

	// store
	InitStore,
	InitBookingStore,

	// applications
	NewBookingApplication,

	// CMD
	NewBookingService,
)

// InitConstructor =====================================================================================================
func InitBookingStore(ctx context.Context, log *zap.Logger, conn *db.Store) (*store.BookingStore, error) {
	st := store.BookingStore{}
	bookStore, err := st.Use(ctx, log, conn)
	if err != nil {
		return nil, err
	}

	return bookStore, nil
}

func NewBookingApplication(store *store.BookingStore, userRPC user_rpc.UserRPCClient) (*book.Service, error) {
	bookService, err := book.New(store, userRPC)
	if err != nil {
		return nil, err
	}

	return bookService, nil
}

func NewBookingRPCClient(ctx context.Context, log *zap.Logger, rpcClient *grpc.ClientConn) (book_rpc.BookingRPCClient, error) {
	bookService, err := book_rpc.Use(ctx, rpcClient)
	if err != nil {
		return nil, err
	}

	return bookService, nil
}

func NewBookingRPCServer(bookService *book.Service, log *zap.Logger, serverRPC *rpc.RPCServer) (*book_rpc.BookingServer, error) {
	bookRPCServer, err := book_rpc.New(serverRPC, log, bookService)
	if err != nil {
		return nil, err
	}

	return bookRPCServer, nil
}

func NewBookingService(log *zap.Logger, bookingRPCServer *book_rpc.BookingServer) (*BookService, error) {
	return &BookService{
		Log: log,

		bookRPCServer: bookRPCServer,
	}, nil
}

func InitializeBookingService(ctx context.Context) (*BookService, func(), error) {
	panic(wire.Build(BookSet))
}
