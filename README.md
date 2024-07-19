# Terraform Provider Relyt 

```
    terraform {
      required_providers {
        relyt = {
          source  = "relytcloud/relyt"
          version = "0.0.2"
        }  
      }
    }
    
    provider "relyt" {
      auth_key = "<api_key>" # Copy the API key you obtained from the previous step.
      role     = "SYSTEMADIN" # The system role for the Relyt cloud account to create, fixed to SYSTEMADMIN.
    }
    
    
    locals {
      cloud_id = {
        id = "aws"
      }
      region_id = {
        id = "<region_id>" # Set the region in which the DW service unit will be created.
      }
      BASIC = {
        id = "basic"
      }
    }
    
    # Create a DW service unit and its default Hybrid DPS cluster.
      resource "relyt_dwsu" "dwsu_example" {
      cloud     = local.cloud_id.id
      region    = local.region_id.id
      domain    = "hdps-example-tf" # The subdomain, customizable.
      alias     = "hdps-test" # The alias of the DW service unit.
      default_dps = {
        name        = "hdps-test" # The name for the Hybrid DPS cluster, customizable.
        description = "The Hybrid DPS cluster" # A short description for the Hybrid DPS cluster, customizable and optional.
        engine      = "hybrid" # The type of the DPS cluster, fixed to hybrid.
        size        = "s" # The size of the Hybrid DPS cluster, customizable.
    }
    }
    
    # Create an Extreme DPS cluster.
    resource "relyt_dps" "edps_example" {
      dwsu_id     = relyt_dwsu.edps_example.id
      name        = "edps1" # Set the name for the Extreme DPS cluster.
      description = "An Extreme DPS cluster" # A short description for the Extreme DPS cluster, customizable.
      engine      = "extreme" # The type of the DPS cluster, fixed to extreme.
      size        = "xs" # The size of the Extreme DPS cluster.
    }
    
    # Create a DW user. You can repeat this code block to create multiple DW users.
    resource "relyt_dwuser" "user1" {
      dwsu_id          = relyt_dwsu.dwsu_example.id
      account_name     = "user1" # Name the DW user to create.
      account_password = "Qwer123!" # The password for the DW user, which must be 8 to 32 characters in length and contain at least one uppercase letter, one lowercase letter, one digit, and one special character.
    
    
    # Optional parameters
      datalake_aws_lakeformation_role_arn = "anotherRole2" # The ARN of the cross-account IAM role, optional.
      async_query_result_location_prefix  = "simple"       # The prefix of the path to the S3 output location.
      async_query_result_location_aws_role_arn = "anotherSimple" # The ARN of the role to access the output location, optional.
    }
    
    
    data "relyt_dwsu_service_account" "tes" {
      dwsu_id = relyt_dwsu.dwsu_example.id
    }
```