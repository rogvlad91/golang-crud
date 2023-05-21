package candidate

import (
	"context"
	"fmt"
	"golang-crud/apps/cli/controllers/candidate"
	"os"
)

type CreateCandidateCommand struct {
	CandidateController candidate.CandidatesCLIProcessor
}

func (c *CreateCandidateCommand) Validate(args []string) {
	if len(args) != 10 || args[2] != "-full_name" || args[4] != "-candidate_age" || args[6] != "-wanted_job" || args[8] != "-wanted_salary" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *CreateCandidateCommand) Execute(ctx context.Context, args []string) error {
	return c.CandidateController.Create(ctx, args[3], args[5], args[7], args[9])
}

func (c *CreateCandidateCommand) Usage() {
	fmt.Println("Command usage: create_candidate -full_name <candidate_full_name> -candidate_age <age> -wanted_job <job> -wanted_salary <salary>")
}

type UpdateCandidateCommand struct {
	CandidateController candidate.CandidatesCLIProcessor
}

func (c *UpdateCandidateCommand) Validate(args []string) {
	if len(args) != 7 || args[2] != "-candidate_id" || args[4] != "-wanted_job" || args[6] != "-wanted_salary" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *UpdateCandidateCommand) Execute(ctx context.Context, args []string) error {
	return c.CandidateController.Update(ctx, args[3], args[5], args[7])
}

func (c *UpdateCandidateCommand) Usage() {
	fmt.Println("Command usage: update_candidate -candidate_id <id> -wanted_job <job> -wanted_salary <salary>")
}

type DeleteCandidateCommand struct {
	CandidateController candidate.CandidatesCLIProcessor
}

func (c *DeleteCandidateCommand) Validate(args []string) {
	if len(args) != 3 || args[2] != "-candidate_id" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *DeleteCandidateCommand) Execute(ctx context.Context, args []string) error {
	return c.CandidateController.Delete(ctx, args[3])
}

func (c *DeleteCandidateCommand) Usage() {
	fmt.Println("Command usage: delete_candidate -candidate_id <id>")
}

type GetCandidateCommand struct {
	CandidateController candidate.CandidatesCLIProcessor
}

func (c *GetCandidateCommand) Validate(args []string) {
	if len(args) != 3 || args[2] != "-candidate_id" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *GetCandidateCommand) Execute(ctx context.Context, args []string) error {
	return c.CandidateController.GetById(ctx, args[3])
}

func (c *GetCandidateCommand) Usage() {
	fmt.Println("Command usage: get_candidate -candidate_id <id>")
}

type GetCandidatesCommand struct {
	CandidateController candidate.CandidatesCLIProcessor
}

func (c *GetCandidatesCommand) Validate(_args []string) {
}

func (c *GetCandidatesCommand) Execute(ctx context.Context, args []string) error {
	return c.CandidateController.GetAll(ctx)
}

func (c *GetCandidatesCommand) Usage() {
	fmt.Println("Command usage: get_candidates")
}
