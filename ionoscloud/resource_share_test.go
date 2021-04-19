package ionoscloud

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v5"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccShare_Basic(t *testing.T) {
	var share ionoscloud.GroupShare
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckShareDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckShareConfig_basic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckShareExists("ionoscloud_share.share", &share),
					resource.TestCheckResourceAttr("ionoscloud_share.share", "share_privilege", "true"),
				),
			},
			{
				Config: testAccCheckShareConfig_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ionoscloud_share.share", "share_privilege", "false"),
				),
			},
		},
	})
}

func testAccCheckShareDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*ionoscloud.APIClient)

	for _, rs := range s.RootModule().Resources {

		grpId := rs.Primary.Attributes["group_id"]
		resourceId := rs.Primary.Attributes["resource_id"]
		_, apiResponse, err := client.UserManagementApi.UmGroupsSharesFindByResourceId(context.TODO(), grpId, resourceId).Execute()

		if err != nil {
			return fmt.Errorf("share for resource %s still exists in group %s %s", resourceId, grpId, apiResponse.Response.Status)
		}
	}

	return nil
}

func testAccCheckShareExists(n string, share *ionoscloud.GroupShare) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*ionoscloud.APIClient)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckShareExists: Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		grpId := rs.Primary.Attributes["group_id"]
		resourceId := rs.Primary.Attributes["resource_id"]
		foundshare, _, err := client.UserManagementApi.UmGroupsSharesFindByResourceId(context.TODO(), grpId, resourceId).Execute()

		if err != nil {
			return fmt.Errorf("Error occured while fetching Share of resource  %s in group %s", rs.Primary.Attributes["resource_id"], rs.Primary.Attributes["group_id"])
		}
		if *foundshare.Id != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		share = &foundshare

		return nil
	}
}

const testAccCheckShareConfig_basic = `
resource "ionoscloud_datacenter" "foobar" {
	name       = "terraform test"
	location = "us/las"
}

resource "ionoscloud_group" "group" {
  name = "terraform test"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
}

resource "ionoscloud_share" "share" {
  group_id = "${ionoscloud_group.group.id}"
  resource_id = "${ionoscloud_datacenter.foobar.id}"
  edit_privilege = true
  share_privilege = true
}`

const testAccCheckShareConfig_update = `
resource "ionoscloud_datacenter" "foobar" {
	name       = "terraform test"
	location = "us/las"
}

resource "ionoscloud_group" "group" {
  name = "terraform test"
  create_datacenter = true
  create_snapshot = true
  reserve_ip = true
  access_activity_log = false
}

resource "ionoscloud_share" "share" {
group_id = "${ionoscloud_group.group.id}"
  resource_id = "${ionoscloud_datacenter.foobar.id}"
  edit_privilege = true
  share_privilege = false
}
`
