MIGRATIONS_PATH = internal/infra/database/sql/migrations

DATABASE_CONNECTION = "mysql://root:root@tcp(localhost:3306)/orders"

createmigration:
	migrate create -ext=sql -dir=$(MIGRATIONS_PATH) -seq init

migrate:
	 migrate -path=$(MIGRATIONS_PATH) -database $(DATABASE_CONNECTION) up

migratedown:
	 migrate -path=$(MIGRATIONS_PATH) -database $(DATABASE_CONNECTION) down

.PHONY: migrate migratedown createmigration