TAG := $$(git log -1 --pretty=%h)
DOCKER_USER := mfigurski80

build:
	@docker build -t ${DOCKER_USER}/tweet-api:${TAG} .

push:
	@docker push ${DOCKER_USER}/tweet-api
