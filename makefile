build:
	docker compose up -d --build
	docker compose exec nginx chown nginx:nginx /var/run/backend/www.sock

autodoc:
	docker compose run autodoc