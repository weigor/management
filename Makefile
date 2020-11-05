
run:
	APP_ID=management go run ./cmd/main.go

create:
	APP_ID=mysql go run ./cmd/mysql/main.go   --do create
drop:
	APP_ID=mysql go run ./cmd/mysql/main.go   --do drop
migrate:
	APP_ID=mysql go run ./cmd/mysql/main.go   --do migrate



