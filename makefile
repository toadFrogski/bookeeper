SCHEMA := ${SCHEMA}
HOSTNAME := ${HOSTNAME}
DB_DSN := ${DB_DSN}

release:	build compile_api build_frontend

build:
	docker compose up -d --build
	docker compose exec nginx chown nginx:nginx /var/run/backend/www.sock

build_frontend:
	docker run -v ${PWD}/frontend:/frontend -w /frontend bookeeper/frontend build

compile_api:
	docker run -v ${PWD}/backend:/app -v ${PWD}/docs:/docs swag/autodoc init -g ./cmd/web/main.go -o /docs -d ./ --parseDependency -ot yaml --parseGoList=false
	docker run -w /work -v ${PWD}:/work openapitools/openapi-generator-cli:v6.6.0 generate -i /work/docs/swagger.yaml -o /work/frontend/src/services/api -g typescript-axios --server-variables schema=$(SCHEMA),hostname=$(HOSTNAME)

migrate:
	docker run -w /work -v ${PWD}/backend:/work -v /var/run/docker.sock:/var/run/docker.sock --network host gomigtools/atlas schema apply -u ${DB_DSN} --env "gorm" --auto-approve
