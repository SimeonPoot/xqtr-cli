# XQTR - executor

## Problem
The problem it should solve: 
When using a CI/CD and IAC/GITOPS, sometimes there's a need for a custom script to do something to an environment (Kubernetes for instance) for one time.

## Use-cases
- If cert-manager is used to create certificates, and the certificate-resource is deleted, the secret with the certificate/ca/key is not deleted. [Under the assumption that --enable-certificate-owner-ref isn't passed as a flag](https://cert-manager.io/docs/usage/certificate/#cleaning-up-secrets-when-certificates-are-deleted). 
  - Effectively, when this is not accounted for, in the deletion of the certificate-resource, a secret is living in the cluster with certs that should not be deployed there. 

When there's no access to a production environment prefered, a CI/CD pipeline is scheduled everyday to release to dev-envs and once a week to prod-envs, a simple script to remove the secret isn't sufficient. 
As, for instance, the pipeline is deploying master-git-branch everyday to dev-envs, once a week to test-envs and once a week to prod-envs, the simple script would need a trigger to go to these environments.

A solution would be to run a script with a timeframe in mind. More concrete, the script will be run within the date it has been provided with. 
So in the context of the use-case: 
- a specific cmd will be run (kubectl delete secret cert-secret)
- a timeframe is given when this cmd is allowed to run, for instance;
  - Monday (dev-env master), Wednesday (test-env master) and Friday (prod-env master).

When the master is deployed with the script (delete secret), this will be run with the command for only the day it is configured.
- preferably, next to a date, a git-commit is also given to validate master.

### xqtr-cli:
Go [COBRA]-cli that runs a command (first phase a bash-script). This command can only be run in a timeframe of a day (later more specified).
