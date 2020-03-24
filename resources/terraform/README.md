# Terraform project by Kreas

The following elements constitute the project :

```
.
├── modules/                # A directory containing any terraform module for Amazon and GCP provider.
├── sample/                 # A directory containing the infrastructure definitions scripts.
└── README.md               # OPTIONAL: A human-readable README file
```

### Usage for GCP

```shell script
$ cd sample/gcp
$ terraform init
$ terraform plan
```

NOTE : It is recommended to always execute a `terraform plan` command before `terraform apply`

### Usefull links

- [terraform](https://www.terraform.io/)
- [cli](https://www.terraform.io/docs/cli-index.html)
- [provider](https://www.terraform.io/docs/providers/index.html)
- [guides](https://www.terraform.io/guides/index.html)
- [tfswitch](https://warrensbox.github.io/terraform-switcher/)

