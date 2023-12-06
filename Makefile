createDb:
	createdb --username=postgres --owner=postgres go_finance

postgres:
	docker run --name postgres -p 5432