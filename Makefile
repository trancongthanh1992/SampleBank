# connect `bank-network` into postgres.
# docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:aA@123123@postgres14.4:5432/sample_bank_db?sslmode=disable" simplebank:latest

# postgres connect to `bank-network`.
# golang connect database from `postgres` via --network at line 2. 
# Golang `sample-bank` will be running on the same network with Postgres.

# pd master password -> "tcthanh1992"
# postgres can set isolation level within transaction.
# show transaction isolation level; # postgres
# https://www.postgresql.org/docs/current/transaction-iso.html
# https://dev.mysql.com/doc/refman/8.0/en/innodb-transaction-isolation-levels.html
postgres:
	docker run --name postgres14.4 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=aA@123123 -d postgres:14.4-alpine
createdb:
	docker exec -it postgres14.4 createdb --username=root --owner=root sample_bank_db
dropdb:
	docker exec -it postgres14.4 dropdb sample_bank_db
migrateup:
	migrate -path db/migration -database "postgresql://root:aA@123123@localhost:5432/sample_bank_db?sslmode=disable" -verbose up
migrateup:
	migrate -path db/migration -database "postgresql://root:aA@123123@localhost:5432/sample_bank_db?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:aA@123123@localhost:5432/sample_bank_db?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:aA@123123@localhost:5432/sample_bank_db?sslmode=disable" -verbose down 1
# To rememberance.
migratecreate:
	migrate create -ext sql -dir db/migration -seq add_users
sqlc:
	sqlc generate

# run all unit-test with `./...`
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/trancongthanh1992/samplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server