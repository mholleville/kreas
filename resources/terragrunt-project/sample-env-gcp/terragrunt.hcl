locals {
  # Automatically load account-level variables
  account_vars = read_terragrunt_config(find_in_parent_folders("account.hcl"))

  # Automatically load environment-level variables
  environment_vars = read_terragrunt_config(find_in_parent_folders("env.hcl"))

  # Extract the variables we need for easy access
  gcp_automation_file_path = local.account_vars.locals.gcp_automation_file_path
  gcp_region   = local.environment_vars.locals.gcp_region
  gcp_project = local.environment_vars.locals.gcp_project
  gcp_zone = local.environment_vars.locals.gcp_zone

}

remote_state {
  backend =  "gcs"
  config  = {
    bucket      = "<bucket_state_terragrunt>"
    prefix      = "${path_relative_to_include()}/terraform.tfstate"
  }
}
# Generate an GCP provider block
generate "provider" {
  path      = "provider.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<EOF
provider "google-beta" {
  credentials = file("${local.gcp_automation_file_path}")
  project     = "${local.gcp_project}"
  region      = "${local.gcp_region}"
  zone        = "${local.gcp_zone}"
  version     = "3.0.0"
}


provider "google" {
  credentials = file("${local.gcp_automation_file_path}")
  project     = "${local.gcp_project}"
  region      = "${local.gcp_region}"
  zone        = "${local.gcp_zone}"
  version     = "3.0.0"
}
EOF
}

# ---------------------------------------------------------------------------------------------------------------------
# GLOBAL PARAMETERS
# These variables apply to all configurations in this subfolder. These are automatically merged into the child
# `terragrunt.hcl` config via the include block.
# ---------------------------------------------------------------------------------------------------------------------

# Configure root level variables that all resources can inherit. This is especially helpful with multi-account configs
# where terraform_remote_state data sources are placed directly into the modules.
inputs = merge(
local.account_vars.locals,
local.environment_vars.locals,
)