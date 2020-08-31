package rules

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAzureVmSizes() []string {
	// TODO: Cache file locally.  Check to see if file exists, if not save it

	resp, err := http.Get("https://tflintrulesstore.z5.web.core.windows.net/vm-size.json")

	// TODO: If there is an error retrieving the file, use local cached file instead
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var sizeData []string
	json.Unmarshal(body, &sizeData)

	return sizeData
}
