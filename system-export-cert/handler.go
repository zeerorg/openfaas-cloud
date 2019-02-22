package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	deployment := os.Getenv("ss_deployment")
	namespace := os.Getenv("ss_namespace")

	resp, err := http.Get(fmt.Sprintf("http://%s.%s:8080/v1/cert.pem", deployment, namespace))

	if err != nil {
		return err.Error()
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		return readErr.Error()
	}

	return string(body)
}
