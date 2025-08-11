resource "aws_iam_role" "this_iam_role" {
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
  role       = aws_iam_role.this_iam_role.name
  policy_arn = aws_iam_policy.this_cloudwatch_write_policy.arn
}

resource "aws_iam_instance_profile" "this_iam_instance_profile" {
  role = aws_iam_role.this_iam_role.name
  tags = var.tags
}
