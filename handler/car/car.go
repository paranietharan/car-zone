package car

import (
	"car-zone/models"
	"car-zone/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler {
	return &CarHandler{
		service: service,
	}
}

func (h *CarHandler) GetCarByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")

	car, err := h.service.GetCarByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var carReq models.CarRequest

	if err := json.NewDecoder(r.Body).Decode(&carReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car, err := h.service.CreateCar(ctx, carReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(body)
	if err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")
	var carReq models.Car

	if err := json.NewDecoder(r.Body).Decode(&carReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car, err := h.service.UpdateCar(ctx, id, carReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")

	err := h.service.DeleteCar(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	brand := r.URL.Query().Get("brand")

	cars, err := h.service.GetCarByBrand(ctx, brand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}
