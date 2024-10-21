---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "relyt_dwsu_schemas Data Source - relyt"
subcategory: ""
description: |-
  
---

# relyt_dwsu_schemas (Data Source)



## Example Usage

```terraform
data "relyt_dwsu_schemas" "schemas" {
  database = "your_database_name"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `database` (String) The name of the database.

### Read-Only

- `schemas` (Attributes List) The list of schema. (see [below for nested schema](#nestedatt--schemas))

<a id="nestedatt--schemas"></a>
### Nested Schema for `schemas`

Read-Only:

- `catalog` (String) The catalog of the schema. null is returned if the schema is not an external schema.
- `database` (String) The database of the schema.
- `external` (Boolean) Whether the schema is an external schema. true indicates yes; false indicates no.
- `name` (String) The name of the schema.
- `owner` (String) The owner of schema.