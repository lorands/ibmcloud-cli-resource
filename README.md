# ibmcloud-cli-resource

IBM Cloud Cli (bx) Concourse Resource

**NOTE: This is a work in progress! Do NOT use it in production!**

Inspired by cf-cli-resource (https://github.com/nulldriver/cf-cli-resource). 
This is a dead simple resource for IBM Cloud Cli. All it does is calls up ibmcloud cli with the
given parameters (after a login) and pipes the output in json format to a file if requested.

Uses IBM Cloud client: https://cloud.ibm.com/docs/cli

An get only resource capable of running IBM Cloud cli commands. 

Note: Reason for being a get and not put only resource is, that for getting the output of a put is only
supported through the subsequent get, thus this would indicate a need of a persistent storage like s3, 
which would complicate this resource a lot. 

## Source Configuration

Note: you must provide either `username` and `password` or `client_id` and `client_secret`.

* `api`: *Optional.* The IBM Cloud API. Defaults to https://cloud.ibm.com
* `region`: *Optional.* The IBM Cloud region to use. Defaults to `eu-gb`
* `username`: *Required.* The username used to authenticate.
* `password`: *Required.* The password used to authenticate.
* `account_id`: *Optional.* The account id used to authenticate.
* `resource_group`: *Optional.* The resource group to use. Will use default defined by the account, usually 'Default'.

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

* `command`: *Required.* The command to execute
* `subcommand`: *Required.* Sub-command to execute
* `params`: *Optional.* List of additional parameters to pass to ibmcloud cli.
* `jsonOutputFile`: *Optional.* Set ibmcloud to produce a json output and feed it to a file by this name

Examples:

```yml
  - get: bx-resource
    resource: bx-dev
    params:
      command: resource
      subcommand: groups
      jsonOutputFile: myOutput.json
```

```yml
  - get: bx-resource
    resource: bx-dev
    params:
      command: resource
      subcommand: service-alias-create
      jsonOutputFile: sac.json
      params:
      - my-service-alias
      - --instance-name my-instance-11
      - -s mySpace
```

## PUT

Does nothing.