package automobile

func (s *AutomobileStorage) DeleteAuto(id int) error {
	query := `DELETE FROM automobiles WHERE id = $1`
	_, err := s.DB.Exec(query, id)
	return err
}
