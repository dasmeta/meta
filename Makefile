export DOCKER_BUILDKIT=1
export VERSION=0.0.1
export YAML_DIR=.

CURRENT_DIR := $(shell pwd)

# general commands
version:
	@echo ${VERSION}

# commands to run on the host
docker-build:
	docker build --platform linux/amd64 -t dasmeta/meta:latest .

docker-publish: docker-build
	docker tag dasmeta/meta:latest dasmeta/meta:${VERSION}
	docker push dasmeta/meta:${VERSION}
	docker push dasmeta/meta:latest

run:
	docker run -it \
		-v $(CURRENT_DIR)/tests/terraform/metacloud.example.yaml:/metacloud.yaml \
		dasmeta/meta:latest

apply-docker:
	docker run -it \
		-v $(CURRENT_DIR):/workspace
		-e GIT_TOKEN=$(GIT_TOKEN)
		-e TERRAFORM_CLOUD_TOKEN=$(TERRAFORM_CLOUD_TOKEN)
		-e AWS_ACCESS_KEY_ID=$(AWS_ACCESS_KEY_ID)
		-e AWS_SECRET_ACCESS_KEY=$(AWS_SECRET_ACCESS_KEY)
		-e AWS_DEFAULT_REGION=$(AWS_DEFAULT_REGION)
		dasmeta/meta:latest create

debug-gitlab:
	docker run -it \
		-v $(CURRENT_DIR)/tests/terraform:/build \
		-e ROOT_DIR=/build \
		-e GIT_TOKEN="54567898656" \
		-e TERRAFORM_CLOUD_TOKEN="54567898656" \
		-e AWS_ACCESS_KEY_ID="54567898656" \
		-e AWS_SECRET_ACCESS_KEY="54567898656" \
		-e AWS_DEFAULT_REGION="54567898656" \
		dasmeta/meta:latest create

debug-docker:
	docker run -it --entrypoint /bin/sh \
		-v $(CURRENT_DIR):/workspace \
		dasmeta/meta:latest

# comands to run inside docker container
init:
	rm -rf _tfc && mkdir _tfc
	cd _tfc && \
	terraform init \
		-backend-config=/build/config.tfbackend \
		-from-module="github.com/dasmeta/terraform-tfe-cloud.git?ref=DMVP-2598-simplify-tfe-use"

apply:
	cd _tfc && \
	cp /build/terraform.tfvars . && \
	terraform apply -auto-approve

debug:
	cd _tfc && \
	cp ../terraform.tfvars . && \
	terraform console

destroy:
	cd _tfc && \
	cp ../terraform.tfvars . && \
	terraform destroy

generate-config-files:
	/scripts/yaml-to-terraform-configs.sh $(ROOT_DIR)/metacloud.yaml
	cp /config.tfbackend /terraform.tfvars /build

adjust-env-vars:
	$(eval export TF_TOKEN_app_terraform_io=$(TERRAFORM_CLOUD_TOKEN))
	$(eval export TFE_TOKEN=$(TERRAFORM_CLOUD_TOKEN))
	$(eval export TF_VAR_token=$(TERRAFORM_CLOUD_TOKEN))

create: generate-config-files adjust-env-vars init apply

# experemental commands
# @todo set to apply from local (remote does not detect yaml files)
create-managing-workspace:
	curl \
		--header "Authorization: Bearer ${TFE_TOKEN}" \
		--header "Content-Type: application/vnd.api+json" \
		--request POST \
		--data "{\"data\": {\"attributes\": {\"name\": \"${TFE_MANAGING_WORKSPACE}\"},\"type\": \"workspaces\"}}" \
		https://app.terraform.io/api/v2/organizations/${TFE_ORG}/workspaces
