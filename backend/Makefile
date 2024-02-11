DB_URL=mysql://root:secret@tcp(localhost:3306)/mydb

migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database $(DB_URL) -verbose up

migratedown:
	migrate -path db/migrations -database $(DB_URL) -verbose down