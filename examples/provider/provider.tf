
terraform {
  required_providers {
    relyt = {
      source  = "LuoYuWuSheng/relyt"
      version = ">= 0.0.13"
    }
  }
}

provider "relyt" {
  auth_key = "9a3727e5b9c0ddabaGbll2HVLVKLLY1AyjOilAqeyPOBAb74A7VlJRAdTi0bJWJd3"
  role     = "SYSTEMADMIN"
}