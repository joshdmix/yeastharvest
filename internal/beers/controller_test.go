package beers_test

import (
	"testing"

	"github.com/joshdmix/yeastharvest/internal/beers"
)

func TestGetBeers(t *testing.T) {
	test.EndpointTest(t, 200, "GET", "/beers", beers.GetBeers())
}
