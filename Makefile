READINESS_PROBE_IMAGE_NAME ?= localhost/readiness-check-probe

CONTAINER_ENGINE := $(shell command -v podman 2> /dev/null | echo docker)
CONTAINER_BUILD_EXTRA_FLAGS =
CONTAINER_PUSH_EXTRA_FLAGS =

.PHONY: image push run

image:
	$(CONTAINER_ENGINE) build -t $(READINESS_PROBE_IMAGE_NAME)  $(CONTAINER_BUILD_EXTRA_FLAGS) -f Containerfile .

push:  
	$(CONTAINER_ENGINE) push $(CONTAINER_PUSH_EXTRA_FLAGS) $(READINESS_PROBE_IMAGE_NAME)

run: image
	$(CONTAINER_ENGINE) run --rm -it -p 8000:8000 --name readiness-probe $(READINESS_PROBE_IMAGE_NAME)
