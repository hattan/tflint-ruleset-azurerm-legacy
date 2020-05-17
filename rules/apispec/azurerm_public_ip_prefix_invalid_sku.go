// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermPublicIPPrefixInvalidSkuRule checks the pattern is valid
type AzurermPublicIPPrefixInvalidSkuRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermPublicIPPrefixInvalidSkuRule returns new rule with default attributes
func NewAzurermPublicIPPrefixInvalidSkuRule() *AzurermPublicIPPrefixInvalidSkuRule {
	return &AzurermPublicIPPrefixInvalidSkuRule{
		resourceType:  "azurerm_public_ip_prefix",
		attributeName: "sku",
		enum: []string{
			"Standard",
		},
	}
}

// Name returns the rule name
func (r *AzurermPublicIPPrefixInvalidSkuRule) Name() string {
	return "azurerm_public_ip_prefix_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermPublicIPPrefixInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermPublicIPPrefixInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermPublicIPPrefixInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermPublicIPPrefixInvalidSkuRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as sku`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}