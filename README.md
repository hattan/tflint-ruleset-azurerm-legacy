# TFLint Ruleset for terraform-provider-azurerm
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-azurerm/workflows/build/badge.svg?branch=master)](https://github.com/terraform-linters/tflint-ruleset-azurerm/actions)
[![GitHub release](https://img.shields.io/github/release/terraform-linters/tflint-ruleset-azurerm.svg)](https://github.com/terraform-linters/tflint-ruleset-azurerm/releases/latest)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

TFLint ruleset plugin for Terraform Provider for Azure (Resource Manager)

## Requirements

- TFLint v0.19+
- Go v1.14

## Installation

### Building the plugin

Clone the repository locally by running the following command:

```
git clone --recurse-submodules
```

Note: the `--recurse-submodules` flag is used to pull the files for the `azure-rest-api-specs` submodule.

Next, run the following command, which will build the plugin:

```
$ make
```

### Installing the plugin

After building, you can easily install the plugin using:

```
$ make install
```

This command will install the plugin and place it in `~/.tflint.d/plugins/tflint-ruleset-azurerm` (or `./.tflint.d/plugins/tflint-ruleset-azurerm`).


### Enabling the plugin

Create a file in the root of your Terraform project directory named `.tflint.hcl` if one does not already exist.

To enable the azurerm ruleset plugin for use, add the following configuration to the `.tflint.hcl` file:

```hcl
plugin "azurerm" {
    enabled = true
}
```

## Rules

200+ rules are available. See the [documentation](docs/README.md).


