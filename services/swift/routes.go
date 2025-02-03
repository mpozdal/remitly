package swift

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"mpozdal/remitly/types"
	"mpozdal/remitly/utils"
)

type Handler struct {
	service *SwiftService
}

func NewHandler(service *SwiftService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/swift-codes/{swift-code}", h.getData).Methods(http.MethodGet)
	router.HandleFunc("/swift-codes/country/{countryISO2code}", h.getDataByCountry).Methods(http.MethodGet)
	router.HandleFunc("/swift-codes", h.addNewData).Methods(http.MethodPost)
	router.HandleFunc("/swift-codes/{swift-code}", h.deleteData).Methods(http.MethodDelete)

}

func (h *Handler) getData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	code, ok := vars["swift-code"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing swift-code parameter"))
		return
	}
	response, err := h.service.GetDataBySwiftCode(code)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, response)

}
func (h *Handler) getDataByCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	code, ok := vars["countryISO2code"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing iso2code parameter"))
		return
	}
	response, err := h.service.GetDataByCountry(code)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, response)
}
func (h *Handler) addNewData(w http.ResponseWriter, r *http.Request) {
	var payload types.AddBankPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := h.service.AddNewData(payload)

	utils.WriteJSONMessage(w, response.Status, response.Message)

}
func (h *Handler) deleteData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code, ok := vars["swift-code"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing swift-code parameter"))
		return
	}

	response := h.service.DeleteData(code)

	utils.WriteJSONMessage(w, response.Status, response.Message)
}
