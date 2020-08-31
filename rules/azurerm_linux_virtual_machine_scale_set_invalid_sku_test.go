package rules

import (
	"testing"

	"github.com/golang/mock/gomock"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

var linuxVmScaleSetInvalidSkuRule = AzurermLinuxVirtualMachineScaleSetInvalidSkuRule(AzurermLinuxVirtualMachineScaleSetInvalidSkuRule{
	resourceType:  "azurerm_linux_virtual_machine_scale_set",
	attributeName: "sku",
})

func Test_AzurermLinuxVirtualMachineScaleSetInvalidSkuRule(t *testing.T) {
	//arrange
	content := `
			resource "azurerm_linux_virtual_machine_scale_set" "test" {
			sku = "Basic_A0"
		}`

	expected := helper.Issues{
		{
			Rule:    &linuxVmScaleSetInvalidSkuRule,
			Message: "\"Basic_A0\" is an invalid value as sku",
			Range: hcl.Range{
				Filename: "instances.tf",
				Start:    hcl.Pos{Line: 3, Column: 10},
				End:      hcl.Pos{Line: 3, Column: 20},
			},
		},
	}

	runner := helper.TestRunner(t, map[string]string{"instances.tf": content})
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//act
	if err := linuxVmScaleSetInvalidSkuRule.Check(runner); err != nil {
		t.Fatalf("Unexpected error occurred: %s", err)
	}

	//assert
	helper.AssertIssues(t, expected, runner.Issues)
}
