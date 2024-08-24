.PHONY: postgres createdb dropdb migrateup migratedown sqlc test

# Định nghĩa target cho PostgreSQL
postgres:
	docker run --name pg12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

# Định nghĩa target để tạo database
createdb:
	docker exec -it pg12 createdb --username=root --owner=root simple_bank

# Định nghĩa target để xóa database
dropdb:
	docker exec -it pg12 dropdb simple_bank

# Định nghĩa target để thực hiện migration lên
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

# Định nghĩa target để thực hiện migration xuống
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

# Định nghĩa target để generate code bằng sqlc
sqlc:
	sqlc generate

# Định nghĩa target để chạy test
test:
	go test -v -cover ./...
