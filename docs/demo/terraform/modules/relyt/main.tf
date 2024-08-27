terraform {
  required_providers {
    relyt = {
      source = "LuoYuWuSheng/relyt-unstable"
    }
  }
}

provider "relyt" {
  api_host = "http://global-gcs-inner-lb-0bc39606d8a0f702.elb.us-east-1.amazonaws.com"
  role     = "SYSTEMADMIN"
}

module "dwsu" {
  source = "./dwsu"
}

module "edps" {
  source = "./edps"
  dwsu_id = module.dwsu.dwsu_id
}

module "dw_user" {
  source = "./dw_user"
  dwsu_id = module.dwsu.dwsu_id
}

module "integration_info" {
  source = "./relyt_dwsu_integration_info"
  dwsu_id = module.dwsu.dwsu_id
  external_id = "20240821"
}

module "privatelink" {
  source = "./relyt_privatelink"
  dwsu_id = module.dwsu.dwsu_id
}