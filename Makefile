# export AWS_PROFILE=
# export AWS_CONFIG_FILE=~/.aws/config
export AWS_REGION=eu-central-1
export AWS_ACCOUNT_ID=

APP=
ENV=staging
ROLE_ARN=
SESSION_NAME=default

CLUSTER=
eks-select:
	aws --region eu-central-1 eks update-kubeconfig --name ${CLUSTER}

.EXPORT_ALL_VARIABLES:

debug-aws-session:
	aws sts get-caller-identity

assume-role:
	aws sts assume-role --role-arn ${ROLE_ARN} --role-session-name ${SESSION_NAME} | jq -r '.Credentials | "export AWS_ACCESS_KEY_ID=\(.AccessKeyId) export AWS_SECRET_ACCESS_KEY=\(.SecretAccessKey) export AWS_SESSION_TOKEN=\(.SessionToken)"'
