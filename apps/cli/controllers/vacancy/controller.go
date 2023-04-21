package vacancy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"homework-7/internal/svc/vacancies"
	"strconv"
)

func (c VacanciesController) Create(ctx context.Context, title string, salary string) error {
	parsedSalary, parsingSalaryErr := strconv.Atoi(salary)
	if parsingSalaryErr != nil {
		return errors.New("invalid salary")
	}
	id, err := c.svc.Create(ctx, vacancies.CreateVacancyDto{
		Title:  title,
		Salary: parsedSalary,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Create vacancy with id %s \n", id)

	return nil
}

func (c VacanciesController) GetById(ctx context.Context, id string) error {
	result, err := c.svc.GetById(ctx, id)
	if err != nil {
		return err
	}

	parsedResult, _ := json.Marshal(result)

	fmt.Printf("Vacancy :  [%s] \n", string(parsedResult))
	return nil
}

func (c VacanciesController) GetAll(ctx context.Context) error {
	result, err := c.svc.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, res := range result {
		parsedResult, _ := json.Marshal(res)
		fmt.Printf("Vacancy :  [%s] \n", string(parsedResult))
	}

	return nil
}

func (c VacanciesController) Update(ctx context.Context, id string, title string, salary string) error {
	parsedSalary, parsingSalaryErr := strconv.Atoi(salary)
	if parsingSalaryErr != nil {
		return errors.New("invalid salary")
	}
	err := c.svc.Update(ctx, id, vacancies.UpdateVacancyDto{
		Title:  title,
		Salary: parsedSalary,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Update vacancy with id %s \n", id)

	return nil
}

func (c VacanciesController) Delete(ctx context.Context, id string) error {
	err := c.svc.Delete(ctx, id)
	if err != nil {
		return err
	}

	fmt.Printf("Delete vacancy with id %s \n", id)

	return nil
}
