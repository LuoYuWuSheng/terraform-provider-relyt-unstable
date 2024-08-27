terraform {
  required_providers {
    relyt = {
      source = "relytcloud/relyt"
    }
  }
}

provider "relyt" {
  role     = var.role
  #   resource_check_timeout = 100
}

resource "relyt_dps" "edps" {
  dwsu_id     = var.dwsu_id
  name        = var.edps_name
  description = "terraform edps test"
  engine      = var.engine
  size        = var.size
}
