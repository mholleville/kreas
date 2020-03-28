provider "google" {
  project     = var.project_id
  version     = "3.0.0"
}

terraform {
  backend "gcs" {
    prefix  = "terraform/state"
    # ...
  }
}

module "network" {
  source = "..\/..\/modules\/gcp\/network"
  project_id = var.project_id
  name = var.name
  subnets = var.subnets
}

module "managed-instance-group" {
  source  = "GoogleCloudPlatform/managed-instance-group/google"
  version = "1.1.15"
}