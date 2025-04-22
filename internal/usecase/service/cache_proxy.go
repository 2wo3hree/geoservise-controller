package service

import (
	"context"
	"encoding/json"
	"fmt"
	"geoservise-jwt/internal/infrastructure/metrics"
	"geoservise-jwt/internal/model"
	"github.com/redis/go-redis/v9"
	"time"
)

type CachedGeoService struct {
	real  GeoService
	cache *redis.Client
	ttl   time.Duration
}

func NewCachedGeoService(real GeoService, cache *redis.Client, ttl time.Duration) *CachedGeoService {
	return &CachedGeoService{
		real:  real,
		cache: cache,
		ttl:   ttl,
	}
}

func (c *CachedGeoService) Search(ctx context.Context, query string) ([]*model.Address, error) {
	start := time.Now()

	key := fmt.Sprintf("search:%s", query)
	//fmt.Println("🔍 Ищу в Redis ключ:", key)
	cached, err := c.cache.Get(ctx, key).Result()
	if err == nil {
		//fmt.Println("⚡ Найдено в кэше:", key)
		var result []*model.Address
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result, nil
		}
	}
	metrics.CacheDuration.WithLabelValues("get").Observe(time.Since(start).Seconds())
	//fmt.Println("Не найдено в кэше, идём в API:", key)

	result, err := c.real.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(result)
	if err != nil {
		//fmt.Println("Ошибка маршалинга:", err)
		return nil, err
	}
	err = c.cache.Set(ctx, key, data, c.ttl).Err()
	if err != nil {
		//fmt.Println("Ошибка записи в Redis:", err)
	} else {
		//fmt.Println("Сохранили в Redis:", key)
	}

	return result, nil
}

func (c *CachedGeoService) Geocode(ctx context.Context, latStr, lngStr string) ([]*model.Address, error) {
	start := time.Now()

	key := fmt.Sprintf("geocode:%s:%s", latStr, lngStr)
	cached, err := c.cache.Get(ctx, key).Result()
	if err == nil {
		var result []*model.Address
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result, nil
		}
	}
	metrics.CacheDuration.WithLabelValues("get").Observe(time.Since(start).Seconds())

	result, err := c.real.Geocode(ctx, latStr, lngStr)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = c.cache.Set(ctx, key, data, c.ttl).Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}
