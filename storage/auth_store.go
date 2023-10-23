package storage

import (
	"context"

	"github.com/luqus/s/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthStorageImpl struct{ collection *mongo.Collection }

func NewAuthStorage(collection *mongo.Collection) *AuthStorageImpl {
	return &AuthStorageImpl{
		collection: collection,
	}
}

func (store *AuthStorageImpl) CreateUser(ctx context.Context, user *types.User) error {
	_, err := store.collection.InsertOne(ctx, user)

	return err
}

func (store *AuthStorageImpl) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	filter := primitive.D{primitive.E{Key: "email", Value: email}}

	user := new(types.User)

	err := store.collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (store *AuthStorageImpl) GetUserByUID(ctx context.Context, uid string) (*types.User, error) {
	filter := primitive.D{primitive.E{Key: "uid", Value: uid}}
	user := new(types.User)

	err := store.collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (store *AuthStorageImpl) CheckIfEmailExists(ctx context.Context, email string) (int64, error) {
	filter := primitive.D{primitive.E{Key: "email", Value: email}}

	result, err := store.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, nil
	}

	return result, nil
}

func (store *AuthStorageImpl) CheckIfPhoneNumberExists(ctx context.Context, phoneNumber string) (int64, error) {
	filter := primitive.D{primitive.E{Key: "phone_number", Value: phoneNumber}}

	result, err := store.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, nil
	}

	return result, nil
}
