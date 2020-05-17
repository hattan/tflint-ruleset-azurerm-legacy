// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule checks the pattern is valid
type AzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule returns new rule with default attributes
func NewAzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule() *AzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule {
	return &AzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule{
		resourceType:  "azurerm_data_factory_integration_runtime_managed",
		attributeName: "edition",
		enum: []string{
			"Standard",
			"Enterprise",
		},
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule) Name() string {
	return "azurerm_data_factory_integration_runtime_managed_invalid_edition"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidEditionRule) Check(runner tflint.Runner) error {
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
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as edition`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}