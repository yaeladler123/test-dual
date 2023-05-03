package handler

import (
	"context"
	"github.com/calculi-corp/api/example/go/potato"
	"github.com/go-playground/assert/v2"
	"testing"
)

// Unit tests are implemented on the Handler and not the gRPC Server in order to avoid pulling Certificates
// and other plumbing code.

func TestGetPotatoes(t *testing.T) {
	ctx := context.Background()
	tested := NewPotatoHandler()

	req := &potato.GetPotatoesRequest{}

	resp, err := tested.GetPotatoes(ctx, req)

	if err != nil {
		t.Error("Error should be nil")
	}
	assert.Equal(t, len(resp.Potatoes), 1)
	assert.Equal(t, resp.Potatoes[0].GetId(), "1")
}

func TestGetPotato(t *testing.T) {
	ctx := context.Background()
	tested := NewPotatoHandler()

	req := &potato.GetPotatoRequest{
		PotatoId: "42",
	}

	resp, err := tested.GetPotato(ctx, req)

	if err != nil {
		t.Error("Error should be nil")
	}
	assert.Equal(t, resp.GetPotato().GetId(), req.PotatoId)
}
