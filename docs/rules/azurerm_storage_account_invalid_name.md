# azurerm_storage_account_invalid_name

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

In this rule, the string must match the regular expression `^[a-z0-9]{3,24}$``.

## Example

```hcl
resource "azurerm_storage_account" "foo" {
  name = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." does not match valid pattern ^[a-z0-9]{3,24}$ (azurerm_storage_account_invalid_name)

  on template.tf line 15:
  15:   name = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.4.0/docs/rules/azurerm_storage_account_invalid_name.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/resource-name-rules#microsoftstorage