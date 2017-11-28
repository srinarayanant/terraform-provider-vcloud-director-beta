#*****************************************************************
# terraform-provider-vcloud-director
# Copyright (c) 2017 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: BSD-2-Clause
#*****************************************************************

export TF_ACC=1
export TF_LOG=TRACE


export TF_VAR_VCD_IP="10.112.83.27"
export TF_VAR_VCD_USER="user1"
export TF_VAR_VCD_PASSWORD="Admin!23"
export TF_VAR_VCD_ORG="O1"

export TF_VAR_CATALOG_DESCRIPTION="accc_ description1"

export TF_VAR_OVA_PATH="/Users/srinarayana/vmws/tiny.ova"

export TF_VAR_MEDIA_PATH="/Users/srinarayana/vmws/file1.txt"

go test github.com/vmware/terraform-provider-vcloud-director/go/src/vcd/provider/ -v | grep --line-buffered -vE 'DEBUG|TRACE|terraform|^$'  