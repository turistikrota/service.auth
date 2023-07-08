package adapters

import (
	"github.com/turistikrota/service.auth/src/adapters/memory"
	"github.com/turistikrota/service.auth/src/adapters/mongo"
	"github.com/turistikrota/service.auth/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
