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
  tags               = var.tags
}

resource "aws_iam_policy" "this_amp_write_policy" {
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
  tags   = var.tags
}

resource "aws_iam_role_policy_attachment" "this_amp_policy_attachment" {
  role       = aws_iam_role.this_iam_role.name
  policy_arn = aws_iam_policy.this_amp_write_policy.arn
}

resource "aws_iam_instance_profile" "this_amp_role_profile" {
  role = aws_iam_role.this_iam_role.name
  tags = var.tags
}
