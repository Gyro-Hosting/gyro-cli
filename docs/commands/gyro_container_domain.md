## gyro container domain

Domains of a service

```
gyro container domain [flags]
```

### Options

```
  -h, --help   help for domain
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

* [gyro container](gyro_container.md)	 - Gyro Containers (ECS), Istanbul: services, versions, logs, domains
* [gyro container domain add](gyro_container_domain_add.md)	 - Attach a domain
* [gyro container domain list](gyro_container_domain_list.md)	 - Domains of a service
* [gyro container domain remove](gyro_container_domain_remove.md)	 - Detach a domain
* [gyro container domain verify](gyro_container_domain_verify.md)	 - Re-check pending domains

