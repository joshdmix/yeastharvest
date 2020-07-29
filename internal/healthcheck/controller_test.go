package healthcheck_test

import "github.com/joshdmix/yeastharvest/internal/healthcheck"

func TestHealthcheck(t *testing.T) {
	test.EndpointTest(t, 200, "GET", "/healthcheck", healthcheck.Healthcheck(, nil, "", nil))
}