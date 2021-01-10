package statuspage

import (
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sp "github.com/sbecker59/statuspage-api-client-go/api/v1/statuspage"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var pageID string
var organizationID string

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"statuspage": testAccProvider,
	}
	pageID = os.Getenv("STATUSPAGE_PAGE_ID")
	organizationID = os.Getenv("STATUSPAGE_ORGANIZATION_ID")
}

func isAPIKeySet() bool {
	if os.Getenv("SP_API_KEY") != "" {
		return true
	}
	if os.Getenv("STATUSPAGE_API_KEY") != "" {
		return true
	}
	return false
}

func isPageIDSet() bool {
	if os.Getenv("STATUSPAGE_PAGE_ID") != "" {
		return true
	}
	return false
}

func isOrganizationIDSet() bool {
	if os.Getenv("STATUSPAGE_ORGANIZATION_ID") != "" {
		return true
	}
	return false
}

// testAccPreCheck validates the necessary test API keys exist
// in the testing environment
func testAccPreCheck(t *testing.T) {
	if !isAPIKeySet() {
		t.Fatal("STATUSPAGE_API_KEY or SP_API_KEY must be set for acceptance tests")
	}
	if !isPageIDSet() {
		t.Fatal("STATUSPAGE_PAGE_ID must be set for acceptance tests")
	}
	if !isOrganizationIDSet() {
		t.Fatal("STATUSPAGE_ORGANIZATION_ID must be set for acceptance tests")
	}
}

func buildStatuspageClientV1(httpclient *http.Client) *sp.APIClient {
	configV1 := sp.NewConfiguration()
	configV1.UserAgent = getUserAgent(configV1.UserAgent)
	return sp.NewAPIClient(configV1)
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
