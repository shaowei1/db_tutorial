db: db.c
	gcc db.c -o db

run: db
	./db mydb.db

clean:
	rm -f db *.db

test: db
	bundle exec rspec

build_test:
	gcc mdb.c -o mdb
	cd gov && go test

format: *.c
	clang-format -style=Google -i *.c
