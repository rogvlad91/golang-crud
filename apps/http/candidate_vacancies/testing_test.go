package candidate_vacancies_controller

import (
	"github.com/golang/mock/gomock"
	mock_candidate_vacancy_svc "homework-7/internal/svc/candidate_vacancies/mocks"
	"testing"
)

type VacancyControllerFixture struct {
	ctrl       *gomock.Controller
	mockSvc    *mock_candidate_vacancy_svc.MockCandidateVacanciesProcessor
	controller *CandidateVacanciesController
}

func setUp(t *testing.T) VacancyControllerFixture {
	ctrl := gomock.NewController(t)

	mockSvc := mock_candidate_vacancy_svc.NewMockCandidateVacanciesProcessor(ctrl)
	controller := NewCandidateVacanciesController(mockSvc)
	return VacancyControllerFixture{mockSvc: mockSvc, ctrl: ctrl, controller: controller}
}

func (c *VacancyControllerFixture) tearDown() {
	c.ctrl.Finish()
}
