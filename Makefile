
run:
	 LOG_DIR=$(LOG_DIR) go run ./cmd/main.go -conf=./config/db.toml -httpconf=./config/http.toml

create:
	 go run ./cmd/mysql/main.go   --do create --db ./config/db.toml
drop:
	 go run ./cmd/mysql/main.go   --do drop  --db ./config/db.toml
migrate:
	 go run ./cmd/mysql/main.go   --do migrate --db ./config/db.toml



