package repo

import (
	"context"

	"github.com/c-4u/pinned-employee/domain/entity"
)

type RepoInterface interface {
	CreateEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
