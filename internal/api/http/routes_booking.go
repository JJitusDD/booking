package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"google.golang.org/protobuf/encoding/protojson"

	"microservice-template-ddd/internal/booking/infrastructure/rpc"
)

// Routes creates a REST router
func (api *API) BookingRoutes() chi.Router {
	r := chi.NewRouter()

	// CRUD
	r.Post("/", api.Booking)
	r.Get("/", api.ListAvailableBook)

	return r
}

// CRUD ================================================================================================================
func (api *API) Booking(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (api *API) ListAvailableBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	resp, err := api.BookService.Get(r.Context(), &book_rpc.GetRequest{Id: "Hello World"})
	if err != nil {
		api.Log.Error(err.Error())
		_, _ = w.Write([]byte(`{"error": "error 0_o"}`))
		return
	}

	m := protojson.MarshalOptions{}
	payload, err := m.Marshal(resp)
	if err != nil {
		api.Log.Error(err.Error())
		_, _ = w.Write([]byte(`{"error": "error 0_o"}`))
	}

	_, _ = w.Write(payload)
}
