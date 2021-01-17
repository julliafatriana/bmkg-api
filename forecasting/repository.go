package forecasting

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/yudhasubki/forecast/config"
)

const (
	ForecastingMapPrefixCache   = "area:map"
	ForecastingArrayPrefixCache = "area:province"
)

type Repository struct {
	Cache *cache.Cache
}

func NewRepository(cfg config.Config) *Repository {
	return &Repository{
		Cache: cache.New(time.Duration(cfg.DefaultExpirationDuration)*time.Minute, time.Duration(cfg.PurgeExpiredItemsDuration)*time.Minute),
	}
}

func (r *Repository) CreateForecastingMap(province string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	r.Cache.Set(fmt.Sprintf("%s:%s", ForecastingMapPrefixCache, strings.ToLower(province)), string(b), cache.DefaultExpiration)

	return nil
}

func (r *Repository) CreateForecastingArray(province string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	r.Cache.Set(fmt.Sprintf("%s:%s", ForecastingArrayPrefixCache, strings.ToLower(province)), string(b), cache.DefaultExpiration)

	return nil
}

func (r *Repository) ResolveForecastingMap(province string) (bool, map[string]Forecasting) {
	if f, exist := r.Cache.Get(fmt.Sprintf("%s:%s", ForecastingMapPrefixCache, strings.ToLower(province))); exist {
		forecasting := make(map[string]Forecasting)
		err := json.Unmarshal([]byte(f.(string)), &forecasting)
		if err != nil {
			return false, nil
		}

		return exist, forecasting
	}

	return false, nil
}

func (r *Repository) ResolveForecastingArray(province string) (bool, []Forecasting) {
	if f, exist := r.Cache.Get(fmt.Sprintf("%s:%s", ForecastingArrayPrefixCache, strings.ToLower(province))); exist {
		forecasting := make([]Forecasting, 0)
		err := json.Unmarshal([]byte(f.(string)), &forecasting)
		if err != nil {
			return false, nil
		}

		return exist, forecasting
	}

	return false, nil
}
