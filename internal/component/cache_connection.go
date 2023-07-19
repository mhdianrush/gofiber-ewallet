package component

import (
	"context"
	"gofiber-ewallet/domain"
	"time"

	"github.com/allegro/bigcache/v3"
)

func GetCacheConnnection() domain.CacheRepository {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		logger.Printf("error connect cache %s", err.Error())
	}
	return cache
}
