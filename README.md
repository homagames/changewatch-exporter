ChangeWatch Exporter
====================

ChangeWatch Exporter aim to keep checking remote versionned content and expose it as Prometheus metrics. 

# Motivation
This tool is developed by the Homa's lovely infra team who got tired of checking for new updates on external dependencies such as Helm repository or external binaries from github releases, etc.
It also help you to prioritize updates and have a better understanding of the lag of your infra.

# Use case

ChangeWatch Exporter use what we call `monitors` to keep track of an external dependencies. You will have to record the version of the version of the dependencie you can to monitor and the ChangeExporter will export a prometheus metric indicating with there is a change and it will add the type of change (`major`, `minor`, `patch`) as label. 

# What can we monitor

1. Github release

You can watch a Github Release by adding this monitor to your config file
```
  - kind: GithubRelease
    apiVersion: v1
    metadata:
      name: <name>
    spec:
      org: <github_org>
      repo:  <github_repo>
      version: <version>
```

With:
* `<name>`: the name you want to assign to this monitor
* `<github_org>`: the github organisation of the release to watch
* `<github_repo>`: the gitub repository of the release to watch
* `<version>`: the version you currently are on your environment

2. Formated remote file

You can watch a remote file by adding this monitor to your config file. This is useful as not all the dependencies are made a with a Github release like most of the Helm chart.
```
  - kind: FormatedRemoteFile
    apiVersion: v1
    metadata:
      name: <name>
    spec:
      url:  <url>
      version: <version>
      format: <format>
      key: <key>
```
With:
* `<name>`: the name you want to assign to this monitor
* `<url>`: the url of the remote file
* `<version>`: the version you currently are on your environment
* `<format>`: the format of the remote file to decode the version
* `<key>`: the key of the file to read the version


# Exemple of usage

```
./changewatch-exporter --config config.example.yaml
```
With [config.example.yaml](config.example.yaml)


# Install it on Kubernetes

See dedicated [helm chart](https://github.com/homagames/helm-charts/tree/master/charts/changewatch-exporter).