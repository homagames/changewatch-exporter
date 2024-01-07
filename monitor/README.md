Kinds
=====

Here the list of the monitors we can define.

# FormatedRemoteFile

| Field | Description |Value |
| ----- | ----- | ----- |
| `Kind` | The Kind of the Monitor | `FormatedRemoteFile` |
| `apiVersion` | Version of the Monitor | `v1` |
| `metadata.name` | This name of the monitor | `""` |
| `metadata.labels` | Labels to attach to the prometeus metrics | `{}` |
| `spec.url` | URL of the remote file | `""` |
| `spec.version` | The version of the remote file that you have on your infrastructure | `""` |
| `spec.format`   | Format of the remote file | `yaml` |
| `spec.key` | The key of the version to keep track of the change | `.version` |

# GithubRelease

| Field | Description |Value |
| ----- | ----- | ----- |
| `Kind` | The Kind of the Monitor | `GithubRelease` |
| `apiVersion` | Version of the Monitor | `v1` |
| `metadata.name` | This name of the monitor | `""` |
| `metadata.labels` | Labels to attach to the prometeus metrics | `{}` |
| `spec.repo` | The Github repo to monitor | `""` |
| `spec.version` | The version of the release you have on your infrastructure | `""` |

