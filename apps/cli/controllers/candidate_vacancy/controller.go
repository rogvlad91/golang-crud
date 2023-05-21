package candidate_vacancy

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-crud/internal/svc/candidate_vacancies"
)

func (c CandidateVacancyController) Create(ctx context.Context, vacancyId string, candidateId string) error {
	id, err := c.candidateVacancySvc.Create(ctx, candidate_vacancies.CreateDto{VacancyId: vacancyId, CandidateId: candidateId})

	if err != nil {
		return err
	}

	fmt.Printf("Create vacancy response with id %s \n", id)
	return nil
}

func (c CandidateVacancyController) DeleteResponseForVacancy(ctx context.Context, vacancyId string, candidateId string) error {
	err := c.candidateVacancySvc.DeleteResponseForVacancy(ctx, vacancyId, candidateId)
	if err != nil {
		return err
	}

	fmt.Printf("Delete response for vacancy %s of candidate %s \n", vacancyId, candidateId)
	return nil

}

func (c CandidateVacancyController) GetCandidatesByVacancyId(ctx context.Context, vacancyId string) error {
	ids, err := c.candidateVacancySvc.GetCandidatesByVacancyId(ctx, vacancyId)
	if err != nil {
		return err
	}

	fmt.Printf("Candidates for vacancy %s \n", vacancyId)

	for _, candidateId := range ids {
		candidate, candidateErr := c.candidateSvc.GetById(ctx, *candidateId)
		if candidateErr != nil {
			return candidateErr
		}

		parsedCandidate, _ := json.Marshal(candidate)

		fmt.Printf("Candidate %s \n", parsedCandidate)
	}

	return nil
}

func (c CandidateVacancyController) GetVacanciesByCandidate(ctx context.Context, candidateId string) error {
	ids, err := c.candidateVacancySvc.GetVacanciesByCandidate(ctx, candidateId)
	if err != nil {
		return err
	}

	fmt.Printf("Vacancies for candidate %s \n", candidateId)

	for _, vacancyId := range ids {
		vacancy, vacancyErr := c.vacancySvc.GetById(ctx, *vacancyId)
		if vacancyErr != nil {
			return vacancyErr
		}

		parsedVacancy, _ := json.Marshal(vacancy)

		fmt.Printf("Vacancy %s \n", parsedVacancy)
	}

	return nil
}
