.PHONY: cert

migrate:
	migrate -database "postgresql://postgres:postgres@localhost:5432/foods?sslmode=disable" -path migrations $(action)

migrate-force:
	migrate -database "postgresql://postgres:postgres@localhost:5432/foods?sslmode=disable" -path migrations force $(number)

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)