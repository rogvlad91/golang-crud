package http

import (
	"context"
	"homework-7/apps"
	candidateVacanciesHttp "homework-7/apps/http/candidate_vacancies"
	candidatesHttp "homework-7/apps/http/candidates"
	vacanciesHttp "homework-7/apps/http/vacancies"
	"homework-7/internal/svc/candidate_vacancies"
	candidateVacancyMemcachedRepo "homework-7/internal/svc/candidate_vacancies/repo/memcached"
	candidateVacancyPgRepo "homework-7/internal/svc/candidate_vacancies/repo/pg"
	"homework-7/internal/svc/candidates"
	candidateMemcachedRepo "homework-7/internal/svc/candidates/repo/memcached"
	candidateRepo "homework-7/internal/svc/candidates/repo/pg"
	"homework-7/internal/svc/vacancies"
	vacancyMemcachedRepo "homework-7/internal/svc/vacancies/repo/memcached"
	vacancyRepo "homework-7/internal/svc/vacancies/repo/pg"
	"net/http"
)

func Run(ctx context.Context) {

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

	///initControllers
	candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)
	vacancyController := vacanciesHttp.NewVacancyController(cacheVacancySvc)
	candidateVacancyController := candidateVacanciesHttp.NewCandidateVacanciesController(candidateVacancyCacheSvc)

	///initRouters
	candidateRouter := candidatesHttp.GetRouter(candidatesController)
	vacancyRouter := vacanciesHttp.GetRouter(vacancyController)
	candidateVacancyRouter := candidateVacanciesHttp.GetRouter(candidateVacancyController)

	mux := http.NewServeMux()

	mux.Handle("/candidates", candidateRouter)
	mux.Handle("/vacancies", vacancyRouter)
	mux.Handle("/candidate-vacancies", candidateVacancyRouter)

	http.ListenAndServe(":8080", mux)
}
