## gyro user

List sub-users

```
gyro user [flags]
```

### Options

```
  -h, --help   help for user
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro user create](gyro_user_create.md)	 - Invite a sub-user
* [gyro user delete](gyro_user_delete.md)	 - Remove a sub-user
* [gyro user list](gyro_user_list.md)	 - List sub-users
* [gyro user resend-invite](gyro_user_resend-invite.md)	 - Resend the invitation e-mail
* [gyro user update](gyro_user_update.md)	 - Update a sub-user (label, status, permissions)

