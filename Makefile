fmt:
	./shell/fmt.sh

gen:
	./shell/gen.sh
	./shell/fmt.sh

lint:
	./shell/lint.sh

build:
	docker compose build

up:
	docker compose up

ps:
	docker compose ps

dev:
	docker compose build
	docker compose up

test:
	go test -v ./...

minio:
	docker container run -d --name minio -p 9000:9000 -p 9001:9001 minio/minio server /data --console-address ":9001"
minio_old:
	docker container run -d --name minio -p 9000:9000 -p 9001:9001 minio/minio:RELEASE.2022-10-08T20-11-00Z server /data --console-address ":9001"

migration:
	set -a && source makefile.env && set +a&&\
	atlas schema apply \
	-u "mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@$${IP}:$${PORT}/$${MYSQL_DATABASE}" \
	--to file://db/atlas/schema.hcl

schema_output:
	set -a && source makefile.env && set +a&&\
	atlas schema inspect -u "mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@$${IP}:$${PORT}/$${MYSQL_DATABASE}" > db/atlas/schema.hcl