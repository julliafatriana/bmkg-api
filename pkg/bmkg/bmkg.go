package bmkg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	areaMap map[string][]Territory
)

type BMKG struct {
	XMLName  xml.Name `xml:"data"`
	Forecast Forecast `xml:"forecast"`
}

type Forecast struct {
	Domain string `xml:"domain,attr"`
	Issue  Issue  `xml:"issue"`
	Area   []Area `xml:"area"`
}

type Issue struct {
	Timestamp string `xml:"timestamp"`
	Year      string `xml:"year"`
	Day       string `xml:"day"`
	Month     string `xml:"month"`
	Hour      string `xml:"hour"`
	Minute    string `xml:"minute"`
	Second    string `xml:"second"`
}

type Area struct {
	ID          string      `xml:"id,attr"`
	Latitude    string      `xml:"latitude,attr"`
	Longitude   string      `xml:"longitude,attr"`
	Coordinate  string      `xml:"coordinate,attr"`
	Type        string      `xml:"type,attr"`
	Region      string      `xml:"region,attr"`
	Level       string      `xml:"level,attr"`
	Description string      `xml:"description,attr"`
	Domain      string      `xml:"domain,attr"`
	Tags        string      `xml:"tags,attr"`
	Names       []Name      `xml:"name"`
	Parameters  []Parameter `xml:"parameter"`
}

type Parameter struct {
	ID          string      `xml:"id,attr"`
	Description string      `xml:"description,attr"`
	Type        string      `xml:"type,attr"`
	Timeranges  []Timerange `xml:"timerange"`
}

type Timerange struct {
	Type     string           `xml:"type,attr"`
	Hour     string           `xml:"h,attr"`
	Datetime string           `xml:"datetime,attr"`
	Value    []TimerangeValue `xml:"value"`
}

type TimerangeValue struct {
	Unit  string `xml:"unit,attr"`
	Value string `xml:",chardata"`
}

type Name struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type Territory struct {
	Name string `json:"name"`
}

func initArea(b BMKG) error {
	provinces := b.GetProvince()

	for _, p := range provinces {
		res, err := b.GetForecasting(p)
		if err != nil {
			return err
		}

		t := make([]Territory, 0)
		for _, r := range res.Forecast.Area {
			t = append(t, Territory{Name: r.Description})
		}
		areaMap[p] = t
	}
	return nil
}

func NewForecast() *BMKG {
	b := BMKG{}
	areaMap = make(map[string][]Territory)
	err := initArea(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finish Init Area")
	return &b
}

// GetForecasting is function retrieve Forecasting based on province name
func (b *BMKG) GetForecasting(province string) (BMKG, error) {
	resp, err := http.Get(fmt.Sprintf("https://data.bmkg.go.id/datamkg/MEWS/DigitalForecast/DigitalForecast-%s.xml", province))
	if err != nil {
		return BMKG{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return BMKG{}, err
	}

	result := BMKG{}

	err = xml.Unmarshal(body, &result)
	if err != nil {
		return BMKG{}, err
	}

	return result, nil
}

func (b *BMKG) GetArea(province string) []Territory {
	return areaMap[province]
}

func (b *BMKG) GetProvince() map[string]string {
	province := make(map[string]string)
	province["Aceh"] = "Aceh"
	province["Bali"] = "Bali"
	province["Banten"] = "Banten"
	province["Bengkulu"] = "Bengkulu"
	province["DI Yogyakarta"] = "DIYogyakarta"
	province["DKI Jakarta"] = "DKIJakarta"
	province["Gorontalo"] = "Gorontalo"
	province["Jambi"] = "Jambi"
	province["Jawa Barat"] = "JawaBarat"
	province["Jawa Tengah"] = "JawaTengah"
	province["Jawa Timur"] = "JawaTimur"
	province["Kalimantan Barat"] = "KalimantanBarat"
	province["Kalimantan Selatan"] = "KalimantanSelatan"
	province["Kalimantan Tengah"] = "KalimantanTengah"
	province["Kalimantan Timur"] = "KalimantanTimur"
	province["Kalimantan Utara"] = "KalimantanUtara"
	province["Bangka Belitung"] = "BangkaBelitung"
	province["Kepulauan Riau"] = "KepulauanRiau"
	province["Lampung"] = "Lampung"
	province["Maluku"] = "Maluku"
	province["Maluku Utara"] = "MalukuUtara"
	province["Nusa Tenggara Barat"] = "NusaTenggaraBarat"
	province["Nusa Tenggara Timur"] = "NusaTenggaraTimur"
	province["Papua"] = "Papua"
	province["Papua Barat"] = "PapuaBarat"
	province["Riau"] = "Riau"
	province["Sulawesi Barat"] = "SulawesiBarat"
	province["Sulawesi Selatan"] = "SulawesiSelatan"
	province["Sulawesi Tengah"] = "SulawesiTengah"
	province["Sulawesi Tenggara"] = "SulawesiTenggara"
	province["Sulawesi Utara"] = "SulawesiUtara"
	province["Sumatera Barat"] = "SumateraBarat"
	province["Sumatera Selatan"] = "SumateraSelatan"
	province["Sumatera Utara"] = "SumateraUtara"

	return province
}
