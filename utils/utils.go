package utils

import (
	"mpozdal/remitly/types"
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