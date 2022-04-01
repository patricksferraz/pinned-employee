package service

import (
	"context"

	"github.com/c-4u/pinned-employee/domain/entity"
	"github.com/c-4u/pinned-employee/domain/repo"
	"github.com/c-4u/pinned-employee/infra/client/kafka/topic"
	"github.com/c-4u/pinned-employee/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateEmployee(ctx context.Context, name *string) (*string, error) {
	employee, err := entity.NewEmployee(name)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateEmployee(ctx, employee); err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(employee)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_EMPLOYEE), utils.PString(string(eMsg)), employee.ID)
	if err != nil {
		return nil, err
	}

	return employee.ID, nil
}

func (s *Service) FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error) {
	employee, err := s.Repo.FindEmployee(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *Service) SearchEmployees(ctx context.Context, pageToken *string, pageSize *int) ([]*entity.Employee, *string, error) {
	pagination, err := entity.NewPagination(pageToken, pageSize)
	if err != nil {
		return nil, nil, err
	}

	searchEmployees, err := entity.NewSearchEmployees(pagination)
	if err != nil {
		return nil, nil, err
	}

	employees, nextPageToken, err := s.Repo.SearchEmployees(ctx, searchEmployees)
	if err != nil {
		return nil, nil, err
	}

	return employees, nextPageToken, nil
}
