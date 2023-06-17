package adapters

import (
	"api.turistikrota.com/auth/src/adapters/memory"
	"api.turistikrota.com/auth/src/adapters/mongo"
	"api.turistikrota.com/auth/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
