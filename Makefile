TERRAFORM_DIR = infra

init:
	@cd $(TERRAFORM_DIR) && terraform init

plan:
	@cd $(TERRAFORM_DIR) && terraform plan

apply:
	@cd $(TERRAFORM_DIR) && terraform apply -auto-approve

destroy:
	@cd $(TERRAFORM_DIR) && terraform destroy -auto-approve

validate:
	@cd $(TERRAFORM_DIR) && terraform validate

fmt:
	@cd $(TERRAFORM_DIR) && terraform fmt

help:
	@echo "Available commands:"
	@echo "  make init        Initialize Terraform"
	@echo "  make plan        Generate execution plan"
	@echo "  make apply       Apply changes"
	@echo "  make destroy     Destroy infrastructure"
	@echo "  make validate    Validate configuration"
	@echo "  make fmt         Format Terraform files"

config:
	@bash config.sh

dev:
	@bash config.sh --dev
