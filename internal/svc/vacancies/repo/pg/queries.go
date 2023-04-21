package pg

const (
	CreateQuery  = "insert into vacancies(id,title,salary,created_at,updated_at) values ($1,$2,$3,$4,$5)"
	GetByIdQuery = "select id, title, salary, created_at, updated_at from vacancies where id=$1"
	GetAllQuery  = "select id, title, salary, created_at, updated_at from vacancies"
	UpdateQuery  = "update vacancies set title = $1, salary = $2, updated_at = $3 where id = $4"
	DeleteQuery  = "delete from vacancies where id = $1"
)
