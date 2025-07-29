resource "azurerm_resource_group" "rg" {
  name     = "rg-monitoring"
  location = "West Europe"
  tags     = var.tags
}


resource "azurerm_user_assigned_identity" "monitoring_identity" {
  name                = "monitoring-identity"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  tags                = var.tags
}

resource "azurerm_role_definition" "cloudwatch_equivalent" {
  name        = "CustomMonitoringRole"
  scope       = var.scope_id
  description = "Custom role: описанные права для мониторинга и логирования, аналог AWS CloudWatch Write"
  permissions {
    actions = [
      # autoscaling:Describe*
      "Microsoft.Compute/virtualMachineScaleSets/read",
      "Microsoft.Compute/virtualMachineScaleSets/list/action",

      # cloudwatch:* и logs:* → общий мониторинг и логирование
      "Microsoft.Insights/metrics/read",
      "Microsoft.Insights/activityLog/read",
      "Microsoft.Insights/alertRules/*",
      "Microsoft.OperationalInsights/workspaces/*",

      # sns:* → эквивалент Pub/Sub в Azure: EventGrid
      "Microsoft.EventGrid/eventSubscriptions/*",

      # iam:GetPolicy, iam:GetPolicyVersion, iam:GetRole
      "Microsoft.Authorization/roleDefinitions/read",
      "Microsoft.Authorization/roleAssignments/read",
      "Microsoft.Resources/subscriptions/resourceGroups/read"
    ]
    not_actions = []
  }
  assignable_scopes = [
    var.scope_id
  ]
  tags = var.tags
}
______________________
resource "azurerm_role_definition" "this_role_definition" {
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

  tags = var.tags
}

resource "aws_iam_policy" "this_cloudwatch_write_policy" {
  description = "Cloudwatch Write Policy"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "autoscaling:Describe*",
                "cloudwatch:*",
                "logs:*",
                "sns:*",
                "iam:GetPolicy",
                "iam:GetPolicyVersion",
                "iam:GetRole"
            ],
            "Effect": "Allow",
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": "iam:CreateServiceLinkedRole",
            "Resource": "arn:aws:iam::*:role/aws-service-role/events.amazonaws.com/AWSServiceRoleForCloudWatchEvents*",
            "Condition": {
                "StringLike": {
                    "iam:AWSServiceName": "events.amazonaws.com"
                }
            }
        }
    ]
}
EOF

  tags = var.tags
}

resource "aws_iam_role_policy_attachment" "this_cloudwatch_policy_attachment" {
  role       = azurerm_role_definition.this_role_definition.name
  policy_arn = aws_iam_policy.this_cloudwatch_write_policy.arn
}

resource "aws_iam_instance_profile" "this_iam_instance_profile" {
  role = azurerm_role_definition.this_role_definition.name
  tags = var.tags
}
