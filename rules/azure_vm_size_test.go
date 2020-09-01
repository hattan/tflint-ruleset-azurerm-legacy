package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAzureVmSizes_ValidUrl_ResultArrayGreaterThanZero(t *testing.T) {
	//arrange
	validUrl := "https://tflintrulesstore.z5.web.core.windows.net/vm-size.json"

	//act
	sizes := GetAzureVmSizesWithUrl(validUrl)

	//assert
	assert.True(t, len(sizes) > 0)
}

func Test_GetAzureVmSizes_InvalidUrl_ResultArrayGreaterThanZero(t *testing.T) {
	//arrange
	validUrl := "https://tflintrulesstore.z5.web.core.windows.net/somenonexistentfile.json"

	//act
	sizes := GetAzureVmSizesWithUrl(validUrl)

	//assert
	assert.True(t, len(sizes) > 0)
}
