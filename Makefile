client:
	cd client/instafake && npm start

services-up:
	docker-compose up -d

services-down:
	docker-compose down --remove-orphans
	