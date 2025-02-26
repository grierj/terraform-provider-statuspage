---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "statuspage_page_access_group Resource - terraform-provider-statuspage"
subcategory: ""
description: |-
  
---

# statuspage_page_access_group (Resource)



## Example Usage

```terraform
resource "statuspage_page_access_group" "my_user_group" {
  page_id    = "my_page_id"
  name       = "My Page Access User Group"
  users      = ["${statuspage_page_access_user.my_user.id}"]
  components = ["${statuspage_component.my_component.id}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) Name for this Group.
- **page_id** (String) the ID of the page this component group belongs to

### Optional

- **components** (Set of String) An array with the IDs of the components in this group
- **external_identifier** (String) Associates group with external group
- **id** (String) The ID of this resource.
- **metrics** (Set of String) An array with the IDs of the metrics in this group
- **users** (Set of String) An array with the Page Access User IDs that are in this group


