package repo

import (
	"context"
	"fmt"

	"github.com/c-4u/pinned-employee/domain/entity"
	"github.com/c-4u/pinned-employee/infra/client/kafka"
	"github.com/c-4u/pinned-employee/infra/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
	Pg *db.PostgreSQL
	Kp *kafka.KafkaProducer
}

func NewRepository(pg *db.PostgreSQL, kp *kafka.KafkaProducer) *Repository {
	return &Repository{
		Pg: pg,
		Kp: kp,
	}
}

func (r *Repository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Pg.Db.Create(employee).Error
	return err
}

func (r *Repository) FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error) {
	var e entity.Employee

	r.Pg.Db.First(&e, "id = ?", *employeeID)

	if e.ID == nil {
		return nil, fmt.Errorf("no employee found")
	}

	return &e, nil
}

func (r *Repository) SaveEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Pg.Db.Save(employee).Error
	return err
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
