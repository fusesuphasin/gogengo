nuclear:
	- docker stop $$(docker ps -a -q)
	- docker kill $$(docker ps -q)
	- docker rm $$(docker ps -a -q)
	- docker rm $$(docker ps -a -q)
	- docker rmi $$(docker images -q)
	- docker system prune --all --force --volumes

dcup-dev:
	docker-compose up

dcup-prod:
	docker-compose -f ./docker-compose.prod.yml up

dc-down:
	docker-compose down

dc-clear:
	docker-compose down
	docker rmi -f $(docker images | grep promptpay)

run-test:
	docker-compose build test-runner
	docker-compose up test-runner
