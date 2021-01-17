package forecasting

import (
	"strings"

	"github.com/yudhasubki/forecast/pkg/bmkg"
)

type Service struct {
	BMKG       *bmkg.BMKG
	Repository *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		BMKG:       bmkg.NewForecast(),
		Repository: repo,
	}
}

func (s *Service) ResolveProvinces() []Province {
	provinces := make([]Province, 0)
	for k, v := range s.BMKG.GetProvince() {
		provinces = append(provinces, Province{
			ID:   v,
			Name: k,
		})
	}
	return provinces
}

func (s *Service) ResolveAreas(province string) []Area {
	areasByProvince := s.BMKG.GetArea(province)

	areas := make([]Area, 0)
	for _, a := range areasByProvince {
		areaKey := strings.ToLower(strings.ReplaceAll(a.Name, " ", "-"))
		areas = append(areas, Area{
			ID:   areaKey,
			Name: a.Name,
		})
	}

	return areas
}

func (s *Service) ResolveForecastingByProvinceAndArea(province, area string) (Forecasting, error) {
	found, forecasting := s.Repository.ResolveForecastingMap(province)
	if !found {
		bmkg, err := s.BMKG.GetForecasting(province)
		if err != nil {
			return Forecasting{}, err
		}

		c := NewForecastingFromBMKG(bmkg)
		err = s.Repository.CreateForecastingMap(province, c)
		if err != nil {
			return Forecasting{}, nil
		}

		return c[area], nil
	}

	return forecasting[area], nil
}

func (s *Service) ResolveForecastingByProvince(province string) ([]Forecasting, error) {
	found, forecasting := s.Repository.ResolveForecastingArray(province)
	if !found {
		bmkg, err := s.BMKG.GetForecasting(province)
		if err != nil {
			return nil, err
		}

		c := NewForecastingToArrayFromMap(NewForecastingFromBMKG(bmkg))
		err = s.Repository.CreateForecastingArray(province, c)
		if err != nil {
			return nil, err
		}

		return c, nil
	}

	return forecasting, nil
}
