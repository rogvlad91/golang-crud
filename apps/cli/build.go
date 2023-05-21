package cli

import (
	"context"
	"fmt"
	"golang-crud/apps"
	candidateCommand "golang-crud/apps/cli/commands/candidate"
	candidateVacancyCommand "golang-crud/apps/cli/commands/candidate_vacancy"
	fmtCommand "golang-crud/apps/cli/commands/fmt"
	"golang-crud/apps/cli/commands/spell"
	vacancyCommand "golang-crud/apps/cli/commands/vacancy"
	"golang-crud/apps/cli/controllers/candidate"
	candidateVacancy "golang-crud/apps/cli/controllers/candidate_vacancy"
	"golang-crud/apps/cli/controllers/vacancy"
	"golang-crud/internal/svc/candidate_vacancies"
	candidateVacancyMemcachedRepo "golang-crud/internal/svc/candidate_vacancies/repo/memcached"
	candidateVacancyPgRepo "golang-crud/internal/svc/candidate_vacancies/repo/pg"
	"golang-crud/internal/svc/candidates"
	candidateMemcachedRepo "golang-crud/internal/svc/candidates/repo/memcached"
	candidateRepo "golang-crud/internal/svc/candidates/repo/pg"
	"golang-crud/internal/svc/vacancies"
	vacancyMemcachedRepo "golang-crud/internal/svc/vacancies/repo/memcached"
	vacancyRepo "golang-crud/internal/svc/vacancies/repo/pg"
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
