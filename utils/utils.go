package utils

import (
	"encoding/json"
	"fmt"
	"mpozdal/remitly/types"
	"net/http"
)

func SortBanks(banks []types.Bank) []types.Bank {
	var headquarters []types.Bank
	var branches []types.Bank

	for _, bank := range banks {
		if bank.IsHeadquarter {
			headquarters = append(headquarters, bank)
		} else {
			branches = append(branches, bank)
		}
	}

	return append(headquarters, branches...)
}

func WriteJSONMessage(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	var msg = map[string]any{"message": v}
	return json.NewEncoder(w).Encode(msg)
}
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func ReturnBankReponse(banks []types.Bank) []types.BankResponse {

	var bankResponses []types.BankResponse

	for _, bank := range banks {
		bankResponses = append(bankResponses, types.BankResponse{
			SwiftCode:     bank.SwiftCode,
			BankName:      bank.BankName,
			Address:       bank.Address,
			CountryISO2:   bank.CountryISO2,
			CountryName:   bank.CountryName,
			IsHeadquarter: bank.IsHeadquarter,
		})
	}
	return bankResponses

}
