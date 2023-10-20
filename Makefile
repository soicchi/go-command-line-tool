file_detect:
	docker compose run --rm app go run ./cmd/${entry}/main.go -dir=${dir} -keyword=${keyword}

go_vet:
	docker compose run --rm app go vet ./...

go_fmt:
	docker compose run --rm app go fmt ./...

go_tidy:
	docker compose run --rm app go mod tidy