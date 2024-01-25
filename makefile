diff:
	atlas migrate diff $(NAME) \
		--dir "file://database/migrations" \
		--to "file://database/schema/schema.sql" \
		--dev-url "sqlite://dev?mode=memory"

migrate:
	atlas migrate apply \
		--dir "file://database/migrations" \
		--url "sqlite://sqlite.db"

gen: 
	sqlc generate && templ generate

seed:
	go run ./cmd/seed/seed.go

ssr:
	go run ./cmd/serve/ssr/main.go

ssg:
	go run ./cmd/serve/ssg/main.go
