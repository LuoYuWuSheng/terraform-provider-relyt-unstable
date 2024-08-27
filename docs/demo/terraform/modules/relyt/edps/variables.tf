variable "role" {
  type        = string
  default     = "SYSTEMADMIN"
  description = "Role"
}

variable "engine" {
  type        = string
  default     = "extreme"
  description = "Engine"
}

variable "size" {
  type        = string
  default     = "S"
  description = "Size"
}

variable "dwsu_id" {
  type        = string
  description = "DWSU id"
}

variable "edps_name" {
  type        = string
  default     = "edps-terraform-test"
  description = "EDPS name"
}