# Azure KeyVault Secrets Kubernetes Operator 🚧

> This project is a work in progress

This Operator was created to keep Azure Key Vault secrets in sync with Kubernetes secrets. This allows all secrets to be managed in Azure Key Vault and keeps secrets out of the CI/CD process. Additionally, only the secrets specified will be pulled from the vault so you don't need `List` permissions on your vault.

## Getting Started

First, deploy the CRDs.

### Deploy using the yamls

```
curl https://git.io/JedVi | kubectl apply -f -

```
