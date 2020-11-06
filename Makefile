
run:
	 LOG_DIR=$(LOG_DIR) go run ./cmd/main.go -conf=./config/db.toml -httpconf=./config/http.toml

create:
	 go run ./cmd/mysql/main.go   --do create
drop:
	 go run ./cmd/mysql/main.go   --do drop
migrate:
	 go run ./cmd/mysql/main.go   --do migrate



