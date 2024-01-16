package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FilterEntity struct {
	Query string `query:"q,omitempty" validate:"omitempty,max=500"`
}

func (r *repo) filterToBson(filter FilterEntity) bson.M {
	list := make([]bson.M, 0)
	list = r.filterByQuery(list, filter)
	listLen := len(list)
	if listLen == 0 {
		return bson.M{}
	}
	if listLen == 1 {
		return list[0]
	}
	return bson.M{
		"$and": list,
	}
}

func (r *repo) filterByQuery(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Query != "" {
		list = append(list, bson.M{
			"$or": []bson.M{
				{
					fields.Email: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					fields.Phone: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					fields.UUID: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
			},
		})
	}
	return list
}

func (r *repo) sort(opts *options.FindOptions, filter FilterEntity) *options.FindOptions {
	opts.SetSort(bson.M{
		fields.CreatedAt: -1,
	})
	return opts
}
