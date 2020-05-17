// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermNotificationHubNamespaceInvalidNamespaceTypeRule checks the pattern is valid
type AzurermNotificationHubNamespaceInvalidNamespaceTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermNotificationHubNamespaceInvalidNamespaceTypeRule returns new rule with default attributes
func NewAzurermNotificationHubNamespaceInvalidNamespaceTypeRule() *AzurermNotificationHubNamespaceInvalidNamespaceTypeRule {
	return &AzurermNotificationHubNamespaceInvalidNamespaceTypeRule{
		resourceType:  "azurerm_notification_hub_namespace",
		attributeName: "namespace_type",
		enum: []string{
			"Messaging",
			"NotificationHub",
		},
	}
}

// Name returns the rule name
func (r *AzurermNotificationHubNamespaceInvalidNamespaceTypeRule) Name() string {
	return "azurerm_notification_hub_namespace_invalid_namespace_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermNotificationHubNamespaceInvalidNamespaceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermNotificationHubNamespaceInvalidNamespaceTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermNotificationHubNamespaceInvalidNamespaceTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermNotificationHubNamespaceInvalidNamespaceTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as namespace_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}