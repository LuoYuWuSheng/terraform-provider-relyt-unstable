terraform {
  required_providers {
    relyt = {
      source = "relytcloud/relyt"
    }
  }
}

provider "relyt" {
  role     = "SYSTEMADMIN"
}

resource "relyt_dwsu_integration_info" "integration_info" {
  dwsu_id = var.dwsu_id
  integration_info = {
    external_id = var.external_id
  }
}