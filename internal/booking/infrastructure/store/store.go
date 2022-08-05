/*
Store Endpoint
*/
package store

import (
	"context"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"microservice-template-ddd/internal/booking/infrastructure/store/redis"
	"microservice-template-ddd/internal/db"
)

// Use return implementation of db
func (store *BookingStore) Use(ctx context.Context, log *zap.Logger, db *db.Store) (*BookingStore, error) { // nolint unused
	// Set configuration
	store.setConfig()

	switch store.typeStore {
	case "redis":
		store.Store = &redis.Store{}
	default:
		store.Store = &redis.Store{}
	}

	// Init store
	if err := store.Store.Init(ctx, db); err != nil {
		return nil, err
	}

	log.Info("init BookingStore", zap.String("db", store.typeStore))

	return store, nil
}

// setConfig - set configuration
func (s *BookingStore) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "redis") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram, scylla, cassandra
	s.typeStore = viper.GetString("STORE_TYPE")
}
