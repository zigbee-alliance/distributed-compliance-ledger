starport scaffold module dclupgrade --dep dclauth,upgrade

# Change `plan` field type to cosmos.upgrade.v1beta1.Plan after scaffolding
starport scaffold --module dclupgrade message ProposeUpgrade plan:string
