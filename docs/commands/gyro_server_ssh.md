## gyro server ssh

Open an SSH session to a server

### Synopsis

Looks up the server's IP and runs your local ssh client. Everything after "--" is passed to ssh,
so you can run a command remotely. Defaults come from the profile (ssh_user, ssh_identity).

```
gyro server ssh <SERVER> [-- SSH-ARGS/COMMAND...] [flags]
```

### Examples

```
  gyro server ssh web-1
  gyro server ssh 47466 -i ~/.ssh/id_ed25519 -- uptime
  gyro server ssh web-1 --ipv6 -- -L 8080:localhost:80
```

### Options

```
  -h, --help              help for ssh
  -i, --identity string   private key file
  -6, --ipv6              connect over IPv6
      --port int          SSH port (default 22)
  -u, --user string       remote user (default root)
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

* [gyro server](gyro_server.md)	 - Cloud servers: create, power, snapshots, SSH ...

