package owner

import (
	"context"

	"api.turistikrota.com/auth/src/adapters/mongo/owner/entity"
	"api.turistikrota.com/auth/src/domain/owner"
	"api.turistikrota.com/shared/jwt"
	"github.com/mixarchitecture/i18np"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Create(ctx context.Context, owner *owner.Entity) *i18np.Error {
	n := &entity.MongoOwner{}
	_, err := r.collection.InsertOne(ctx, n.FromOwner(owner))
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	return nil
}

func (r *repo) AddUser(ctx context.Context, nickName string, user *owner.User) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
		"$or": []bson.M{
			{entity.UserField(entity.UserFields.Name): bson.M{"$ne": user.Name}},
			{entity.UserField(entity.UserFields.Code): bson.M{"$ne": user.Code}},
		},
	}
	setter := bson.M{
		"$addToSet": bson.M{
			entity.Fields.Users: bson.M{
				entity.UserFields.UUID:   user.UUID,
				entity.UserFields.Name:   user.Name,
				entity.UserFields.Code:   user.Code,
				entity.UserFields.Roles:  user.Roles,
				entity.UserFields.JoinAt: user.JoinAt,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) RemoveUser(ctx context.Context, nickName string, user owner.UserDetail) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
		entity.UserField(entity.UserFields.Code): user.Code,
	}
	setter := bson.M{
		"$pull": bson.M{
			entity.Fields.Users: bson.M{
				entity.UserFields.Name: user.Name,
				entity.UserFields.Code: user.Code,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) RemoveUserPermission(ctx context.Context, nickName string, user owner.UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
		entity.UserField(entity.UserFields.Code): user.Code,
	}
	setter := bson.M{
		"$pull": bson.M{
			entity.UserArrayFieldInArray(entity.UserFields.Roles): permission,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) AddUserPermission(ctx context.Context, nickName string, user owner.UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
		entity.UserField(entity.UserFields.Code): user.Code,
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

func (r *repo) GetAllAsClaim(ctx context.Context, userUUID string) ([]jwt.UserClaimOwner, *i18np.Error) {
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
	owners := make([]jwt.UserClaimOwner, 0)
	for cursor.Next(ctx) {
		var owner entity.MongoOwner
		if err := cursor.Decode(&owner); err != nil {
			return nil, r.factory.Errors.Failed("get")
		}
		owners = append(owners, owner.ToClaim(userUUID))
	}
	return owners, nil
}
