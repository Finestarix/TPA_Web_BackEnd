package types

import "github.com/graphql-go/graphql"

var carType *graphql.Object

func GetCarType() *graphql.Object {
	if carType == nil {
		carType = graphql.NewObject(graphql.ObjectConfig{
			Name: "CarType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"location": &graphql.Field{
					Type: GetLocationType(),
				},
				"carModel": &graphql.Field{
					Type: GetCarModelType(),
				},
			},
			Description: "Type for Car Model",
		})
	}
	return carType
}
