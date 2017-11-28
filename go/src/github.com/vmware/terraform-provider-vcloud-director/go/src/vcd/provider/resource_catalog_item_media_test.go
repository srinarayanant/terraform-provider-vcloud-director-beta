/*****************************************************************
* terraform-provider-vcloud-director
* Copyright (c) 2017 VMware, Inc. All Rights Reserved.
* SPDX-License-Identifier: BSD-2-Clause
******************************************************************/

package provider

import (
	"fmt"
	"testing"

	//"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/vmware/terraform-provider-vcloud-director/go/src/util/logging"
	"github.com/vmware/terraform-provider-vcloud-director/go/src/vcd/proto"
)

func TestAccResourceCatalogItemMedia(t *testing.T) {
	logging.Plog("__INIT__TestAccResourceCatalogItemMedia")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCatalogItemDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCatalogItem_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCatalogItemUpload(),
				),
			},
		},
	})

	logging.Plog("__DONE__TestAccResourceCatalogItemMedia")
}

func testAccCheckCatalogItemUpload() resource.TestCheckFunc {
	return func(s *terraform.State) error {

		logging.Plog("__INIT__testAccCheckCatalogItemUpload")
		m := testAccProvider.Meta()

		vcdClient := m.(*VCDClient)

		provider := vcdClient.getProvider()
		logging.Plog(fmt.Sprintf("%#v", provider))
		//TEST against VCD that the catalog item is uploaded

		catalogIsPresInfo := proto.IsPresentCatalogItemInfo{CatalogName: "testcata1", ItemName: "item1"}
		isPresResp, isPreErr := provider.IsPresentCatalogItem(catalogIsPresInfo)

		if isPreErr != nil {
			return fmt.Errorf("Error in validating catalog item creation")
		}
		if !isPresResp.Present {
			return fmt.Errorf("Error Catlog Item not created as expected")
		}

		logging.Plog("__DONE__testAccCheckCatalogItemUpload")
		return nil

	}
}

func testAccCheckCatalogItemDestroy(s *terraform.State) error {

	logging.Plog("__INIT__testAccCheckCatalogItemDestroy_")

	m := testAccProvider.Meta()

	vcdClient := m.(*VCDClient)

	provider := vcdClient.getProvider()
	logging.Plog(fmt.Sprintf("%#v", provider))
	//TEST against VCD that the catalog item is DELETED

	catalogIsPresInfo := proto.IsPresentCatalogItemInfo{CatalogName: "testcata1", ItemName: "item1"}
	isPresResp, isPreErr := provider.IsPresentCatalogItem(catalogIsPresInfo)

	if isPreErr != nil {
		return fmt.Errorf("Error in validating catalog item deletion")
	}
	if isPresResp.Present {
		return fmt.Errorf("Error Catlog Item not deleted as expected")
	}

	logging.Plog("__DONE__testAccCheckCatalogItemDestroy_")
	return nil

}

const testAccCatalogItem_basic = `
variable "MEDIA_PATH" { 

 type    = "string"
 default = "nonepath" 
}

variable "VCD_IP" { 

 type    = "string"
 default = "notdefinedip" 
}

variable "VCD_USER" { 

 type    = "string"
 default = "notdefineduser" 
}
variable "VCD_PASSWORD" { 

 type    = "string"
 default = "notdefinedpass" 
}
variable "VCD_ORG" { 

 type    = "string"
 default = "notdefinedorg" 
}


provider "vcloud-director" {
  user                 = "${var.VCD_USER}"
  password             = "${var.VCD_PASSWORD}"
  org                  = "${var.VCD_ORG}"
  ip                   = "${var.VCD_IP}"
  allow_unverified_ssl = "true"
}



resource "vcloud-director_catalog" "catalog1" {
        name    = "testcata1"
        description = "desc"
        shared  = "true"

}

resource "vcloud-director_catalog_item_media" "item1" {
	item_name = "item1"
	catalog_name= "${vcloud-director_catalog.catalog1.name}"

	file_path="${var.MEDIA_PATH}"
}
`

///Users/srinarayana/vmws/file1.txt
