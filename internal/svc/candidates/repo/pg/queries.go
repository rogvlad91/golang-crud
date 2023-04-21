package pg

const (
	CreateQuery  = "insert into candidates(id,full_name,age,wanted_job,wanted_salary,created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7)"
	GetByIdQuery = "select id, full_name,age,wanted_job,wanted_salary, created_at, updated_at from candidates where id=$1"
	GetAllQuery  = "select id, full_name,age,wanted_job,wanted_salary, created_at, updated_at from candidates"
	UpdateQuery  = "update candidates set wanted_job = $1, wanted_salary = $2, updated_at = $3 where id = $4"
	DeleteQuery  = "delete from candidates where id = $1"
)
