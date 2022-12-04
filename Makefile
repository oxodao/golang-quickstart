run:
	go run .

create-db:
	mkdir -p 0_DATA
	rm -rf 0_DATA/app.db
	sqlite3 0_DATA/app.db < sql/init.sql

reset-db:
	mkdir -p 0_DATA
	rm -rf 0_DATA/app.db
	sqlite3 0_DATA/app.db < sql/init.sql
	sqlite3 0_DATA/app.db < sql/fixtures.sql
