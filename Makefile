start:
	docker-compose up -d

start_build:
	docker-compose up -d --build

stop:
	docker-compose down

start_prod:
	docker-compose -f docker-compose.prod.yaml -p buf-connect-prod up -d --build

stop_prod:
	docker-compose -f docker-compose.prod.yaml -p buf-connect-prod down