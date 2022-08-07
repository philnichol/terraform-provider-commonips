package provider

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCommonIPsCloudflare(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCommonIPsCloudflareConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCommonIPsCloudflare("data.commonips_cloudflare.some"),
				),
			},
		},
	})
}

func testAccCommonIPsCloudflare(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		var (
			cidrBlockSize int
			err           error
		)

		if cidrBlockSize, err = strconv.Atoi(a["cidr_blocks.#"]); err != nil {
			return err
		}

		if cidrBlockSize < 10 {
			return fmt.Errorf("cidr_blocks seem suspiciously low: %d", cidrBlockSize)
		}

		var cidrBlocks sort.StringSlice = make([]string, cidrBlockSize)

		for i := range make([]string, cidrBlockSize) {
			block := a[fmt.Sprintf("cidr_blocks.%d", i)]

			if _, _, err := net.ParseCIDR(block); err != nil {
				return fmt.Errorf("malformed CIDR block %s: %w", block, err)
			}

			cidrBlocks[i] = block
		}

		if !sort.IsSorted(cidrBlocks) {
			return fmt.Errorf("unexpected order of cidr_blocks: %s", cidrBlocks)
		}

		return nil
	}
}

const testAccCommonIPsCloudflareConfig = `
data "commonips_cloudflare" "some" {}
`
