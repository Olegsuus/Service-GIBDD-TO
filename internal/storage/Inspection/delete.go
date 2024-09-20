package Inspection

func (s *InspectionStorage) DeleteInsp(id int) error {
	query := `DELETE FROM inspections WHERE id = $1`
	_, err := s.DB.Exec(query, id)
	return err
}
