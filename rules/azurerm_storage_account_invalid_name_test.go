package rules

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/golang/mock/gomock"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

var rule = AzurermStorageAccountInvalidNameRule(AzurermStorageAccountInvalidNameRule{
	resourceType:  "azurerm_storage_account",
	attributeName: "name",
	pattern:       regexp.MustCompile(`^[a-z0-9]{3,24}$`),
})

func Test_AzurermStorageAccountInvalidName_invalid(t *testing.T) {
	for _, test := range testCases {
		testDisplay := fmt.Sprintf("%d - %s", test.testID, test.testName)
		t.Run(testDisplay, func(t *testing.T) {
			//arrange
			runner := helper.TestRunner(t, map[string]string{"instances.tf": test.hcl})
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//act
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			//assert
			helper.AssertIssues(t, test.expected, runner.Issues)

		})
	}
}

var testCases = []struct {
	hcl      string
	testID   int
	testName string
	expected helper.Issues
}{

	{
		testID:   1,
		testName: "Valid Name",
		hcl: `
			resource "azurerm_storage_account" "invalid" {
				name = "thisisvalid123"
			}
		`,
		expected: helper.Issues{},
	},

	{
		testID:   2,
		testName: "Upper Case Characters not allowed",
		hcl: `
			resource "azurerm_storage_account" "invalid" {
				name = "Notavalidname"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    &rule,
				Message: "\"Notavalidname\" does not match valid pattern ^[a-z0-9]{3,24}$",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 12},
					End:      hcl.Pos{Line: 3, Column: 27},
				},
			},
		},
	},

	{
		testID:   3,
		testName: "Non Alpha Numeric Characters not allowed",
		hcl: `
			resource "azurerm_storage_account" "invalid" {
				name = "notavalid_name!"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    &rule,
				Message: "\"notavalid_name!\" does not match valid pattern ^[a-z0-9]{3,24}$",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 12},
					End:      hcl.Pos{Line: 3, Column: 29},
				},
			},
		},
	},

	{
		testID:   4,
		testName: "Non Alpha Numeric Characters with numbers not allowed",
		hcl: `
			resource "azurerm_storage_account" "invalid" {
				name = "Notavalid_name!3"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    &rule,
				Message: "\"Notavalid_name!3\" does not match valid pattern ^[a-z0-9]{3,24}$",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 12},
					End:      hcl.Pos{Line: 3, Column: 30},
				},
			},
		},
	},

	{
		testID:   5,
		testName: "Less than 3 characters is Invalid",
		hcl: `
			resource "azurerm_storage_account" "invalid" {
				name = "ab"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    &rule,
				Message: "\"ab\" does not match valid pattern ^[a-z0-9]{3,24}$",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 12},
					End:      hcl.Pos{Line: 3, Column: 16},
				},
			},
		},
	},

	{
		testID:   6,
		testName: "Greater than 24 characters is invalid",
		hcl: `
			resource "azurerm_storage_account" "invalid" {
				name = "abcdefghijklmnopqrstuvwxyz"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    &rule,
				Message: "\"abcdefghijklmnopqrstuvwxyz\" does not match valid pattern ^[a-z0-9]{3,24}$",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 12},
					End:      hcl.Pos{Line: 3, Column: 40},
				},
			},
		},
	},

	{
		testID:   7,
		testName: "Between 3 and 24 characters is valid",
		hcl: `
			resource "azurerm_storage_account" "invalid" {
				name = "abc1235"
			}
		`,
		expected: helper.Issues{},
	},
}
