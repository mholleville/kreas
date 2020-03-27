# Terragrunt project by Kreas

The following elements constitute the project :

```
├── README.md
└── sample-env-gcp          # Directory to group all the resources to be created in the env
    ├── account.hcl         # Variable file to fetch account infos
    ├── buckets             # Resource folder
    ├── env.hcl             # Variable file contains some specific env variables
    ├── gce                 # Resource folder
    ├── terragrunt.hcl      # Root Terragrunt file
    └── vpc                 # Resource folder
```

## Usage for GCP

```shell script
$ cd terragrunt-project/sample-env-gcp
$ terragrunt apply-all 
```

NOTE : It is recommended to always execute a `terragrunt plan-all` command before `terragrunt apply-all`

## Useful links

- [terragrunt docs](https://terragrunt.gruntwork.io/docs/)


