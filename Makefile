PROJECT = gateway-service
DOCKER_REPO = shourov
VERSION = 0.0.1
# will be used to upload to docker hub
docker:
	docker build -t $(DOCKER_REPO)/$(PROJECT):$(VERSION) .
	docker push $(DOCKER_REPO)/$(PROJECT):$(VERSION)