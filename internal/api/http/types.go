package http

import (
	"context"

	"go.uber.org/zap"

	booking_rpc "microservice-template-ddd/internal/booking/infrastructure/rpc"
	"microservice-template-ddd/internal/user/infrastructure/rpc"
)

// API ...
type API struct { // nolint unused
	ctx context.Context
	Log *zap.Logger

	UserService    user_rpc.UserRPCClient
	BookingService    booking_rpc.BookingServer
}
