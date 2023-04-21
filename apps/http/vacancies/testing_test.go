package vacancies_controller

import (
	"github.com/golang/mock/gomock"
	mockvacancysvc "homework-7/internal/svc/vacancies/mocks"
	"testing"
)

type VacancyControllerFixture struct {
	ctrl       *gomock.Controller
	mockSvc    *mockvacancysvc.MockVacancyProcessor
	controller *VacancyController
}

func setUp(t *testing.T) VacancyControllerFixture {
	ctrl := gomock.NewController(t)

	mockSvc := mockvacancysvc.NewMockVacancyProcessor(ctrl)
	controller := NewVacancyController(mockSvc)
	return VacancyControllerFixture{mockSvc: mockSvc, ctrl: ctrl, controller: controller}
}

func (c *VacancyControllerFixture) tearDown() {
	c.ctrl.Finish()
}
