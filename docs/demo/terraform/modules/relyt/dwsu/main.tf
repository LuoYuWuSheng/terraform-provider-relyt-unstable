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
  #   resource_check_timeout = 30
}

resource "relyt_dwsu" "dwsu" {
  cloud  = var.cloud
  region = var.region
  domain = var.domain
  alias  = "terraform-test"
  default_dps = {
    name        = var.name
    description = "terraform test"
    engine      = var.engine
    size        = var.size
  }
}


