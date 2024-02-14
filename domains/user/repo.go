package user

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo interface {
	GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	GetByEmail(ctx context.Context, email string) (*Entity, *i18np.Error)
	GetByToken(ctx context.Context, token string) (*Entity, *i18np.Error)
	CheckEmail(ctx context.Context, email string) (bool, *i18np.Error)
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	SetToken(ctx context.Context, email string, token string) *i18np.Error
	Verify(ctx context.Context, token string) *i18np.Error
	AddRoles(ctx context.Context, uuid string, roles []string) *i18np.Error
	RemoveRoles(ctx context.Context, uuid string, roles []string) *i18np.Error
	Delete(ctx context.Context, uuid string) *i18np.Error
	Recover(ctx context.Context, uuid string) *i18np.Error
	SetPassword(ctx context.Context, uuid string, password []byte) *i18np.Error
	List(ctx context.Context, filter FilterEntity, listConf list.Config) (*list.Result[*Entity], *i18np.Error)
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

func (r *repo) GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error) {
	id, _err := mongo2.TransformId(uuid)
	if _err != nil {
		return nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	res, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound(uuid)
	}
	return *res, nil
}

func (r *repo) GetByEmail(ctx context.Context, email string) (*Entity, *i18np.Error) {
	filter := bson.M{
		fields.Email: email,
	}
	res, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound(email)
	}
	return *res, nil
}

func (r *repo) GetByToken(ctx context.Context, token string) (*Entity, *i18np.Error) {
	filter := bson.M{
		fields.VerifyToken: token,
	}
	res, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound(token)
	}
	return *res, nil
}

func (r *repo) CheckEmail(ctx context.Context, email string) (bool, *i18np.Error) {
	filter := bson.M{
		fields.Email: email,
	}
	_, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return false, r.factory.Errors.Failed("get")
	}
	return !exist, nil
}

func (r *repo) Create(ctx context.Context, e *Entity) (*Entity, *i18np.Error) {
	res, err := r.collection.InsertOne(ctx, e)
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	e.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return e, nil
}

func (r *repo) SetToken(ctx context.Context, email string, token string) *i18np.Error {
	filter := bson.M{
		fields.Email: email,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsVerified:  false,
			fields.VerifyToken: token,
			fields.UpdatedAt:   time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Verify(ctx context.Context, token string) *i18np.Error {
	filter := bson.M{
		fields.VerifyToken: token,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsVerified:  true,
			fields.VerifyToken: "",
			fields.UpdatedAt:   time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) AddRoles(ctx context.Context, uuid string, roles []string) *i18np.Error {
	id, _err := mongo2.TransformId(uuid)
	if _err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$addToSet": bson.M{
			fields.Roles: bson.M{
				"$each": roles,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) RemoveRoles(ctx context.Context, uuid string, roles []string) *i18np.Error {
	id, _err := mongo2.TransformId(uuid)
	if _err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$pull": bson.M{
			fields.Roles: bson.M{
				"$in": roles,
			},
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Delete(ctx context.Context, uuid string) *i18np.Error {
	id, _err := mongo2.TransformId(uuid)
	if _err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsDeleted: true,
			fields.DeletedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Recover(ctx context.Context, uuid string) *i18np.Error {
	id, _err := mongo2.TransformId(uuid)
	if _err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsDeleted: false,
			fields.DeletedAt: nil,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) SetPassword(ctx context.Context, uuid string, password []byte) *i18np.Error {
	id, _err := mongo2.TransformId(uuid)
	if _err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.Password: password,
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) List(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	filters := r.filterToBson(filter)
	l, err := r.helper.GetListFilter(ctx, filters, r.sort(r.filterOptions(listConfig), filter))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, bson.M{})
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) filterOptions(listConfig list.Config) *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetProjection(bson.M{
		fields.Password: 0,
	}).SetSkip(listConfig.Offset).SetLimit(listConfig.Limit)
	return opts
}
