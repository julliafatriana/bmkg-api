package forecasting

import (
	"encoding/json"
	"net/http"

	"github.com/yudhasubki/forecast/libs/response"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.WithMessage(w, http.StatusOK, "Health")
}

func (h *Handler) ResolveProvinces(w http.ResponseWriter, r *http.Request) {
	response.WithJSON(w, http.StatusOK, h.Service.ResolveProvinces())
}

func (h *Handler) ResolveAreas(w http.ResponseWriter, r *http.Request) {
	var req AreaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WithMessage(w, http.StatusBadRequest, "")
		return
	}

	response.WithJSON(w, http.StatusOK, h.Service.ResolveAreas(req.Province))
}

func (h *Handler) ResolveForecastingByProvinceAndArea(w http.ResponseWriter, r *http.Request) {
	var req ForecastingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WithMessage(w, http.StatusBadRequest, "")
		return
	}

	data, err := h.Service.ResolveForecastingByProvinceAndArea(req.ProvinceID, req.AreaID)
	if err != nil {
		response.WithMessage(w, http.StatusBadRequest, "")
		return
	}
	response.WithJSON(w, http.StatusOK, data)
}

func (h *Handler) ResolveForecastingByProvince(w http.ResponseWriter, r *http.Request) {
	var req ForecastingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WithMessage(w, http.StatusBadRequest, "")
		return
	}

	data, err := h.Service.ResolveForecastingByProvince(req.ProvinceID)
	if err != nil {
		response.WithMessage(w, http.StatusBadRequest, "")
		return
	}
	response.WithJSON(w, http.StatusOK, data)
}
