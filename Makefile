
run:
	 LOG_DIR=$(LOG_DIR) go run ./cmd/main.go -conf=./config/db.toml -httpconf=./config/http.toml

create:
	APP_ID=mysql go run ./cmd/mysql/main.go   --do create
drop:
	APP_ID=mysql go run ./cmd/mysql/main.go   --do drop
migrate:
	APP_ID=mysql go run ./cmd/mysql/main.go   --do migrate



