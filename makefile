db_up:
	# Docker compose start running in a detached mode
	docker-compose up -d

db_down:
	# postgres down - delete postgres server
	docker-compose down
