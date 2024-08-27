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

resource "relyt_dwsu_privatelink" "pl_api" {
  dwsu_id      = var.dwsu_id
  service_type = "data_api"
  allow_principals = [
    { principal = "*" }
  ]
}

resource "relyt_dwsu_privatelink" "pl_db" {
  dwsu_id      = var.dwsu_id
  service_type = "database"
  allow_principals = [{ principal = "*" }]
}
