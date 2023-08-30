FROM alpine

RUN apk add --no-cache make yq git
# RUN apk-install terraform aws-cli gcloud make

# install terraform
ARG TERRAFORM_VERSION=1.5.6
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip
RUN unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip
RUN mv terraform /usr/bin/terraform

COPY ./scripts /scripts

COPY Makefile /

ENTRYPOINT [ "make" ]
