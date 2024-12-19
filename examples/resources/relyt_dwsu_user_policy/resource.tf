resource "relyt_dwsu_user_policy" "security_constraints" {

  dwsu_id               = "your_dwsu_id"
  mfa                   = "OPTIONAL"
  mfa_protection_scopes = ["DPS_OPERATIONS"]
  reset_init_password   = false
}