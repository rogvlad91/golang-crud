package candidate

import (
	"context"
	"homework-7/internal/svc/candidates"
)

type CandidatesController struct {
	svc candidates.CandidateProcessor
}

func NewCandidatesController(svc candidates.CandidateProcessor) *CandidatesController {
	return &CandidatesController{svc: svc}
}

type CandidatesCLIProcessor interface {
	Create(ctx context.Context, fullName string, age string, wantedJob string, wantedSalary string) error
	GetById(ctx context.Context, id string) error
	GetAll(ctx context.Context) error
	Update(ctx context.Context, id string, wantedJob string, wantedSalary string) error
	Delete(ctx context.Context, id string) error
}
