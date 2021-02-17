.PHONY: all build-binary build-image push clean install uninstall

all: build-binary build-image push


TAG ?= 1.0.1
REGISTRY ?= sakibalamin
APP_NAME ?= apiserver
RELEASE_NAME ?= apiserver


build-binary:
	@echo Building Go binary file
	go build -o ${APP_NAME} .


build-image: build-binary
	@echo Building the API Server Project ...
	docker build -t ${REGISTRY}/${APP_NAME}:${TAG} .


push: build-image
	@echo Pushing the Image into ${REGISTRY}
	docker push ${REGISTRY}/${APP_NAME}:${TAG}


clean:
	@echo Cleaning up...
	rm -rf ${APP_NAME}

