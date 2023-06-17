package account

import (
	"context"

	"api.turistikrota.com/auth/src/adapters/mongo/account/entity"
	"api.turistikrota.com/auth/src/domain/account"
	"api.turistikrota.com/shared/jwt"
	"github.com/mixarchitecture/i18np"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) updateOne(ctx context.Context, filter bson.M, setter bson.M, opts ...*options.UpdateOptions) *i18np.Error {
	res, err := r.collection.UpdateOne(ctx, filter, setter, opts...)
	if err != nil {
		return r.factory.Errors.Failed("update")
	}
	if res.MatchedCount == 0 {
		return r.factory.Errors.NotFound()
	}
	return nil
}

func (r *repo) Create(ctx context.Context, account *account.Entity) *i18np.Error {
	n := &entity.MongoAccount{}
	_, err := r.collection.InsertOne(ctx, n.FromEntity(account))
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	return nil
}

func (r *repo) Update(ctx context.Context, u account.UserUnique, account *account.Entity) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UserUUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.UserName:  account.UserName,
			entity.Fields.UserCode:  account.UserCode,
			entity.Fields.BirthDate: account.BirthDate,
		},
	}
	return r.updateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, u account.UserUnique) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UserUUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsActive: false,
		},
	}
	return r.updateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, u account.UserUnique) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UserUUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsActive: true,
		},
	}
	return r.updateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, u account.UserUnique) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UserUUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted: true,
		},
	}
	return r.updateOne(ctx, filter, setter)
}

func (r *repo) ListAsClaims(ctx context.Context, userUUID string) ([]jwt.UserClaimAccount, *i18np.Error) {
	filter := bson.M{
		entity.Fields.UserUUID: userUUID,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
		entity.Fields.IsActive: true,
	}
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("list")
	}
	defer cur.Close(ctx)
	claims := make([]jwt.UserClaimAccount, 0)
	for cur.Next(ctx) {
		n := &entity.MongoAccount{}
		err := cur.Decode(n)
		if err != nil {
			return nil, r.factory.Errors.Failed("list")
		}
		claims = append(claims, n.ToClaim())
	}
	if err := cur.Err(); err != nil {
		return nil, r.factory.Errors.Failed("list")
	}
	return claims, nil
}
