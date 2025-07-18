create_migrations :
				migrate create -ext sql -dir db/migrations -seq $(name)
			
migrate_up :
			migrate -database 'postgres://postgres:mysecret123@localhost:5432/kairodb?sslmode=disable' -path  db/migrations up

migrate_down :
			migrate -database 'postgres://postgres:mysecret123@localhost:5432/kairodb?sslmode=disable' -path  db/migrations down

postgres_up :
			# Create postgres server 			
			docker compose up -d

postgres_down :
			# Delete postgres server
			docker compose down

db_up :
		docker exec -it kairo_postgres createdb --username=postgres --owner=postgres  kairodb
db_down :
		docker exec -it kairo_postgres dropdb --username=postgres   kairodb

sql_gen :
		sqlc generate