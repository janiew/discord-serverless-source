.PHONY: docker \
	docker-build docker-push \
	docker-run kubernetes-run

DOCKER_TAG_VERSION ?= latest
DOCKER_TAG_C ?= pratikmahajan/twitter-stream-source:${DOCKER_TAG_VERSION}

# Build and Push to docker hub
docker: docker-build docker-push

# build Docker image
docker-build:
	docker build -t ${DOCKER_TAG_C} .

# Push the docker image to docker hub
docker-push:
	docker push ${DOCKER_TAG_C}

# Runs the docker image on local machine
docker-run:
	docker run -it --rm --net host ${DOCKER_TAG_C}

