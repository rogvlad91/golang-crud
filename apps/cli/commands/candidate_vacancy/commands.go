package candidate_vacancy

import (
	"context"
	"fmt"
	candidateVacancy "homework-7/apps/cli/controllers/candidate_vacancy"
	"os"
)

type CreateResponseCommand struct {
	VacancyCandidateController candidateVacancy.CandidateVacancyCLIProcessor
}

func (c *CreateResponseCommand) Validate(args []string) {
	if len(args) != 5 || args[2] != "-vacancy_id" || args[4] != "-candidate_id" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *CreateResponseCommand) Execute(ctx context.Context, args []string) error {
	return c.VacancyCandidateController.Create(ctx, args[3], args[5])
}

func (c *CreateResponseCommand) Usage() {
	fmt.Println("Command usage: create_response -vacancy_id <id> -candidate_id <id>")
}

type DeleteResponseCommand struct {
	VacancyCandidateController candidateVacancy.CandidateVacancyCLIProcessor
}

func (c *DeleteResponseCommand) Validate(args []string) {
	if len(args) != 5 || args[2] != "-vacancy_id" || args[4] != "-candidate_id" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *DeleteResponseCommand) Execute(ctx context.Context, args []string) error {
	return c.VacancyCandidateController.Create(ctx, args[3], args[5])
}

func (c *DeleteResponseCommand) Usage() {
	fmt.Println("Command usage: delete_response -vacancy_id <id> -candidate_id <id>")
}

type GetCandidatesByVacancyCommand struct {
	VacancyCandidateController candidateVacancy.CandidateVacancyCLIProcessor
}

func (c *GetCandidatesByVacancyCommand) Validate(args []string) {
	if len(args) != 3 || args[2] != "-vacancy_id" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *GetCandidatesByVacancyCommand) Execute(ctx context.Context, args []string) error {
	return c.VacancyCandidateController.GetCandidatesByVacancyId(ctx, args[3])
}

func (c *GetCandidatesByVacancyCommand) Usage() {
	fmt.Println("Command usage: get_candidates_by_vacancy -vacancy_id <id>")
}

type GetVacanciesByCandidateCommand struct {
	VacancyCandidateController candidateVacancy.CandidateVacancyCLIProcessor
}

func (c *GetVacanciesByCandidateCommand) Validate(args []string) {
	if len(args) != 3 || args[2] != "-candidate_id" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *GetVacanciesByCandidateCommand) Execute(ctx context.Context, args []string) error {
	return c.VacancyCandidateController.GetVacanciesByCandidate(ctx, args[3])
}

func (c *GetVacanciesByCandidateCommand) Usage() {
	fmt.Println("Command usage: get_vacancies_by_candidate -candidate_id <id>")
}
