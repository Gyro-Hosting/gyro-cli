## gyro sms message

List sms messages

```
gyro sms message [flags]
```

### Options

```
  -h, --help   help for message
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

* [gyro sms](gyro_sms.md)	 - SMS messages
* [gyro sms message get](gyro_sms_message_get.md)	 - Get sms message (path: {id})
* [gyro sms message list](gyro_sms_message_list.md)	 - List sms messages

