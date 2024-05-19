build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.auth . 

run:
	docker service create --name auth-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6015:6015 github.com/turistikrota/service.auth:latest

remove:
	docker service rm auth-api-turistikrota-com

stop:
	docker service scale auth-api-turistikrota-com=0

start:
	docker service scale auth-api-turistikrota-com=1

restart: remove build run
