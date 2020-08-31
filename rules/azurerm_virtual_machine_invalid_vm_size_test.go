package rules

import (
	"testing"

	"github.com/golang/mock/gomock"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

var vmInvalidSizeRule = AzurermVirtualMachineInvalidVMSizeRule(AzurermVirtualMachineInvalidVMSizeRule{
	resourceType:  "azurerm_linux_virtual_machine",
	attributeName: "vm_size",
})

func Test_AzurermVirtualMachineInvalidVMSizeRule(t *testing.T) {
	//arrange
	content := `
			resource "azurerm_linux_virtual_machine" "test" {
			vm_size = "Basic_A0"
		}`

	expected := helper.Issues{
		{
			Rule:    &vmInvalidSizeRule,
			Message: "\"Basic_A0\" is an invalid value as vm_size",
			Range: hcl.Range{
				Filename: "instances.tf",
				Start:    hcl.Pos{Line: 3, Column: 14},
				End:      hcl.Pos{Line: 3, Column: 24},
			},
		},
	}

	runner := helper.TestRunner(t, map[string]string{"instances.tf": content})
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//act
	if err := vmInvalidSizeRule.Check(runner); err != nil {
		t.Fatalf("Unexpected error occurred: %s", err)
	}

	//assert
	helper.AssertIssues(t, expected, runner.Issues)
}
