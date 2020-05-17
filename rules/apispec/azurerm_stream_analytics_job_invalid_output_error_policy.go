// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule checks the pattern is valid
type AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule returns new rule with default attributes
func NewAzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule() *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule {
	return &AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule{
		resourceType:  "azurerm_stream_analytics_job",
		attributeName: "output_error_policy",
		enum: []string{
			"Stop",
			"Drop",
		},
	}
}

// Name returns the rule name
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Name() string {
	return "azurerm_stream_analytics_job_invalid_output_error_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as output_error_policy`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}