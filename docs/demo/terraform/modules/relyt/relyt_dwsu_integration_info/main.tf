terraform {
  required_providers {
    relyt = {
      source = "LuoYuWuSheng/relyt-unstable"
    }
  }
}

provider "relyt" {
  role     = "SYSTEMADMIN"
  api_host = "http://global-gcs-inner-lb-0bc39606d8a0f702.elb.us-east-1.amazonaws.com"
}

resource "relyt_dwsu_integration_info" "integration_info" {
  dwsu_id = var.dwsu_id
  integration_info = {
    external_id = var.external_id
  }
}