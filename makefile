.PHONY: clean
rm:
	docker-compose -f docker/compose.yml down --rmi all --volumes --remove-orphans
