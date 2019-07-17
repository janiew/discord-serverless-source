# Twitter-Knative-Serverless-App-Source

### Create Knative events using twitter 

Twitter event source implementation for Knative Eventing using `ContainerSource`.

## Setup
### Twitter API keys 
Add your Twitter API keys in `secrets.yaml` file in `deploy/templates`.

Check `secrets.yaml.example` file in `deploy/templates` for reference

_All the secrets should be in `base64` format in the `secrets.yaml`_ \
_P.S. your secrets file should always be gitignored and dockerignored to avoid leaking secrets_

### Query
Add your Twitter search query as `QUERY` parameter in `configmap.yaml` \
`QUERY` is taken as an environment variable in the application.\
By default `QUERY` is set to `google` in the `configmap`


## RUN

### Build the docker image 
`make docker-build`

### Run the docker image on local machine
`make docker-run` \
_The default entrypoint while running the docker image on local machine is `/bin/sh`_


### Upload the docker image to dockerhub
`make docker-push`

### Build and Upload the docker image to repository 
`make docker`

### Deploy the image to kubernetes/knative cluster
*Make sure you have knative installed on the cluster* 

#### Deploy the image to `Staging`
`make staging` \
*It creates a docker image from your current working directory and uploads it to docker hub before depolying the image* \

**If you are looking to deploy to staging without building and pushing docker image, use** \
`make staging-rollout`

#### Deploy the image to `Production`
`make production` \
*Takes the latest `production` image from docker hub to deploy* \
Make sure you have `github webhooks` configured on `dockerhub`



## Misc.
* _To change the default namespace and docker repository see `Makefile`_
* In order to run the application properly, make sure to set `knative-eventing-injection=enabled` as a `label` on your `namespace`
* `--sink` is taken as a commandline argument in the GoLang app