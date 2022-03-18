run_local:
	go run main.go

run_migrate:
	go run ./db/migrate/migrate.go

run_seed:
	psql -U postgres codedoct_gin_golang177 < db/seed.sql
