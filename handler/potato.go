package handler

import (
	"context"
	"github.com/calculi-corp/api/example/go/potato"
	handler "github.com/calculi-corp/grpc-handler"
	hostflags "github.com/calculi-corp/grpc-hostflags"
	"github.com/calculi-corp/log"
	"google.golang.org/grpc"
)

const (
	handlerName = "PotatoHandler"
	handlerDesc = "Potato Example Handler"
)

var (
	dependencies = []string{
		hostflags.DbService,
	}
)

type PotatoHandler struct {
	potato.UnimplementedPotatoServiceServer
	metrics *handler.Map
}

func NewPotatoHandler() *PotatoHandler {
	log.Debug("Initializing PotatoHandler...")
	return &PotatoHandler{
		metrics: handler.NewMap(handlerName),
	}
}

func (ph *PotatoHandler) RegisterPotatoServer(s *grpc.Server) {
	log.Debug("Registering PotatoHandler...")
	potato.RegisterPotatoServiceServer(s, ph)
}

func (ph *PotatoHandler) GetPotato(ctx context.Context, req *potato.GetPotatoRequest) (*potato.GetPotatoResponse, error) {
	log.Debug("GetPotato called")
	return &potato.GetPotatoResponse{
		Potato: &potato.Potato{
			Id:   req.PotatoId,
			Name: "Sweet Potato",
			Size: 7,
		},
	}, nil
}

func (ph *PotatoHandler) GetPotatoes(ctx context.Context, req *potato.GetPotatoesRequest) (*potato.GetPotatoesResponse, error) {
	log.Debug("GetPotatoes called")
	return &potato.GetPotatoesResponse{
		Potatoes: []*potato.Potato{
			&potato.Potato{
				Id:   "1",
				Name: "Sweet Potato",
				Size: 7,
			},
		},
	}, nil
}
