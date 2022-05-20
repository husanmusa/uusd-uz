export POSTGRESQL_URL='postgres://husanmusa:1234@localhost:5432/uusduz?sslmode=disable'
#export POSTGRESQL_URL='postgres://postgres:pass@149.154.65.200:5432:5432/ussduz?sslmode=disable'

migrate -database ${POSTGRESQL_URL} -path migrations down
#force
# migrate -database ${POSTGRESQL_URL} -path migrations force 1