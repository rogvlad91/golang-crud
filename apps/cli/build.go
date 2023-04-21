package cli

import (
	"context"
	"fmt"
	"homework-7/apps"
	candidateCommand "homework-7/apps/cli/commands/candidate"
	candidateVacancyCommand "homework-7/apps/cli/commands/candidate_vacancy"
	fmtCommand "homework-7/apps/cli/commands/fmt"
	"homework-7/apps/cli/commands/spell"
	vacancyCommand "homework-7/apps/cli/commands/vacancy"
	"homework-7/apps/cli/controllers/candidate"
	candidateVacancy "homework-7/apps/cli/controllers/candidate_vacancy"
	"homework-7/apps/cli/controllers/vacancy"
	"homework-7/internal/svc/candidate_vacancies"
	candidateVacancyMemcachedRepo "homework-7/internal/svc/candidate_vacancies/repo/memcached"
	candidateVacancyPgRepo "homework-7/internal/svc/candidate_vacancies/repo/pg"
	"homework-7/internal/svc/candidates"
	candidateMemcachedRepo "homework-7/internal/svc/candidates/repo/memcached"
	candidateRepo "homework-7/internal/svc/candidates/repo/pg"
	"homework-7/internal/svc/vacancies"
	vacancyMemcachedRepo "homework-7/internal/svc/vacancies/repo/memcached"
	vacancyRepo "homework-7/internal/svc/vacancies/repo/pg"
	"os"
)

func RunCli(ctx context.Context) {

	pgRepo := apps.BuildPg(ctx)

	memCachedRepo := apps.BuildMemcached()

	////init pgRepos
	candidatePgRepo := candidateRepo.NewCandidatePGRepo(pgRepo)
	vacancyPgRepo := vacancyRepo.NewVacancyPGRepo(pgRepo)
	candidateVacancyRepo := candidateVacancyPgRepo.NewCandidateVacancyPGRepo(pgRepo)

	/////init cache
	candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
	vacancyCacheRepo := vacancyMemcachedRepo.NewVacancyMemcachedRepo(memCachedRepo)
	candidateVacancyCacheRepo := candidateVacancyMemcachedRepo.NewCandidateVacancyMemcachedRepo(memCachedRepo)

	///initSvc
	candidateSvc := candidates.NewCandidateSvc(candidatePgRepo)
	vacancySvc := vacancies.NewVacancySvc(vacancyPgRepo)
	candidateVacancySvc := candidate_vacancies.NewCandidateVacanciesSvc(candidateVacancyRepo)

	/////init cacheSvc
	cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)
	cacheVacancySvc := vacancies.NewCacheVacancySvc(vacancyCacheRepo, vacancySvc)
	candidateVacancyCacheSvc := candidate_vacancies.NewCandidateVacanciesCacheSvc(candidateVacancyCacheRepo, candidateVacancySvc)

	///init controllers
	candidateController := candidate.NewCandidatesController(cacheCandidateSvc)
	vacancyController := vacancy.NewVacanciesController(cacheVacancySvc)
	candidateVacancyController := candidateVacancy.NewCandidateVacancyController(candidateVacancyCacheSvc, cacheCandidateSvc, cacheVacancySvc)

	createCandidateCommand := &candidateCommand.CreateCandidateCommand{CandidateController: candidateController}
	updateCandidateCommand := &candidateCommand.UpdateCandidateCommand{CandidateController: candidateController}
	deleteCandidateCommand := &candidateCommand.DeleteCandidateCommand{CandidateController: candidateController}
	getCandidateCommand := &candidateCommand.GetCandidateCommand{CandidateController: candidateController}
	getCandidatesCommand := &candidateCommand.GetCandidatesCommand{CandidateController: candidateController}

	createVacancyCommand := &vacancyCommand.CreateVacancyCommand{VacancyController: vacancyController}
	updateVacancyCommand := &vacancyCommand.UpdateVacancyCommand{VacancyController: vacancyController}
	deleteVacancyCommand := &vacancyCommand.DeleteVacancyCommand{VacancyController: vacancyController}
	getVacancyCommand := &vacancyCommand.GetVacancyCommand{VacancyController: vacancyController}
	getVacanciesCommand := &vacancyCommand.GetVacanciesCommand{VacancyController: vacancyController}

	createResponseCommand := &candidateVacancyCommand.CreateResponseCommand{VacancyCandidateController: candidateVacancyController}
	deleteResponseCommand := &candidateVacancyCommand.DeleteResponseCommand{VacancyCandidateController: candidateVacancyController}
	getCandidatesByVacancyCommand := &candidateVacancyCommand.GetCandidatesByVacancyCommand{VacancyCandidateController: candidateVacancyController}
	getVacanciesByCandidateCommand := &candidateVacancyCommand.GetVacanciesByCandidateCommand{VacancyCandidateController: candidateVacancyController}

	formatCommand := &fmtCommand.FmtCommand{}
	spellCommand := &spell.SpellCommand{}

	commands := map[string]Command{
		"create_candidate":           createCandidateCommand,
		"update_candidate":           updateCandidateCommand,
		"delete_candidate":           deleteCandidateCommand,
		"get_candidate":              getCandidateCommand,
		"get_candidates":             getCandidatesCommand,
		"create_vacancy":             createVacancyCommand,
		"update_vacancy":             updateVacancyCommand,
		"delete_vacancy":             deleteVacancyCommand,
		"get_vacancy":                getVacancyCommand,
		"get_vacancies":              getVacanciesCommand,
		"create_response":            createResponseCommand,
		"delete_response":            deleteResponseCommand,
		"get_candidates_by_vacancy":  getCandidatesByVacancyCommand,
		"get_vacancies_by_candidate": getVacanciesByCandidateCommand,
		"spell":                      spellCommand,
		"fmt":                        formatCommand,
	}

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: app [command]")
		os.Exit(1)
	}

	if args[1] == "help" {
		printHelp(commands)
		os.Exit(0)
	}

	cmd, ok := commands[args[1]]
	if !ok {
		printHelp(commands)
		os.Exit(1)
	}

	cmd.Validate(args)

	err := cmd.Execute(ctx, args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printHelp(commands map[string]Command) {
	fmt.Println("Usage: ")
	commandsArr := []string{"create_candidate", "update_candidate", "delete_candidate", "get_candidate", "get_candidates", "create_vacancy", "update_vacancy", "delete_vacancy", "get_vacancy", "get_vacancies", "create_response", "delete_response", "get_candidates_by_vacancy", "get_vacancies_by_candidate", "spell", "fmt"}
	for _, command := range commandsArr {
		commands[command].Usage()
	}
}
