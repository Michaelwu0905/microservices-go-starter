package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct { // 行程模型
	ID       primitive.ObjectID // ID
	UserID   string             // 用户ID
	Status   string             // 行程状态
	RideFare *RideFareModel     // 对应的车费模型
}

type TripRepository interface { // 行程的存储层接口
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface { // 行程的服务层接口
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}
