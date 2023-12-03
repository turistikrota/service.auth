package business

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.auth/src/adapters/mongo/business/entity"
	"github.com/turistikrota/service.auth/src/domain/business"
	"github.com/turistikrota/service.shared/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Create(ctx context.Context, business *business.Entity) *i18np.Error {
	n := &entity.MongoBusiness{}
	_, err := r.collection.InsertOne(ctx, n.FromBusiness(business))
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	return nil
}

func (r *repo) AddUser(ctx context.Context, nickName string, user *business.User) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
		"$or": []bson.M{
			{entity.UserField(entity.UserFields.Name): bson.M{"$ne": user.Name}},
		},
	}
	setter := bson.M{
		"$addToSet": bson.M{
			entity.Fields.Users: bson.M{
				entity.UserFields.UUID:   user.UUID,
				entity.UserFields.Name:   user.Name,
				entity.UserFields.Roles:  user.Roles,
				entity.UserFields.JoinAt: user.JoinAt,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) RemoveUser(ctx context.Context, nickName string, user business.UserDetail) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
	}
	setter := bson.M{
		"$pull": bson.M{
			entity.Fields.Users: bson.M{
				entity.UserFields.Name: user.Name,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) RemoveUserPermission(ctx context.Context, nickName string, user business.UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
	}
	setter := bson.M{
		"$pull": bson.M{
			entity.UserArrayFieldInArray(entity.UserFields.Roles): permission,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) AddUserPermission(ctx context.Context, nickName string, user business.UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
	}
	setter := bson.M{
		"$push": bson.M{
			entity.UserArrayFieldInArray(entity.UserFields.Roles): permission,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsEnabled:  true,
			entity.Fields.IsVerified: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsEnabled:  false,
			entity.Fields.IsVerified: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted:  true,
			entity.Fields.IsVerified: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Recover(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted:  false,
			entity.Fields.IsVerified: false,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Verify(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsVerified: true,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) GetAllAsClaim(ctx context.Context, userUUID string) ([]jwt.UserClaimBusiness, *i18np.Error) {
	filter := bson.M{
		entity.UserField(entity.UserFields.UUID): userUUID,
	}
	projection := bson.M{
		entity.Fields.NickName: 1,
		entity.Fields.Users:    1,
	}
	cursor, err := r.collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	defer cursor.Close(ctx)
	businesses := make([]jwt.UserClaimBusiness, 0)
	for cursor.Next(ctx) {
		var business entity.MongoBusiness
		if err := cursor.Decode(&business); err != nil {
			return nil, r.factory.Errors.Failed("get")
		}
		businesses = append(businesses, business.ToClaim(userUUID))
	}
	return businesses, nil
}
