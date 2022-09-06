package datastore

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect(uri string, database string, configList ...mgm.Config) error {
	config := mgm.Config{
		CtxTimeout: 10 * time.Second,
	}

	if len(configList) > 0 {
		config = configList[0]
	}

	opts := options.Client().ApplyURI(uri)
	return mgm.SetDefaultConfig(&config, database, opts)
}
