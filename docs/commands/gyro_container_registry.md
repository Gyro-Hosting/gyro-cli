## gyro container registry

Gyro image registry of the environment

```
gyro container registry [flags]
```

### Options

```
  -h, --help   help for registry
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
* [gyro container registry credentials](gyro_container_registry_credentials.md)	 - Registry password
* [gyro container registry get](gyro_container_registry_get.md)	 - Gyro image registry of the environment
* [gyro container registry image-delete](gyro_container_registry_image-delete.md)	 - Delete an image tag
* [gyro container registry images](gyro_container_registry_images.md)	 - Images in the registry
* [gyro container registry import](gyro_container_registry_import.md)	 - Import an image from another registry
* [gyro container registry imports](gyro_container_registry_imports.md)	 - Image imports (progress and history)
* [gyro container registry policy](gyro_container_registry_policy.md)	 - Image clean-up policy
* [gyro container registry rotate](gyro_container_registry_rotate.md)	 - Rotate the registry password

