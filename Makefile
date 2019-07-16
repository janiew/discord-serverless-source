.PHONY: docker \
	docker-build docker-push \
	docker-run kubernetes-run

DOCKER_TAG_VERSION ?= staging-latest
DOCKER_TAG_C ?= pratikmahajan/twitter-stream-source:${DOCKER_TAG_VERSION}


NAMESPACE ?= svrless
POD ?= staging-bot-api
PROD_POD ?= prod-bot-api



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
	docker run -it --rm --net host ${DOCKER_TAG_C} /bin/sh

#Staging the app
staging: docker staging-rollout

# Deploy code to Production:
production:
	./deploy/deploy.sh -n ${NAMESPACE} -p ${PROD_POD} -t prod

#deploy code to staging:
staging-rollout:
	./deploy/deploy.sh -n ${NAMESPACE} -p ${POD} -t staging