package dtos

import "go.mongodb.org/mongo-driver/bson"

type ListOptions struct {
	Limit  int
	Offset int
	Filter bson.M
}
