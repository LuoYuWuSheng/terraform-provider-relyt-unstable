---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "relyt_dwsu_database Resource - relyt"
subcategory: ""
description: |-
  
---

# relyt_dwsu_database (Resource)



## Example Usage

```terraform
resource "relyt_dwsu_database" "database" {
  name = "example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the database. The database name must not exceed 127 characters.

### Read-Only

- `owner` (String) The owner of the database.


## Import

Using `terraform import`, import database using the `database_name`. For example:
```
terraform import relyt_dwsu_database.database-import your_db_name
```
