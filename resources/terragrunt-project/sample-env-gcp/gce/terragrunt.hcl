locals {
  # Automatically load environment-level variables
  environment_vars = read_terragrunt_config(find_in_parent_folders("env.hcl"))

  # Extract out common variables for reuse
  gcp_region   = local.environment_vars.locals.gcp_region
  gcp_project = local.environment_vars.locals.gcp_project
  gcp_zone = local.environment_vars.locals.gcp_zone
  env = local.environment_vars.locals.environment
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

dependency "vpc" {
  config_path = "../vpc"
}
# These are the variables we have to pass in to use the module specified in the terragrunt configuration above
inputs = {
  gce_zone                        = "${local.gcp_zone}"
  gce_region                      = "${local.gcp_region}"
  gce_name                        = "${local.gcp_project}-${local.env}-bastion"
  gcp_project                     = "${local.gcp_project}"

  vpc_name                        = dependency.vpc.outputs.vpc_name
  vpc_self_link                   = dependency.vpc.outputs.vpc_self_link
  subnetwork_range                = "10.185.0.0/24"
}