variable "role" {
  type        = string
  default     = "SYSTEMADMIN"
  description = "Role"
}

variable "dwsu_id" {
  type        = string
  description = "DWSU id"
}

variable "external_id" {
  type        = string
  default     = "relyt_dwsu_integration_test"
  description = "External id"
}
