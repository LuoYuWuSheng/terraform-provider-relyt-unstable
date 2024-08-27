output "aksk" {
  value = data.relyt_dwsu_boto3_access_info.boto3
}
output "user1_ak" {
  value = data.relyt_dwsu_boto3_access_info.boto3.boto3_access_infos[0].access_key
}

output "user1_sk" {
  value = data.relyt_dwsu_boto3_access_info.boto3.boto3_access_infos[0].secret_key
}