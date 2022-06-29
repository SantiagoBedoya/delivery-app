.PHONY: cert
cert:
	openssl genrsa -out cert/id_rsa 4096
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub
	
migrate:
	migrate -database "postgresql://postgres:postgres@localhost:5432/customers?sslmode=disable" -path migrations $(action)

migrate-force:
	migrate -database "postgresql://postgres:postgres@localhost:5432/customers?sslmode=disable" -path migrations force $(number)

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)