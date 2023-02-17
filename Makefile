run_local:
	go run main.go

run_migrate:
	go run ./db/migrate/migrate.go

run_seed:
	psql -U postgres codedoct_gin_golang177 < db/seed.sql

run_test:
	go test  -coverprofile=coverage.out -v -coverpkg=./... blog-gin_golang_v177/test/root

run_docker:
	docker stop basecodeapiserver || true && docker rm basecodeapiserver || true
	docker build --tag basecode-api:dev .
	docker run --name basecodeapiserver -d -p 4001:4001 basecode-api:dev
