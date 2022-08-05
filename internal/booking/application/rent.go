/*
Book Service. Application layer
*/
package booking

import (
	"context"

	"microservice-template-ddd/internal/booking/domain"
	"microservice-template-ddd/internal/booking/infrastructure/store"
	"microservice-template-ddd/internal/user/infrastructure/rpc"
)

type Service struct {
	Store *store.BookStore

	// ServiceClients
	UserService    user_rpc.UserRPCClient
}

func New(store *store.BookStore, userService user_rpc.UserRPCClient) (*Service, error) {
	return &Service{
		Store: store,

		UserService:    userService,
	}, nil
}

// Get - get book from store
func (s *Service) Get(ctx context.Context, bookId string) (*domain.Book, error) {
	// Get book from store
	book, err := s.Store.Store.Get(ctx, bookId)
	if err != nil {
		// For example create book
		_, _ = s.Store.Store.Add(ctx, &domain.Book{
			Title:  "Hello World",
			Author: "God",
			IsRent: false,
		})

		return nil, err
	}

	return book, nil
}

func (s *Service) Rent(ctx context.Context, bookId string) (*domain.Book, error) {
	// Get user
	_, err := s.UserService.Get(ctx, &user_rpc.GetRequest{Id: bookId})
	if err != nil {
		return nil, err
	}

	// Get book from store
	book, err := s.Get(ctx, bookId)
	if err != nil {
		return nil, err
	}

	// Change state in DB
	book.IsRent = !book.IsRent
	book, err = s.Store.Store.Update(ctx, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}
