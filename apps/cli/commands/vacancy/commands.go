package vacancy

import (
	"context"
	"fmt"
	"golang-crud/apps/cli/controllers/vacancy"
	"os"
)

type CreateVacancyCommand struct {
	VacancyController vacancy.VacanciesCLIProcessor
}

func (c *CreateVacancyCommand) Validate(args []string) {
	if len(args) != 5 || args[2] != "-vacancy_title" || args[4] != "-vacancy_salary" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *CreateVacancyCommand) Execute(ctx context.Context, args []string) error {
	return c.VacancyController.Create(ctx, args[3], args[5])
}

func (c *CreateVacancyCommand) Usage() {
	fmt.Println("Command usage: create_vacancy -vacancy_title <title> -vacancy_salary <salary>")
}

type UpdateVacancyCommand struct {
	VacancyController vacancy.VacanciesCLIProcessor
}

func (c *UpdateVacancyCommand) Validate(args []string) {
	if len(args) != 7 || args[2] != "-vacancy_id" || args[4] != "-vacancy_title" || args[6] != "-vacancy_salary" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *UpdateVacancyCommand) Execute(ctx context.Context, args []string) error {
	return c.VacancyController.Update(ctx, args[3], args[5], args[7])
}

func (c *UpdateVacancyCommand) Usage() {
	fmt.Println("Command usage: update_vacancy -vacancy_id <id> -vacancy_title <title> -vacancy_salary <salary>")
}

type DeleteVacancyCommand struct {
	VacancyController vacancy.VacanciesCLIProcessor
}

func (c *DeleteVacancyCommand) Validate(args []string) {
	if len(args) != 3 || args[2] != "-vacancy_id" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *DeleteVacancyCommand) Execute(ctx context.Context, args []string) error {
	return c.VacancyController.Delete(ctx, args[3])
}

func (c *DeleteVacancyCommand) Usage() {
	fmt.Println("Command usage: delete_vacancy -vacancy_id <id>")
}

type GetVacancyCommand struct {
	VacancyController vacancy.VacanciesCLIProcessor
}

func (c *GetVacancyCommand) Validate(args []string) {
	if len(args) != 3 || args[2] != "-vacancy_id" {
		c.Usage()
		os.Exit(1)
	}
}

func (c *GetVacancyCommand) Execute(ctx context.Context, args []string) error {
	return c.VacancyController.GetById(ctx, args[3])
}

func (c *GetVacancyCommand) Usage() {
	fmt.Println("Command usage: get_vacancy -vacancy_id <id>")
}

type GetVacanciesCommand struct {
	VacancyController vacancy.VacanciesCLIProcessor
}

func (c *GetVacanciesCommand) Validate(_args []string) {
}

func (c *GetVacanciesCommand) Execute(ctx context.Context, args []string) error {
	return c.VacancyController.GetAll(ctx)
}

func (c *GetVacanciesCommand) Usage() {
	fmt.Println("Command usage: get_vacancies")
}
