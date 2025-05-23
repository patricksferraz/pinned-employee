package repo

import (
	"context"

	"github.com/patricksferraz/pinned-employee/domain/entity"
)

type RepoInterface interface {
	CreateEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error
	SearchEmployees(ctx context.Context, searchEmployees *entity.SearchEmployees) ([]*entity.Employee, *string, error)

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
