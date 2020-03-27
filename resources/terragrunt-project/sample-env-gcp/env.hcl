# Set common variables for the environment. This is automatically pulled in in the root terragrunt.hcl configuration to
# feed forward to the child modules.
locals {
  environment = "sample-env-gcp"
  gcp_region = "europe-west1"
  gcp_project = ""
  gcp_zone = "europe-west1-d"
}