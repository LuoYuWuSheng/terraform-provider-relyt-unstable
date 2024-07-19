
resource "relyt_dwuser" "user1" {
  dwsu_id                                  = "dwsu-id-from-an-dwsu-resource"
  account_name                             = "UniqueAccountName"
  account_password                         = "daf#$dgdfe&Abce%64"
  datalake_aws_lakeformation_role_arn      = "role arn arn=//xxxx"
  async_query_result_location_prefix       = "s3=//bucket-name/prefix/..."
  async_query_result_location_aws_role_arn = "role arn arn=//xxxx"
}