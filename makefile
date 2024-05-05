.PHONY: run
run:
	docker-compose -f docker/compose.yml up -d

.PHONY: stop
down:
	docker-compose -f docker/compose.yml down

.PHONY: clean
rm:
	docker-compose -f docker/compose.yml down --rmi all --volumes --remove-orphans
