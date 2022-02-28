mysql:
	docker-compose up -d

migrate-up:
	migrate -path db/migration -database "mysql://admin:Qwe12345@tcp(127.0.0.1:3306)/simple_bank" -verbose up

migrate-down:
	migrate -path db/migration -database "mysql://admin:Qwe12345@localhost:3306/simple_bank?sslmode=disable" -verbose down 