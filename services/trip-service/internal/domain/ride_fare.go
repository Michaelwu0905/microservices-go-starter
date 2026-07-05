package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct { // 价格方案模型
	ID                primitive.ObjectID // ID
	UserID            string             // 用户ID
	PackageSlug       string             // 套餐方案
	TotalPriceInCents float64            // 总价
}
