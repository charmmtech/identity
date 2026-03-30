docker-build:
	docker build --tag localhost:3000/charmm/identity:latest .

docker-push:
	docker push localhost:3000/charmm/identity:latest

docker-deploy:
	@make docker-build
	@make docker-push
