package pg

const (
	CreateQuery                    = "insert into candidate_vacancy(id, candidate_id, vacancy_id, created_at) values ($1, $2, $3, $4)"
	DeleteQuery                    = "delete from candidate_vacancy where vacancy_id = $1 and candidate_id = $2"
	GetCandidatesByVacancyIdQuery  = "select candidate_id from candidate_vacancy where vacancy_id = $1"
	GetVacanciesByCandidateIdQuery = "select vacancy_id from candidate_vacancy where candidate_id = $1"
)
