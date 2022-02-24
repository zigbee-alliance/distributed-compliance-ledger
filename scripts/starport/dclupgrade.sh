starport scaffold module dclupgrade --dep dclauth,upgrade

# Change `plan` field type to cosmos.upgrade.v1beta1.Plan after scaffolding
starport scaffold --module dclupgrade message ProposeUpgrade plan:string

starport scaffold --module dclupgrade message ApproveUpgrade name:string

# Change `plan` field type to cosmos.upgrade.v1beta1.Plan after scaffolding.
# Then change index to plan.name field and remove redundant name field.
starport scaffold --module dclupgrade map ProposedUpgrade plan:string creator:string approvals:array.string --index name:string --no-message

# Change `plan` field type to cosmos.upgrade.v1beta1.Plan after scaffolding.
# Then change index to plan.name field and remove redundant name field.
starport scaffold --module dclupgrade map ApprovedUpgrade plan:string creator:string approvals:array.string --index name:string --no-message