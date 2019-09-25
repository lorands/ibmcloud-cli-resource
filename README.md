# ibmcloud-cli-resource

IBM Cloud Cli (bx) Concourse Resource

**NOTE: This is a work in progress! Do NOT use it in production!**

Inspired by cf-cli-resource (https://github.com/nulldriver/cf-cli-resource). 
This is a dead simple resource for IBM Cloud Cli. All it does is calls up ibmcloud cli with the
given parameters (after a login) and pipes the output in json format to a file if requested.

Uses IBM Cloud client: https://cloud.ibm.com/docs/cli

An output only resource capable of running IBM Cloud cli commands.

## Source Configuration

Note: you must provide either `username` and `password` or `client_id` and `client_secret`.

* `region`: *Optional.* The IBM Cloud region to use. Defaults to `eu-gb`
* `username`: *Required.* The username used to authenticate.
* `password`: *Required.* The password used to authenticate.
* `account_id`: *Optional.* The account id used to authenticate.
* `resource_group`: *Optional.* The resource group to use. Will use default defined by the account, usually 'Default'.
* `org`: *Optional.* Sets the default organization to target (can be overridden in the params config).
* `space`: *Optional.* Sets the default space to target (can be overridden in the params config).

```yml
resource_types:
- name: bx-cli-resource
  type: docker-image
  source:
    repository: lorands/ibmcloud-cli-resource
    tag: latest

resources:
- name: bx-env
  type: bx-cli-resource
  source:
    region: eu-gb
    username: admin
    password: admin
    account_id: 4433443344
    resource_group: my-resource-group
```
## Check

Does nothing.

## GET

Call an ibmcloud cli and stores the output of json to a file if requested. 
Use this if there is no side effect.

```yml
  - get: bx-resource
    resource: bx-dev
    params:
      command: resource
      subcommand: groups
      jsonOutputFile: myOutput.json
```

## PUT

Call an ibmcloud cli and stores the output of json to a file if requested. 
Use this if there is a side effect (resource created).

```yml
  - put: bx-resource
    resource: bx-dev
    params:
      command: resource
      subcommand: group-create
      params:
      - my-resource-group
```

```yml
  - put: bx-resource
    resource: bx-dev
    params:
      command: resource
      subcommand: group-create
      params:
      - my-resource-group
      tags: 
      - tag1
      - tag2
```

