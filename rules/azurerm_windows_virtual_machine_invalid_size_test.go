package rules

import (
	"testing"

	"github.com/golang/mock/gomock"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

var windowsVmInvalidSizeRule = AzurermWindowsVirtualMachineInvalidSizeRule(AzurermWindowsVirtualMachineInvalidSizeRule{
	resourceType:  "azurerm_windows_virtual_machine",
	attributeName: "size",
})

func Test_AzurermWindowsVirtualMachineInvalidSizeRule(t *testing.T) {
	//arrange
	content := `
			resource "azurerm_windows_virtual_machine" "test" {
			size = "Basic_A0"
		}`

	expected := helper.Issues{
		{
			Rule:    &windowsVmInvalidSizeRule,
			Message: "\"Basic_A0\" is an invalid value as size",
			Range: hcl.Range{
				Filename: "instances.tf",
				Start:    hcl.Pos{Line: 3, Column: 11},
				End:      hcl.Pos{Line: 3, Column: 21},
			},
		},
	}

	runner := helper.TestRunner(t, map[string]string{"instances.tf": content})
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//act
	if err := windowsVmInvalidSizeRule.Check(runner); err != nil {
		t.Fatalf("Unexpected error occurred: %s", err)
	}

	//assert
	helper.AssertIssues(t, expected, runner.Issues)
}
