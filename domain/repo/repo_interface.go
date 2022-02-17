package repo

import (
	"context"

	"github.com/c-4u/attendant/domain/entity"
)

type RepoInterface interface {
	CreateAttendant(ctx context.Context, attendant *entity.Attendant) error
	FindAttendant(ctx context.Context, attendantID *string) (*entity.Attendant, error)
	SaveAttendant(ctx context.Context, attendant *entity.Attendant) error

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
