package main

import "ride-sharing/shared/types"

// 定义API gateway自己接收的请求结构

// 定义前端请求 trip preview 时传的数据
type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

// 也就是说前端可能传来的数据如下
// {
//   "userID": "42",
//   "pickup": {
//     "latitude": 31.23,
//     "longitude": 121.47
//   },
//   "destination": {
//     "latitude": 31.20,
//     "longitude": 121.50
//   }
// }	Coordinate 来自共享包 shared
