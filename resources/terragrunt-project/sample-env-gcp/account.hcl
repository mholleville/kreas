# Set common variables for the environment. This is automatically pulled in in the root terragrunt.hcl configuration to
# feed forward to the child modules.
locals {
  gcp_automation_file_path = "${get_terragrunt_dir()}/<path-to-JSON-SA-file>.json"
}