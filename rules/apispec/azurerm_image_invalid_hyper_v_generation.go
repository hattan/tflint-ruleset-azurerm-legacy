// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermImageInvalidHyperVGenerationRule checks the pattern is valid
type AzurermImageInvalidHyperVGenerationRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermImageInvalidHyperVGenerationRule returns new rule with default attributes
func NewAzurermImageInvalidHyperVGenerationRule() *AzurermImageInvalidHyperVGenerationRule {
	return &AzurermImageInvalidHyperVGenerationRule{
		resourceType:  "azurerm_image",
		attributeName: "hyper_v_generation",
		enum: []string{
			"V1",
			"V2",
		},
	}
}

// Name returns the rule name
func (r *AzurermImageInvalidHyperVGenerationRule) Name() string {
	return "azurerm_image_invalid_hyper_v_generation"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermImageInvalidHyperVGenerationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermImageInvalidHyperVGenerationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermImageInvalidHyperVGenerationRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermImageInvalidHyperVGenerationRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as hyper_v_generation`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}