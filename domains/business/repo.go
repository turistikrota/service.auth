package business

import (
	"context"

	"github.com/cilloparch/cillop/i18np"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDetail struct {
	Name string
	UUID string
}

type Repo interface {
	Create(ctx context.Context, entity *Entity) *i18np.Error
	AddUser(ctx context.Context, nickName string, user *User) *i18np.Error
	RemoveUser(ctx context.Context, nickName string, user UserDetail) *i18np.Error
	RemoveUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error
	AddUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error
	Enable(ctx context.Context, nickName string) *i18np.Error
	Disable(ctx context.Context, nickName string) *i18np.Error
	Delete(ctx context.Context, nickName string) *i18np.Error
	Recover(ctx context.Context, nickName string) *i18np.Error
	Verify(ctx context.Context, nickName string) *i18np.Error
	GetAllAsClaim(ctx context.Context, userUUID string) ([]jwt.UserClaimBusiness, *i18np.Error)
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

func (r *repo) AddUser(ctx context.Context, nickName string, user *User) *i18np.Error {
	filter := bson.M{
		fields.NickName:            nickName,
		userField(userFields.Name): bson.M{"$ne": user.Name},
	}
	update := bson.M{
		"$addToSet": bson.M{
			fields.Users: bson.M{
				userFields.UUID:   user.UUID,
				userFields.Name:   user.Name,
				userFields.Roles:  user.Roles,
				userFields.JoinAt: user.JoinAt,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) RemoveUser(ctx context.Context, nickName string, user UserDetail) *i18np.Error {
	filter := bson.M{
		fields.NickName:            nickName,
		userField(userFields.Name): user.Name,
	}
	update := bson.M{
		"$pull": bson.M{
			fields.Users: bson.M{
				userFields.Name: user.Name,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) RemoveUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		fields.NickName:            nickName,
		userField(userFields.Name): user.Name,
	}
	setter := bson.M{
		"$pull": bson.M{
			userArrayFieldInArray(userFields.Roles): permission,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) AddUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		fields.NickName:            nickName,
		userField(userFields.Name): user.Name,
	}
	setter := bson.M{
		"$push": bson.M{
			userArrayFieldInArray(userFields.Roles): permission,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsEnabled:  true,
			fields.IsVerified: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsEnabled:  false,
			fields.IsVerified: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsDeleted:  true,
			fields.IsVerified: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Recover(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsDeleted:  false,
			fields.IsVerified: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Verify(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsVerified: true,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) GetAllAsClaim(ctx context.Context, userUUID string) ([]jwt.UserClaimBusiness, *i18np.Error) {
	filter := bson.M{
		userField(userFields.UUID): userUUID,
	}
	projection := bson.M{
		fields.NickName: 1,
		fields.Users:    1,
	}
	res, err := r.helper.GetListFilter(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	businesses := make([]jwt.UserClaimBusiness, 0)
	for _, item := range res {
		businesses = append(businesses, item.ToClaim(userUUID))
	}
	return businesses, nil
}
