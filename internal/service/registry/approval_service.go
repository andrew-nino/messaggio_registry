package service

import "github.com/andrew-nino/messaggio/internal/domain/models"

func (s *ApplicationServices) Approve(answer models.Answer) error {

	err := s.approval.SetApproval(answer)
	if err != nil {
		s.log.Error("Error set approving to BD: ", err)
		return err
	}
	return nil
}
