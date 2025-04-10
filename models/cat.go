package models

import "go.mongodb.org/mongo-driver/bson"

type OptionsCatFunc func(c *CatOptions) error

type CatOptions struct {
	Args []string
}

type CatResponse struct {
	Documents []bson.M
	Count     int64
}
