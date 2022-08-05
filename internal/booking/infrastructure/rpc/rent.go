package booking_rpc

import (
	"context"
)

func (m *BookingServer) Rent(ctx context.Context, in *BookingRequest) (*BookingResponse, error) {
	book, err := m.service.Rent(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &BookingResponse{
		Book: book,
	}, nil
}
