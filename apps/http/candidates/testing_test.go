package candidate_controller

import (
	"github.com/golang/mock/gomock"
	mockcandidatesvc "homework-7/internal/svc/candidates/mocks"
	"testing"
)

type CandidateControllerFixture struct {
	ctrl       *gomock.Controller
	mockSvc    *mockcandidatesvc.MockCandidateProcessor
	controller *CandidateController
}

func setUp(t *testing.T) CandidateControllerFixture {
	ctrl := gomock.NewController(t)

	mockSvc := mockcandidatesvc.NewMockCandidateProcessor(ctrl)
	controller := NewCandidateController(mockSvc)
	return CandidateControllerFixture{mockSvc: mockSvc, ctrl: ctrl, controller: controller}
}

func (c *CandidateControllerFixture) tearDown() {
	c.ctrl.Finish()
}
