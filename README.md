# Terraform Provider for Windows Active Directory (AD)

![Status: Tech Preview](https://img.shields.io/badge/status-experimental-EAAA32) [![Releases](https://img.shields.io/github/release/Sidoran/terraform-provider-ad.svg)](https://github.com/Sidoran/terraform-provider-ad/releases)
[![LICENSE](https://img.shields.io/github/license/Sidoran/terraform-provider-ad.svg)](https://github.com/Sidoran/terraform-provider-ad/blob/main/LICENSE)
![Unit tests](https://github.com/Sidoran/terraform-provider-ad/workflows/Unit%20tests/badge.svg)

This Windows AD provider for Terraform allows you to manage users, groups and group policies in your AD installation.

This provider is a technical preview, which means it's a community supported project. Please [file issues](https://github.com/Sidoran/terraform-provider-ad/issues/new/choose) generously and detail your experience while using the provider. We welcome your feedback.

By using the software in this repository (the AD provider), you acknowledge that:

* The AD provider is provided on an "as-is" basis, and may include bugs, errors, or other issues.
* The AD provider is NOT INTENDED FOR PRODUCTION USE, use of the Software may result in unexpected results, loss of data, or other unexpected results.

## Fork Notice

This project is a fork of [terraform-provider-ad](https://github.com/Sidoran/terraform-provider-ad), originally developed by [Hashicorp]. At the time of the fork creation  (May 2025) the original project has no changes for 2 years and a lot of uncommitted PRs. It is licensed under the Mozilla Public License 2.0 (MPL-2.0).

This fork is independently maintained and may diverge from the original project in features and functionality.

Significant modifications have been made since the fork. See the [CHANGELOG](./CHANGELOG.md) for details.

## Requirements

* [Terraform](https://www.terraform.io/downloads.html) version 0.12.x+
* [Windows Server](https://www.microsoft.com/en-us/windows-server) 2012R2 or greater
* [Go](https://golang.org/doc/install) version 1.16.x+

## Getting Started

If this is your first time here, you can get an overview of the provider by reading our [introductory blog post](https://www.hashicorp.com/blog/manage-active-directory-objects-new-windows-ad-provider-hashicorp-terraform). Otherwise, start by downloading a copy of our latest build from the [registry](https://registry.terraform.io/providers/hashicorp/ad/latest).

Once you have the plugin installed, review the [docs](docs/) folder to understand which configuration options are available. You can find examples and more in [our examples folder](examples/). Don't forget to run `terraform init` in your Terraform configuration directory to allow Terraform to detect the provider plugin.

## Contributing

We welcome your contribution. Please understand that the experimental nature of this repository means that contributing code may be a bit of a moving target. If you have an idea for an enhancement or bug fix, and want to take on the work yourself, please first [create an issue](https://github.com/Sidoran/terraform-provider-ad/issues/new/choose) so that we can discuss the implementation with you before you proceed with the work.

You can review our [contribution guide](_about/CONTRIBUTING.md) to begin. You can also check out our [frequently asked questions](_about/FAQ.md).
