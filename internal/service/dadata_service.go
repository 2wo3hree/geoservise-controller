package service

import (
	"context"
	"fmt"
	"geoservise-jwt/internal/model"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"strconv"
	"strings"
)

type GeoService interface {
	Search(ctx context.Context, query string) ([]*model.Address, error)
	Geocode(ctx context.Context, latStr, lngStr string) ([]*model.Address, error)
}

type Service struct {
	api *suggest.Api
}

func NewService(api *suggest.Api) *Service {
	return &Service{
		api: api,
	}
}

func (s *Service) Search(ctx context.Context, query string) ([]*model.Address, error) {
	if strings.TrimSpace(query) == "" {
		return nil, fmt.Errorf("query пустой")
	}

	params := suggest.RequestParams{
		Query: query,
	}
	suggestions, err := s.api.Address(ctx, &params)
	if err != nil {
		return nil, err
	}

	addresses := make([]*model.Address, 0, len(suggestions))
	for _, sugg := range suggestions {
		addresses = append(addresses, &model.Address{
			Source:       query,
			Result:       sugg.Value,
			PostalCode:   sugg.Data.PostalCode,
			Country:      sugg.Data.Country,
			Region:       sugg.Data.Region,
			CityArea:     sugg.Data.CityArea,
			CityDistrict: sugg.Data.CityDistrict,
			Street:       sugg.Data.Street,
			House:        sugg.Data.House,
			GeoLat:       sugg.Data.GeoLat,
			GeoLon:       sugg.Data.GeoLon,
			QCGeo:        sugg.Data.QualityCodeGeoRaw,
		})
	}

	return addresses, nil
}

func (s *Service) Geocode(ctx context.Context, latStr, lngStr string) ([]*model.Address, error) {
	lat, err1 := strconv.ParseFloat(latStr, 64)
	lng, err2 := strconv.ParseFloat(lngStr, 64)
	if err1 != nil || err2 != nil || lat == 0 || lng == 0 {
		return nil, fmt.Errorf("невалидные координаты")
	}

	params := suggest.GeolocateParams{
		Lat: fmt.Sprintf("%f", lat),
		Lon: fmt.Sprintf("%f", lng),
	}
	suggestions, err := s.api.GeoLocate(ctx, &params)
	if err != nil {
		return nil, err
	}

	addresses := make([]*model.Address, 0, len(suggestions))
	for _, sugg := range suggestions {
		addresses = append(addresses, &model.Address{
			Result:       sugg.Value,
			PostalCode:   sugg.Data.PostalCode,
			Country:      sugg.Data.Country,
			Region:       sugg.Data.Region,
			CityArea:     sugg.Data.CityArea,
			CityDistrict: sugg.Data.CityDistrict,
			Street:       sugg.Data.Street,
			House:        sugg.Data.House,
			GeoLat:       sugg.Data.GeoLat,
			GeoLon:       sugg.Data.GeoLon,
			QCGeo:        sugg.Data.QualityCodeGeoRaw,
		})
	}

	return addresses, nil
}
