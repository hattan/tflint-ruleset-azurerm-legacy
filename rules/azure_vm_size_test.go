package rules

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAzureVmSizes_ValidUrl_ResultArrayGreaterThanZero(t *testing.T) {
	//arrange
	validUrl := "https://tflintrulesstore.z5.web.core.windows.net/vm-size.json"
	localJsonFile := ".tflint-azure-config/vm-size.json"

	//act
	sizes := GetAzureVmSizesWithUrl(validUrl, localJsonFile)

	//assert
	assert.True(t, len(sizes) > 0)
}

func Test_GetAzureVmSizes_InvalidUrl_FileDoesNotExist_ResultArrayGreaterThanZero(t *testing.T) {
	//arrange
	validUrl := "https://tflintrulesstore.z5.web.core.windows.net/somenonexistentfile.json"
	localJsonFile := ".tflint-azure-config/vm-size.json"

	// Ensure that the file does not exist
	if fileExists(localJsonFile) {
		err := os.Remove(localJsonFile)
		if err != nil {
			log.Fatalln(err)
		}
	}

	//act
	sizes := GetAzureVmSizesWithUrl(validUrl, localJsonFile)

	//assert
	assert.True(t, len(sizes) > 0)
}

func Test_GetAzureVmSizes_InvalidUrl_FileExists_ResultArrayGreaterThanZero(t *testing.T) {
	//arrange
	validUrl := "https://tflintrulesstore.z5.web.core.windows.net/somenonexistentfile.json"
	localJsonFile := ".tflint-azure-config/vm-size.json"

	// Ensure that the local cached file exists
	if !fileExists(localJsonFile) {
		generateVMSizeJSON(localJsonFile)
	}

	//act
	sizes := GetAzureVmSizesWithUrl(validUrl, localJsonFile)

	//assert
	assert.True(t, len(sizes) > 0)
}
