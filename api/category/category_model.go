package category

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Title string			 `json:"title,omitempty" validate:"required"`
	Color string 			 `json:"color,omitempty" validate:"required"`
}