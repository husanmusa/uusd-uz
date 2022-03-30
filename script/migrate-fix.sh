export POSTGRESQL_URL='postgres://husanmusa:1234@localhost:5432/uusduz?sslmode=disable'

migrate -database ${POSTGRESQL_URL} -path migrations force 1