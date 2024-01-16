package account

import (
	"context"

	"github.com/cilloparch/cillop/i18np"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserUnique struct {
	UserUUID string
	Name     string
}

type Repo interface {
	Create(ctx context.Context, account *Entity) *i18np.Error
	Update(ctx context.Context, u UserUnique, account *Entity) *i18np.Error
	Disable(ctx context.Context, u UserUnique) *i18np.Error
	Enable(ctx context.Context, u UserUnique) *i18np.Error
	Restore(ctx context.Context, u UserUnique) *i18np.Error
	Delete(ctx context.Context, u UserUnique) *i18np.Error
	ListAsClaims(ctx context.Context, userUUID string) ([]jwt.UserClaimAccount, *i18np.Error)
}

type repo struct {
	factory    Factory
	collection *mongo.Collection
	helper     mongo2.Helper[*Entity, *Entity]
}

func NewRepo(collection *mongo.Collection, factory Factory) Repo {
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[*Entity, *Entity](collection, createEntity),
	}
}

func createEntity() **Entity {
	return new(*Entity)
}

func (r *repo) Create(ctx context.Context, e *Entity) *i18np.Error {
	_, err := r.collection.InsertOne(ctx, e)
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	return nil
}

func (r *repo) Update(ctx context.Context, u UserUnique, account *Entity) *i18np.Error {
	filter := bson.M{
		fields.UserUUID: u.UserUUID,
		fields.UserName: u.Name,
	}
	update := bson.M{
		"$set": bson.M{
			fields.UserName:  account.UserName,
			fields.BirthDate: account.BirthDate,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Disable(ctx context.Context, u UserUnique) *i18np.Error {
	filter := bson.M{
		fields.UserUUID: u.UserUUID,
		fields.UserName: u.Name,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsActive: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Enable(ctx context.Context, u UserUnique) *i18np.Error {
	filter := bson.M{
		fields.UserUUID: u.UserUUID,
		fields.UserName: u.Name,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsActive: true,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Restore(ctx context.Context, u UserUnique) *i18np.Error {
	filter := bson.M{
		fields.UserUUID: u.UserUUID,
		fields.UserName: u.Name,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsDeleted: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Delete(ctx context.Context, u UserUnique) *i18np.Error {
	filter := bson.M{
		fields.UserUUID: u.UserUUID,
		fields.UserName: u.Name,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsDeleted: true,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) ListAsClaims(ctx context.Context, userUUID string) ([]jwt.UserClaimAccount, *i18np.Error) {
	filter := bson.M{
		fields.UserUUID: userUUID,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
		fields.IsActive: true,
	}
	res, err := r.helper.GetListFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("list")
	}
	claims := make([]jwt.UserClaimAccount, 0)
	for _, v := range res {
		claims = append(claims, jwt.UserClaimAccount{
			Name: v.UserName,
			ID:   v.UserUUID,
		})
	}
	return claims, nil
}
