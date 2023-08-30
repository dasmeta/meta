#!/bin/sh

# Check if argument provided
if [ -z "$1" ]; then
    echo "Usage: $0 <path/to/your/input.yaml>"
    exit 1
fi

# Input YAML file
input_file="$1"

# Output files
backend_config_file="config.tfbackend"
variables_file="terraform.tfvars"

# Extract values, prioritizing environment variables
terraform_cloud_org=${TERRAFORM_CLOUD_ORG:-$(yq e '.terraform_cloud_org' $input_file)}
terraform_cloud_workspace=${TERRAFORM_CLOUD_WORKSPACE:-$(yq e '.terraform_cloud_workspace' $input_file)}
terraform_cloud_token=${TERRAFORM_CLOUD_TOKEN:-$(yq e '.terraform_cloud_token' $input_file)}
git_provider=${GIT_SERVICE_PROVIDER:-$(yq e '.git_provider' $input_file)}
git_org=${GIT_ORG:-$(yq e '.git_org' $input_file)}
git_repo=${GIT_REPO:-$(yq e '.git_repo' $input_file)}
git_token=${GIT_TOKEN:-$(yq e '.git_token' $input_file)}
aws_access_key_id=${AWS_ACCESS_KEY_ID:-$(yq e '.aws_access_key_id' $input_file)}
aws_secret_access_key=${AWS_SECRET_ACCESS_KEY:-$(yq e '.aws_secret_access_key' $input_file)}
aws_default_region=${AWS_DEFAULT_REGION:-$(yq e '.aws_default_region' $input_file)}

# Write to backend.tf
cat > $backend_config_file <<EOF
organization = "$terraform_cloud_org"
workspaces { name = "$terraform_cloud_workspace" }
hostname     = "app.terraform.io"
EOF

# Write to variables.tfvars
cat > $variables_file <<EOF
org = "$terraform_cloud_org"
token = "$terraform_cloud_token"

git_provider = "$git_provider"
git_org = "$git_org"
git_repo = "$git_repo"
git_token = "$git_token"

aws_access_key_id = "$aws_access_key_id"
aws_secret_access_key = "$aws_secret_access_key"
aws_default_region = "$aws_default_region"
EOF

echo "Files created: $backend_config_file, $variables_file"
