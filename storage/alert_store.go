package storage

import "go.mongodb.org/mongo-driver/mongo"

type AlertStorageImpl struct{ collection *mongo.Collection }

func NewAlertStorage(collection *mongo.Collection) *AlertStorageImpl {
	return &AlertStorageImpl{
		collection: collection,
	}
}
