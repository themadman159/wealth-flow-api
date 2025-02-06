package transactionsvc

func (s *TransactionService) Delete(username string, id int) error {

	err := s.TransactionRepository.Delete(username, id)
	if err != nil {
		return err
	}

	return nil
}
