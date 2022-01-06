app_start:
	docker-compose up && exit $(docker wait proxy)