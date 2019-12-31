# terraform-provider-appstream
Terraform appstream provider

# Provider usage

```
$ go build -o terraform-provider-appstream
$ terraform init
$ terraform plan
$ terraform apply
```

#Development notes
This project references some underlying functions from [terraform-provider-aws](https://github.com/terraform-providers/terraform-provider-aws), which subsequently uses `Go Modules`

#Authentication
AWS Authentication is supported via static credentials or via shared credentials file

* Set static credentials
```hcl-terraform
provider "appstream" {
    access_key              = "${var.access_key}",
    secret_key              = "${var.secret_key}",
    region                  = "${var.region}",
    token                   = "${var.token}",
}
```

* Set the profile to use via shared credentials file
```hcl-terraform
provider "appstream" {
  region  = var.aws_region
  profile = var.aws_profile[terraform.workspace]
  version = "~> 2.0"
}
```