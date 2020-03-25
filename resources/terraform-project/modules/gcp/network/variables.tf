variable "project_id" {
  description = "Your GCP project ID"
  default     = ""
}

variable "name" {
  description = "Name to be used on all the resources as identifier"
  default     = ""
}

variable "subnets" {
  description = "A list of subnets inside the VPC"
  type        = "list"
  default     = []
}