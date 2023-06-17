package user

import (
	"context"
	"time"

	"api.turistikrota.com/auth/src/adapters/mongo/user/entity"
	"api.turistikrota.com/auth/src/domain/user"
	"github.com/mixarchitecture/i18np"
	sharedMongo "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) Create(ctx context.Context, email string, password []byte, token string) (*user.User, *i18np.Error) {
	user := r.userFactory.NewUser(email, password, token)
	exist, error := r.checkExist(ctx, user.Email)
	if error != nil {
		return nil, error
	}
	if exist {
		return nil, r.userFactory.Errors.AlreadyExists(user.Email)
	}
	u := &entity.MongoUser{}
	res, err := r.collection.InsertOne(ctx, u.FromUser(user))
	if err != nil {
		return nil, r.userFactory.Errors.Failed("create")
	}
	user.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return user, nil
}

func (r *repo) Update(ctx context.Context, user *user.User) (*user.User, *i18np.Error) {
	u := &entity.MongoUser{}
	exist, error := r.checkExist(ctx, user.Email)
	if error != nil {
		return nil, error
	}
	if !exist {
		return nil, r.userFactory.Errors.NotFound(user.Email)
	}
	res, err := r.collection.UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": u.FromUser(user)})
	if err != nil {
		return nil, r.userFactory.Errors.Failed("update")
	}
	if res.MatchedCount == 0 {
		return nil, r.userFactory.Errors.NotFound(user.Email)
	}
	return u.ToUser(), nil
}

func (r *repo) UpdateByUUID(ctx context.Context, user *user.User) (*user.User, *i18np.Error) {
	u := &entity.MongoUser{}
	id, err := sharedMongo.TransformId(user.UUID)
	if err != nil {
		return nil, r.userFactory.Errors.NotFound(user.UUID)
	}
	update := bson.M{"$set": bson.M{
		"email":              user.Email,
		"phone":              user.Phone,
		"roles":              user.Roles,
		"two_factor_enabled": user.TwoFactorEnabled,
		"is_active":          user.IsActive,
	}}
	res, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, r.userFactory.Errors.Failed("update")
	}
	if res.MatchedCount == 0 {
		return nil, r.userFactory.Errors.NotFound(user.UUID)
	}
	return u.ToUser(), nil
}

func (r *repo) GetByEmail(ctx context.Context, email string) (*user.User, *i18np.Error) {
	u := &entity.MongoUser{}
	res := r.collection.FindOne(ctx, bson.M{"email": email})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, r.userFactory.Errors.NotFound(email)
		}
		return nil, r.userFactory.Errors.Failed("get")
	}
	err := res.Decode(u)
	if err != nil {
		return nil, r.userFactory.Errors.Failed("get")
	}
	return u.ToUser(), nil
}

func (r *repo) CheckEmail(ctx context.Context, email string) (bool, *i18np.Error) {
	return r.checkExist(ctx, email)
}

func (r *repo) GetByUUID(ctx context.Context, uuid string) (*user.User, *i18np.Error) {
	u := &entity.MongoUser{}
	id, error := sharedMongo.TransformId(uuid)
	if error != nil {
		return nil, r.userFactory.Errors.NotFound(uuid)
	}
	res := r.collection.FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, r.userFactory.Errors.NotFound(uuid)
		}
		return nil, r.userFactory.Errors.Failed("get")
	}
	err := res.Decode(u)
	if err != nil {
		return nil, r.userFactory.Errors.Failed("get")
	}
	return u.ToUser(), nil
}

func (r *repo) GetByToken(ctx context.Context, token string) (*user.User, *i18np.Error) {
	u := &entity.MongoUser{}
	res := r.collection.FindOne(ctx, bson.M{"token": token})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, r.userFactory.Errors.NotFound(token)
		}
		return nil, r.userFactory.Errors.Failed("get")
	}
	err := res.Decode(u)
	if err != nil {
		return nil, r.userFactory.Errors.Failed("get")
	}
	return u.ToUser(), nil
}

func (r *repo) Verify(ctx context.Context, token string) *i18np.Error {
	filter := bson.M{
		"token": token,
	}
	update := bson.M{
		"$set": bson.M{
			"is_verified": true,
			"token":       "",
			"updated_at":  time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) SetToken(ctx context.Context, mail string, token string) *i18np.Error {
	filter := bson.M{
		"email": mail,
	}
	update := bson.M{
		"$set": bson.M{
			"is_verified": false,
			"token":       token,
			"updated_at":  time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) checkExist(ctx context.Context, email string) (bool, *i18np.Error) {
	res := r.collection.FindOne(ctx, bson.M{"email": email})
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, r.userFactory.Errors.Failed(err.Error())
	}
	return true, nil
}
