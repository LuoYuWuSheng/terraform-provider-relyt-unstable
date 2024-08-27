variable "role" {
  type        = string
  default     = "SYSTEMADMIN"
  description = "Role"
}

variable "dwsu_id" {
  type        = string
  description = "DWSU Id"
}

variable "account_name" {
  type        = string
  default     = "user1@zbyte-inc.com"
  description = "Account name"
}

variable "account_password" {
  type        = string
  default     = "User1123."
  description = "Account password"
}

variable "datalake_aws_lakeformation_role_arn" {
  type        = string
  default     = "arn:aws:iam::905418298243:role/lake-r1"
  description = "Datalake AWS lakeformation role arn"
}

variable "async_query_result_location_prefix" {
  type        = string
  default     = "s3://relytqaresult-us-east-1/user1/result1/"
  description = "Async query result location prefix"
}

variable "async_query_result_location_aws_role_arn" {
  type        = string
  default     = "arn:aws:iam::905418298243:role/lake-r1"
  description = "Async query result location aws role arn"
}
