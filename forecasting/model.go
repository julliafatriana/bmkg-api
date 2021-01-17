package forecasting

import (
	"fmt"
	"strings"

	"github.com/yudhasubki/forecast/pkg/bmkg"
)

type IDParameter string

const (
	windSpeed      IDParameter = "ws"
	temperature    IDParameter = "t"
	humidity       IDParameter = "hu"
	maxHumidity    IDParameter = "humax"
	minHumidity    IDParameter = "humin"
	maxTemperature IDParameter = "tmax"
	minTemperature IDParameter = "tmin"
	weather        IDParameter = "weather"
	windDirection  IDParameter = "wd"
)

type Forecasting struct {
	Territory       Territory   `json:"territory"`
	Humidities      []Parameter `json:"humidity"`
	Temperatures    []Parameter `json:"temperature"`
	WindSpeeds      []Parameter `json:"windSpeed"`
	MinTemperatures []Parameter `json:"minTemperature"`
	MaxTemperatures []Parameter `json:"maxTemperature"`
	MinHumidities   []Parameter `json:"minHumidity"`
	MaxHumidities   []Parameter `json:"maxHumidity"`
	Weather         []Parameter `json:"weather"`
	WindDirection   []Parameter `json:"windDirection"`
}

type Territory struct {
	Name       string `json:"name"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Coordinate string `json:"coordinate"`
}

type Parameter struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Hour        string   `json:"hour"`
	Date        string   `json:"date"`
	Time        string   `json:"time"`
	Scheme      []Scheme `json:"scheme"`
}

type Scheme struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func NewForecastingToArrayFromMap(fs map[string]Forecasting) []Forecasting {
	forecasting := make([]Forecasting, 0)
	for _, f := range fs {
		forecasting = append(forecasting, f)
	}

	return forecasting
}

// NewForecastingFromBMKG transform to internal Forecasting from BMKG
func NewForecastingFromBMKG(bmkg bmkg.BMKG) map[string]Forecasting {
	area := make(map[string]Forecasting)
	for _, a := range bmkg.Forecast.Area {
		f := Forecasting{}
		for _, p := range a.Parameters {
			switch p.ID {
			case string(humidity):
				f.Humidities = NewParameter(p)
				break
			case string(temperature):
				f.Temperatures = NewParameter(p)
				break
			case string(windSpeed):
				f.WindSpeeds = NewParameter(p)
				break
			case string(maxHumidity):
				f.MaxHumidities = NewParameter(p)
				break
			case string(minHumidity):
				f.MinHumidities = NewParameter(p)
				break
			case string(maxTemperature):
				f.MaxTemperatures = NewParameter(p)
				break
			case string(minTemperature):
				f.MinTemperatures = NewParameter(p)
				break
			case string(weather):
				f.Weather = NewParameter(p)
				break
			case string(windDirection):
				f.WindDirection = NewParameter(p)
				break
			}
		}
		territory := Territory{
			Name:       a.Description,
			Latitude:   a.Latitude,
			Longitude:  a.Longitude,
			Coordinate: a.Coordinate,
		}
		f.Territory = territory
		areaKey := strings.ToLower(strings.ReplaceAll(a.Description, " ", "-"))
		area[areaKey] = f
	}

	return area
}

// NewParameter is function to construct/transform Parameter from BKMG to Internal Parameter
func NewParameter(parameter bmkg.Parameter) []Parameter {
	parameters := make([]Parameter, 0)

	for _, t := range parameter.Timeranges {
		schemes := make([]Scheme, 0)
		for _, s := range t.Value {
			schemes = append(schemes, Scheme{Label: s.Unit, Value: s.Value})
		}

		parameters = append(parameters, Parameter{
			Description: parameter.Description,
			Type:        parameter.Type,
			Hour:        t.Datetime[8:10],
			Date:        fmt.Sprintf("%s-%s-%s", t.Datetime[0:4], t.Datetime[4:6], t.Datetime[6:8]),
			Time:        fmt.Sprintf("%s:%s", t.Datetime[8:10], t.Datetime[10:12]),
			Scheme:      schemes,
		})
	}

	return parameters
}

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Area struct {
	ID   string `json:"id"`
	Name string `json:"area"`
}

type AreaRequest struct {
	Province string `json:"province"`
}

type ForecastingRequest struct {
	ProvinceID string `json:"provinceId"`
	AreaID     string `json:"areaId"`
}
