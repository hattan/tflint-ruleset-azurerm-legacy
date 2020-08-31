package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermWindowsVirtualMachineInvalidSizeRule checks the pattern is valid
type AzurermWindowsVirtualMachineInvalidSizeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermWindowsVirtualMachineInvalidSizeRule returns new rule with default attributes
func NewAzurermWindowsVirtualMachineInvalidSizeRule() *AzurermWindowsVirtualMachineInvalidSizeRule {
	sizeData := GetAzureVmSizes()

	return &AzurermWindowsVirtualMachineInvalidSizeRule{
		resourceType:  "azurerm_windows_virtual_machine",
		attributeName: "size",
		enum:          sizeData,
	}
}

// Name returns the rule name
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Name() string {
	return "azurerm_windows_virtual_machine_invalid_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as size`, val),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
