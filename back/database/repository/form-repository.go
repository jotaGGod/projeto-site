package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"site-backend/database"
	"site-backend/database/entities"
)

func GetFormRepository() FormRepository {
	return FormRepository{collection: "UserForm"}
}

type FormRepository struct {
	collection string
}

func (repo *FormRepository) CreateForm(entity entities.FormEntity) (*entities.FormEntity, error) {
	collection := database.GetCollection(repo.collection)

	formId, err := collection.InsertOne(context.TODO(), entity)
	if err != nil {
		return nil, err
	}

	collection.FindOne(context.TODO(), bson.M{"Id": formId.InsertedID}).Decode(&entity)

	return &entity, err
}
