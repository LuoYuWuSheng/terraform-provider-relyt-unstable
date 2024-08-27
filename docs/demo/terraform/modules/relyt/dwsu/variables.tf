variable "role" {
  type        = string
  default     = "SYSTEMADMIN"
  description = "Role"
}

variable "cloud" {
  type        = string
  default     = "aws"
  description = "Cloud"
}

variable "region" {
  type        = string
  default     = "us-east-1"
  description = "Region"

}

variable "name" {
  type        = string
  default     = "terraform-test"
  description = "Name"
}

variable "domain" {
  type        = string
  default     = "terraform-test"
  description = "Domain"
}

variable "engine" {
  type        = string
  default     = "hybrid"
  description = "Engine"
}

variable "size" {
  type        = string
  default     = "S"
  description = "Size"
}
