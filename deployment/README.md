# Terraform deployment

[Terraform][1] is an open-source infrastructure as code software tool that
codifies cloud APIs into declarative configuration files.

## Requirements

1. Install [Terraform][2] and [AWS][3] CLIs
2. Install Ansible requirements. `pip install -r requirements.txt`
3. Configure AWS access keys for CLI. `aws configure`
4. Modify the [deploy ansible playbook](ansible/deploy.yml) accordingly.

## Deployment

The following instructions automates the provision of a DCL node in AWS Cloud.

```bash
cd terraform/
terraform init
terraform apply -auto-approve
```

[1]: https://www.terraform.io/
[2]: https://learn.hashicorp.com/tutorials/terraform/install-cli
[3]: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
