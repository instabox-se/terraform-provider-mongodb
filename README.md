# terraform-provider-mongodb

## Installation
Download source and build with `go build -o terraform-provider-mongodb`. Move resulting binary to `~/.terraform.d/plugins`.

The file name must be exactly `terraform-provider-mongodb` for this to work.

## Example
See `example.tf` for usage example.

## Features
- Roles
  - Roles
  - Privileges
- Users
  - Change password
  - Roles

*Note: No support for detecting server side changes at the moment. Please make sure there are none, terraform should be the owner of the created roles and users*

## Contributing
I am a Golang noob, please correct any idiomatic mistakes and don't shame me for them.

All the code is based on the examples in [Terraform: Writing Custom Providers](https://www.terraform.io/docs/extend/writing-custom-providers.html) as well as inspiration from [`terraform-provider-kafka`](https://github.com/Mongey/terraform-provider-kafka).
