locals {
  # Automatically load environment-level variables
  environment_vars = read_terragrunt_config(find_in_parent_folders("env.hcl"))
  env = local.environment_vars.locals.environment
  gcp_project = local.environment_vars.locals.gcp_project
}
# Terragrunt will copy the Terraform configurations specified by the source parameter, along with any files in the
# working directory, into a temporary folder, and execute your Terraform commands in that folder.
terraform {
  source = ""
}

# Include all settings from the root terragrunt.hcl file
include {
  path = find_in_parent_folders()
}

# These are the variables we have to pass in to use the module specified in the terragrunt configuration above
inputs = {
  bucket_name = "test-${local.gcp_project}-${local.env}-bucket"
  location = "EUROPE-WEST1"
  versioning_state = false
  lifecycle_action_type = "Delete"
  lifecycle_condition_age = 1
  retention_policy_period = 1
}