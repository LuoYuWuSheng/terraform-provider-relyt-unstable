

resource "relyt_dwsu_privatelink" "privatelink" {
  dwsu_id      = "dwsu-id-from-an-duws-resource"
  service_type = "private link target service type"
  allow_principals = [
    { principal = "*" }, { principal = "arn:aws:iam::093584080162:user/*" }
  ]
}