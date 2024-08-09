

resource "relyt_dwsu_integration_info" "integration_info" {
  dwsu_id = "dwsu-id-from-an-duws-resource"
  integration_info = {
    external_id     = "external id. can be update"
    relyt_principal = "relyt user arn"
    relyt_vpc       = "relyt vpc"
  }
}