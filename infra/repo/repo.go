package repo

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/pinned-employee/domain/entity"
	"github.com/patricksferraz/pinned-employee/infra/client/kafka"
	"github.com/patricksferraz/pinned-employee/infra/db"
)

type Repository struct {
	Orm *db.DbOrm
	Kp  *kafka.KafkaProducer
}

func NewRepository(orm *db.DbOrm, kp *kafka.KafkaProducer) *Repository {
	return &Repository{
		Orm: orm,
		Kp:  kp,
	}
}

func (r *Repository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Orm.Db.Create(employee).Error
	return err
}

func (r *Repository) FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error) {
	var e entity.Employee

	r.Orm.Db.First(&e, "id = ?", *employeeID)

	if e.ID == nil {
		return nil, fmt.Errorf("no employee found")
	}

	return &e, nil
}

func (r *Repository) SaveEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Orm.Db.Save(employee).Error
	return err
}

func (r *Repository) SearchEmployees(ctx context.Context, searchEmployees *entity.SearchEmployees) ([]*entity.Employee, *string, error) {
	var e []*entity.Employee

	q := r.Orm.Db
	if *searchEmployees.PageToken != "" {
		q = q.Where("token < ?", *searchEmployees.PageToken)
	}
	err := q.Order("token DESC").
		Limit(*searchEmployees.PageSize).
		Find(&e).Error
	if err != nil {
		return nil, nil, err
	}

	if len(e) < *searchEmployees.PageSize {
		return e, nil, nil
	}

	return e, e[len(e)-1].Token, nil
}

func (r *Repository) PublishEvent(ctx context.Context, topic, msg, key *string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: topic, Partition: ckafka.PartitionAny},
		Value:          []byte(*msg),
		Key:            []byte(*key),
	}
	err := r.Kp.Producer.Produce(message, r.Kp.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
