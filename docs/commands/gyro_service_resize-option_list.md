## gyro service resize-option list

Products this service can be upgraded to, with prices

### Synopsis

Products this service can be upgraded to, with prices

Prepaid services answer `billing: prepaid`; each option has `plan`, `current` and one `hourly` price with `monthly_cap`. **Data safety (Gyro Cloud Istanbul, tr-ist-1)**: a plan change keeps the disk and all its data, the IP addresses, additional disks and snapshots. The root disk is enlarged in place and the file system is extended automatically; a disk never shrinks, so a plan with a smaller disk is refused with 409 `disk_shrink` (also in `dry_run`; resize-options marks such plans `allowed: false`). Bigger CPU/RAM is applied live when the hypervisor allows it, otherwise the server is shut down cleanly and started again (measured: 1-2 -> 2-4 restarted, about 12 seconds without SSH; 2-4 -> 4-8 live, no restart). While the server has snapshots the disk cannot grow: 409 `resize_refused` (delete the snapshots first). The answer carries `data_preserved: true`, `disk: grows_only`, `restart: if_needed`. Only a rebuild (`POST /cloud/servers/{id}/actions/rebuild`, needs `confirm: true`) erases data: the disk, additional disks and snapshots are deleted (IP addresses stay). **`recreate_required`** (422, invoiced products only): the target plan would need the service to be re-created on another platform, which deletes its data; the API never does that and nothing is changed or charged. Pick another plan from resize-options, or create a new server, move the data and delete the old one.

API: GET /cloud/services/{id}/resize-options (operationId getResizeOptions, scope services_read)

```
gyro service resize-option list <ID> [flags]
```

### Options

```
  -h, --help   help for list
```

### Options inherited from parent commands

```
      --api-url string     API base URL (env GYRO_API_URL)
      --debug              log HTTP requests to stderr (secrets masked)
      --lang string        message language: tr or en (default from locale)
      --no-color           disable colours (also NO_COLOR)
      --no-header          omit the table header
  -o, --output string      output format: table, json, yaml, text
  -p, --profile string     profile to use (env GYRO_PROFILE)
      --query string       JMESPath query applied to the result
  -q, --quiet              print only ids
      --timeout duration   overall timeout, e.g. 30s, 10m
      --wide               show all table columns
  -y, --yes                do not ask for confirmation
```

### SEE ALSO

* [gyro service resize-option](gyro_service_resize-option.md)	 - Products this service can be upgraded to, with prices

