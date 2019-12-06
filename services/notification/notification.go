package notification

import (
	"github.com/kanhaiya15/gopf/services/apiclient"
)

// Alert alert
func Alert(alertURL string, requestData string) (interface{}, error) {
	responseData, err := apiclient.PostService(alertURL, requestData)
	if err != nil {
		return responseData, err
	}
	return responseData, nil
}
