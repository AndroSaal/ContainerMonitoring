local:
	docker-compose -f docker-compose.yml up -d

down:
	docker compose down -v

stop:
	docker-compose down
