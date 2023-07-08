package mongo

import (
	mongo_account "github.com/turistikrota/service.auth/src/adapters/mongo/account"
	mongo_owner "github.com/turistikrota/service.auth/src/adapters/mongo/owner"
	mongo_user "github.com/turistikrota/service.auth/src/adapters/mongo/user"
	"github.com/turistikrota/service.auth/src/domain/account"
	"github.com/turistikrota/service.auth/src/domain/owner"
	"github.com/turistikrota/service.auth/src/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewUser(userFactory user.Factory, collection *mongo.Collection) user.Repository
	NewOwner(factory owner.Factory, collection *mongo.Collection) owner.Repository
	NewAccount(factory account.Factory, collection *mongo.Collection) account.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewUser(userFactory user.Factory, collection *mongo.Collection) user.Repository {
	return mongo_user.New(userFactory, collection)
}

func (m *mongodb) NewOwner(factory owner.Factory, collection *mongo.Collection) owner.Repository {
	return mongo_owner.New(factory, collection)
}

func (m *mongodb) NewAccount(factory account.Factory, collection *mongo.Collection) account.Repository {
	return mongo_account.New(factory, collection)
}
