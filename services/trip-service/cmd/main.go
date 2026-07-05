package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {

	ctx := context.Background()                  // 创建上下文
	inmemRepo := repository.NewInmemRepository() // 初始化内存数据库
	svc := service.NewService(inmemRepo)         // 初始化service，注入内存数据库依赖

	fare := &domain.RideFareModel{ // 创建 行程费模型
		UserID: "42",
	}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	// 保持程序运行（暂时）
	for {
		time.Sleep(time.Second)
	}
}
