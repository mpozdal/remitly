package swift

import (
	"database/sql"
	"net/http"
	"strings"

	"mpozdal/remitly/db"
	"mpozdal/remitly/types"
	"mpozdal/remitly/utils"
)

type SwiftService struct {
	dbm *db.DBManager
}

func NewSwiftService(dbm *db.DBManager) *SwiftService {
	return &SwiftService{dbm: dbm}
}

func (s *SwiftService) GetDataBySwiftCode(swiftCode string) (*types.BankWithBranchesReponse, error) {
	var response types.BankWithBranchesReponse

	bank, err := s.dbm.GetBankBySwiftCode(swiftCode)
	if err != nil {
		return nil, err
	}
	response = types.BankWithBranchesReponse{
		Address:       bank.Address,
		BankName:      bank.BankName,
		CountryISO2:   bank.CountryISO2,
		CountryName:   bank.CountryName,
		IsHeadquarter: bank.IsHeadquarter,
		SwiftCode:     bank.SwiftCode,
	}
	if bank.IsHeadquarter {
		branches, err := s.dbm.GetBranchesByHQSwiftCode(swiftCode)
		branchesReponse := utils.ReturnBankReponse(branches)
		if err != nil {
			return nil, err
		}
		response = types.BankWithBranchesReponse{
			Address:       bank.Address,
			BankName:      bank.BankName,
			CountryISO2:   bank.CountryISO2,
			CountryName:   bank.CountryName,
			IsHeadquarter: bank.IsHeadquarter,
			SwiftCode:     bank.SwiftCode,
			Branches:      branchesReponse,
		}
	}

	return &response, nil

}

func (s *SwiftService) GetDataByCountry(countryISO2code string) (*types.ReponseByCountry, error) {

	country, err := s.dbm.GetCountry(countryISO2code)
	if err != nil {
		return nil, err
	}
	banks, err := s.dbm.GetBanksByCountry(countryISO2code)

	if err != nil {
		return nil, err
	}
	banksResponse := utils.ReturnBankReponse(banks)
	return &types.ReponseByCountry{
		CountryISO2: country.CountryISO2,
		CountryName: country.CountryName,
		Data:        banksResponse,
	}, err

}
func (s *SwiftService) AddNewData(paylaod types.AddBankPayload) types.Response {

	bank := types.Bank{
		SwiftCode:            paylaod.SwiftCode,
		BankName:             paylaod.BankName,
		CountryISO2:          paylaod.CountryISO2,
		IsHeadquarter:        paylaod.IsHeadquarter,
		HeadquarterSwiftCode: sql.NullString{Valid: false},
	}
	country := types.Country{CountryISO2: strings.ToUpper(paylaod.CountryISO2), CountryName: strings.ToUpper(paylaod.CountryName)}

	err := s.dbm.AddCountry(country)
	if err != nil {
		return types.Response{
			Message: "Error adding country",
			Status:  http.StatusInternalServerError,
		}
	}

	isAdded, err := s.dbm.AddBank(bank)
	if err != nil {
		return types.Response{
			Message: "Error adding country",
			Status:  http.StatusInternalServerError,
		}
	}
	if !isAdded {
		return types.Response{
			Message: "Swift Code already exists",
			Status:  http.StatusConflict,
		}
	}
	return types.Response{Message: "Successfully added data", Status: http.StatusOK}
}

func (s *SwiftService) DeleteData(code string) types.Response {
	isDeleted, err := s.dbm.DeleteBank(code)
	if err != nil {
		return types.Response{Message: "Error deleting data", Status: http.StatusInternalServerError}
	}
	if !isDeleted {
		return types.Response{Message: "Data not found", Status: http.StatusNotFound}
	}

	return types.Response{Message: "Successfully deleted data", Status: http.StatusOK}

}
