default: pg

pg:
	export POSTGRESQL_URL=postgres://mshev:123qwe@localhost:5432/calendar?sslmode=disable
	echo $$POSTGRESQL_URL
.PHONY: pg
