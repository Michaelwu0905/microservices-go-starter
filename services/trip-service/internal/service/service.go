package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 行程的服务层，包含其存储层实现
type service struct {
	repo domain.TripRepository
}

// NewService 新建/实例化一个service层
func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

// CreateTrip 创建一个行程，存入数据库
func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending", // 状态：排队中
		RideFare: fare,
	}
	return s.repo.CreateTrip(ctx, t)
}
