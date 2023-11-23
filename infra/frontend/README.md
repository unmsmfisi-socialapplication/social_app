## Local Testing with Docker

To locally test the Terraform configuration using Docker, follow these steps:

1. Pull the latest Terraform Docker image:
   ```bash
    docker pull hashicorp/terraform:latest
2. Initialize the Terraform working directory:
    ```bash
    docker run --rm -it -v $pwd\:/fe -w /fe hashicorp/terraform:latest init
3. Generate and view the Terraform plan:
     ```bash
     docker run --rm -it -v $pwd\:/fe -w /fe hashicorp/terraform:latest plan -var-file="dev.tfvars"

**Warning**: Before running the Docker commands, ensure that you have defined your own Google Cloud credentials path in the dev.tfvars file. Failing to do so may result in authentication issues.