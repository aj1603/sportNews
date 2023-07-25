run:
	go run main.go

refresh_db:
	psql -U postgres -d postgres -f ./sql/refresh.sql

create_tables:
	psql -U postgres -d sport_news -f ./sql/create.sql
	
insert_data:
	psql -U postgres -d sport_news -f ./sql/insert.sql