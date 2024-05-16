package intelisys

import (
	"context"
	"project-v/internal/model/agent"
)

//go:generate mockgen -source=repository.go -destination ./repository_mock.go -package=workflow
type IRepository interface {
	GetReservations(
		ctx context.Context, pnr string,
	) (*agent.Reservation, error)
}
