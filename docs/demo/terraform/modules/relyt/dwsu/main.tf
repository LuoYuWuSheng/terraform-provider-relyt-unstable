terraform {
  required_providers {
    relyt = {
      source = "relytcloud/relyt"
    }
  }
}
provider "relyt" {
  role     = var.role
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


