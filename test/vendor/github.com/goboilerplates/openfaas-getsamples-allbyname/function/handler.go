package function

import (
	"encoding/json"
	"log"

	"github.com/goboilerplates/core"
)

var (
	api GetSamplesAllByNameAPI
)

// Handle a serverless request .
func Handle(req []byte) string {
	log.Println("Request with ", req)
	keyword := string(req)
	api = CreateAPI()
	result, err := api.AllByName(keyword)
	if err != nil {
		return err.Error()
	}
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return err.Error()
	}
	return string(jsonBytes)
}

// CreateAPI creates API for GetSamplesAll.
func CreateAPI() GetSamplesAllByNameAPI {
	if api == nil {
		return GetSamplesAllByNameAPIImpl{
			Interactor: core.CreateDefaultGetSamples("Kaka", "Ronaldo"),
		}
	}
	return api
}
