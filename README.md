# terraform-provider-mongodbacl

## Installation
Download source and build with `go build -o terraform-provider-mongodb`. Move resulting binary to `~/.terraform.d/plugins`.

The file name must be exactly `terraform-provider-mongodb` for this to work.

## Example
See `example.tf` for usage example.

## Contributing
I am a Golang noob, please correct any idiomatic mistakes and don't shame me for them.

All the code is based on the examples in [Terraform: Writing Custom Providers](https://www.terraform.io/docs/extend/writing-custom-providers.html) as well as inspiration from [`terraform-provider-kafka`](https://github.com/Mongey/terraform-provider-kafka).
