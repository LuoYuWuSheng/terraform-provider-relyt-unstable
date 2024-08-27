terraform {
  required_providers {
    relyt = {
      source = "LuoYuWuSheng/relyt-unstable"
    }
  }
}

provider "relyt" {
  role     = var.role
  api_host = "http://global-gcs-inner-lb-0bc39606d8a0f702.elb.us-east-1.amazonaws.com"
}

resource "relyt_dwuser" "dw_user1" {
  dwsu_id                                  = var.dwsu_id
  account_name                             = var.account_name
  account_password                         = var.account_password
  datalake_aws_lakeformation_role_arn      = var.datalake_aws_lakeformation_role_arn
  async_query_result_location_prefix       = var.async_query_result_location_prefix
  async_query_result_location_aws_role_arn = var.async_query_result_location_aws_role_arn
}

data "relyt_dwsu_boto3_access_info" "boto3" {
  dwsu_id    = relyt_dwuser.dw_user1.dwsu_id
  account_id = relyt_dwuser.dw_user1.account_name
}
