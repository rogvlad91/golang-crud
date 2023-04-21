-- +goose Up
-- +goose StatementBegin

create table if not exists vacancies
(
    id              text        default gen_random_uuid() primary key,
    title           varchar(100),
    salary          integer,
    created_at      timestamp,
    updated_at      timestamp
);

create table if not exists candidates
(
    id              text         default gen_random_uuid() primary key,
    full_name       varchar(100),
    age             smallint     not null,
    wanted_job      varchar(100) not null,
    wanted_salary   integer,
    created_at      timestamp,
    updated_at      timestamp
);

create table if not exists candidate_vacancy
(
    id           text           default gen_random_uuid() primary key,
    candidate_id text           not null,
    vacancy_id   text           not null,
    created_at   timestamp,
    unique (candidate_id, vacancy_id)
);

create index wanted_job_idx on candidates (wanted_job);
create index wanted_salary_idx on candidates(wanted_salary);
create index vacancy_title_idx on vacancies (title);
create index salary_idx on vacancies (salary);

INSERT INTO vacancies (title, salary, created_at, updated_at)
SELECT
        'Vacancy ' || i,
        floor(random() * 100000),
        now() - (i || ' days')::interval,
        now() - (i || ' days')::interval
FROM generate_series(1, 10) s(i);

INSERT INTO candidates (full_name, age, wanted_job, wanted_salary, created_at, updated_at)
SELECT
        'Candidate ' || i,
        floor(random() * 50) + 18,
        'Job ' || floor(random() * 5) + 1,
        floor(random() * 100000),
        now() - (i || ' days')::interval,
        now() - (i || ' days')::interval
FROM generate_series(1, 10) s(i);

INSERT INTO candidate_vacancy (candidate_id, vacancy_id, created_at)
SELECT
    c.id,
    v.id,
    now() - (i || ' days')::interval
FROM candidates c
         CROSS JOIN vacancies v
         CROSS JOIN generate_series(1, 20) s(i)
WHERE random() < 0.5
ON CONFLICT (candidate_id, vacancy_id) DO NOTHING;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table candidate_vacancy, candidates, vacancies
-- +goose StatementEnd
