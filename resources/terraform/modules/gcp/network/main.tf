module "network" {
  source  = "terraform-google-modules/network/google"
  version = "2.1.1"
  project_id = var.project_id
  network_name = var.name
  subnets = var.subnets
}