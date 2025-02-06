package transactionsvc

import (
	"go-api/constants/transactionconstants"
	"go-api/types"
)

func (s *TransactionService) GetAll(username string) (*types.TransactionGetAllResponse, error) {

	transcaction, err := s.TransactionRepository.GetAll(username)
	if err != nil {
		return nil, err
	}

	var response []types.TransactionGetAll
	var income, expense, total int64
	for _, tx := range transcaction {
		resp := types.TransactionGetAll{
			ID:              tx.ID,
			Type:            tx.Type,
			Category:        tx.Category,
			Amount:          tx.Amount,
			Description:     tx.Description,
			TransactionDate: tx.TransactionDate,
		}
		response = append(response, resp)

		if tx.Type == transactionconstants.Transactionconstants.INCOME {
			income += tx.Amount
		} else if tx.Type == transactionconstants.Transactionconstants.EXPENSE {
			expense += tx.Amount
		}
	}

	total = income - expense

	return &types.TransactionGetAllResponse{
		Transaction: response,
		Total:       total,
	}, nil
}
