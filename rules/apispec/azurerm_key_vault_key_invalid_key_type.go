// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermKeyVaultKeyInvalidKeyTypeRule checks the pattern is valid
type AzurermKeyVaultKeyInvalidKeyTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermKeyVaultKeyInvalidKeyTypeRule returns new rule with default attributes
func NewAzurermKeyVaultKeyInvalidKeyTypeRule() *AzurermKeyVaultKeyInvalidKeyTypeRule {
	return &AzurermKeyVaultKeyInvalidKeyTypeRule{
		resourceType:  "azurerm_key_vault_key",
		attributeName: "key_type",
		enum: []string{
			"EC",
			"EC-HSM",
			"RSA",
			"RSA-HSM",
			"oct",
		},
	}
}

// Name returns the rule name
func (r *AzurermKeyVaultKeyInvalidKeyTypeRule) Name() string {
	return "azurerm_key_vault_key_invalid_key_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermKeyVaultKeyInvalidKeyTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermKeyVaultKeyInvalidKeyTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermKeyVaultKeyInvalidKeyTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermKeyVaultKeyInvalidKeyTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as key_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}