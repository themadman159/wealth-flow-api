package transactionsvc

import "go-api/types"

func (s *TransactionService) Create(username string, req types.TransactionRequest) error {

	err := s.TransactionRepository.Create(username, req)
	if err != nil {
		return err
	}

	return nil
}
