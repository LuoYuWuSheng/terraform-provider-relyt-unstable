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
  #   resource_check_timeout = 100
}

resource "relyt_dps" "edps" {
  dwsu_id     = var.dwsu_id
  name        = var.edps_name
  description = "terraform edps test"
  engine      = var.engine
  size        = var.size
}
