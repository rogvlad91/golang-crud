package candidate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang-crud/internal/svc/candidates"
	"strconv"
)

func (c CandidatesController) Create(ctx context.Context, fullName string, age string, wantedJob string, wantedSalary string) error {
	parsedAge, parsingAgeErr := strconv.Atoi(age)
	if parsingAgeErr != nil {
		return errors.New("invalid age")
	}

	parsedSalary, parsingSalaryErr := strconv.Atoi(wantedSalary)
	if parsingSalaryErr != nil {
		return errors.New("invalid salary")
	}

	id, err := c.svc.Create(ctx, candidates.CreateCandidateDTO{
		FullName:     fullName,
		Age:          parsedAge,
		WantedJob:    wantedJob,
		WantedSalary: parsedSalary,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Create candidate with id %s \n", id)

	return nil
}

func (c CandidatesController) GetById(ctx context.Context, id string) error {
	result, err := c.svc.GetById(ctx, id)
	if err != nil {
		return err
	}

	parsedResult, _ := json.Marshal(result)

	fmt.Printf("Candidate :  [%s] \n", string(parsedResult))
	return nil
}

func (c CandidatesController) GetAll(ctx context.Context) error {
	result, err := c.svc.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, res := range result {
		parsedResult, _ := json.Marshal(res)
		fmt.Printf("Candidate :  [%s] \n", string(parsedResult))
	}

	return nil
}

func (c CandidatesController) Update(ctx context.Context, id string, wantedJob string, wantedSalary string) error {
	parsedSalary, parsingSalaryErr := strconv.Atoi(wantedSalary)
	if parsingSalaryErr != nil {
		return errors.New("invalid salary")
	}

	err := c.svc.Update(ctx, id, candidates.UpdateCandidateDto{
		WantedJob:    wantedJob,
		WantedSalary: parsedSalary,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Update candidate with id %s \n", id)

	return nil
}

func (c CandidatesController) Delete(ctx context.Context, id string) error {
	err := c.svc.Delete(ctx, id)
	if err != nil {
		return err
	}

	fmt.Printf("Delete candidate with id %s \n", id)

	return nil
}
