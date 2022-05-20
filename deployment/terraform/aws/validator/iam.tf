resource "aws_iam_role" "this_amp_role" {
  count = var.enable_prometheus ? 1 : 0

  name = "amp-role"

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
}

resource "aws_iam_policy" "this_amp_write_policy" {
  count = var.enable_prometheus ? 1 : 0

  name        = "amp-write-policy"
  description = "AMP Write Policy"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "aps:RemoteWrite"
            ],
            "Effect": "Allow",
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "this_amp_policy_attachment" {
  count = var.enable_prometheus ? 1 : 0

  role       = aws_iam_role.this_amp_role[0].name
  policy_arn = aws_iam_policy.this_amp_write_policy[0].arn
}

resource "aws_iam_instance_profile" "this_amp_role_profile" {
  count = var.enable_prometheus ? 1 : 0
  name = "prometheus-node-profile"
  role = aws_iam_role.this_amp_role[0].name
}