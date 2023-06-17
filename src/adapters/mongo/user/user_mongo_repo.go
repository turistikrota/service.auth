package user

import (
	"api.turistikrota.com/auth/src/adapters/mongo/user/entity"
	"api.turistikrota.com/auth/src/domain/user"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	userFactory user.Factory
	collection  *mongo.Collection
	helper      mongo2.Helper[entity.MongoUser, *user.User]
}

func New(userFactory user.Factory, collection *mongo.Collection) user.Repository {
	validate(userFactory, collection)
	return &repo{
		userFactory: userFactory,
		collection:  collection,
		helper:      mongo2.NewHelper[entity.MongoUser, *user.User](collection, createEntity),
	}
}

func validate(userFactory user.Factory, collection *mongo.Collection) {
	if userFactory.IsZero() {
		panic("exampleFactory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}

func createEntity() *entity.MongoUser {
	return &entity.MongoUser{}
}
