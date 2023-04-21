package candidate_vacancy

import (
	"context"
	"homework-7/internal/svc/candidate_vacancies"
	"homework-7/internal/svc/candidates"
	"homework-7/internal/svc/vacancies"
)

type CandidateVacancyController struct {
	candidateVacancySvc candidate_vacancies.CandidateVacanciesProcessor
	candidateSvc        candidates.CandidateProcessor
	vacancySvc          vacancies.VacancyProcessor
}

func NewCandidateVacancyController(candidateVacancySvc candidate_vacancies.CandidateVacanciesProcessor, candidateSvc candidates.CandidateProcessor, vacancySvc vacancies.VacancyProcessor) *CandidateVacancyController {
	return &CandidateVacancyController{candidateVacancySvc: candidateVacancySvc, candidateSvc: candidateSvc, vacancySvc: vacancySvc}
}

type CandidateVacancyCLIProcessor interface {
	Create(ctx context.Context, vacancyId string, candidateId string) error
	DeleteResponseForVacancy(ctx context.Context, vacancyId string, candidateId string) error
	GetCandidatesByVacancyId(ctx context.Context, vacancyId string) error
	GetVacanciesByCandidate(ctx context.Context, candidateId string) error
}
