package testhelper

import (
	"fmt"
	"testing"

	"github.com/scrapli/scrapligo/driver/core"
)

// SendConfigsFromFileTestHelper helper function to handle send configs from file tests.
func SendConfigsFromFileTestHelper(driverName string) func(t *testing.T) {
	sessionFile := fmt.Sprintf("../../test_data/driver/network/sendconfigs/%s", driverName)

	return func(t *testing.T) {
		d, driverErr := core.NewCoreDriver(
			"localhost",
			driverName,
			WithPatchedTransport(sessionFile),
		)

		if driverErr != nil {
			t.Fatalf("failed creating test device: %v", driverErr)
		}

		openErr := d.Open()
		if openErr != nil {
			t.Fatalf("failed opening patched driver: %v", openErr)
		}

		r, cmdErr := d.SendConfigsFromFile(
			fmt.Sprintf(
				"../../test_data/driver/network/sendconfigsfromfile/%s_configs",
				driverName,
			),
		)
		if cmdErr != nil {
			t.Fatalf("failed sending configs: %v", cmdErr)
		}

		if r.Failed() {
			t.Fatal("response object indicates failure")
		}
	}
}
