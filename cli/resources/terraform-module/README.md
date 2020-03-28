# Terraform network Module

It creates a network.

It supports creating:

- A Network

## Usage

```hcl-terraform
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = "v1.44.0"
  name = "kreas-sample"
  cidr = "0.0.0.0/0"
  public_subnets = []
}
```

## Inputs

| Name | Description | Type | Default | Required |
| ------ | ------ | ------ | ------ | ------ |
| name | The name of the network being created | string | `""` | yes |
| cidr | The CIDR block for the VPC. Default value is a valid CIDR, but not acceptable by AWS and should be overriden | string | - | yes |
| azs | A list of availability zones in the region | list | `[]` | no |
| public_subnets | A list of public subnets inside the VPC | list | `[]` | no |

## Outputs

| Name | Description |
| ------ | ------ |
| vpc_id | The ID of the VPC |