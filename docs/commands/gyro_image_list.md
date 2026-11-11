## gyro image list

List operating system images

### Synopsis

List operating system images

API: GET /cloud/images (operationId listImages, scope cloud_read)

```
gyro image list [flags]
```

### Options

```
  -h, --help             help for list
      --plan string      With region: only the images of this plan in that region.
      --product-id int   Only the images offered on this product (implies its region).
      --region string    Only images offered in this region (e.g. tr-ist-1). Without it: the images of the global cloud catalogue (unchanged).
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

* [gyro image](gyro_image.md)	 - Operating system images

