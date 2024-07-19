

resource "relyt_dwsu" "dwsu" {
  cloud  = "aws"
  region = "ap-east-1"
  domain = "your-subdomain-prefix"
  alias  = "your-alias-of-dwsu"
  default_dps = {
    name        = "hybrid"
    description = "An Dwsu Example"
    engine      = "hybrid"
    size        = "S"
  }
}