## gyro completion

Generate a shell completion script

### Synopsis

Load completions:
  bash:  source <(gyro completion bash)        # or: gyro completion bash > /etc/bash_completion.d/gyro
  zsh:   gyro completion zsh > "${fpath[1]}/_gyro"
  fish:  gyro completion fish > ~/.config/fish/completions/gyro.fish
  pwsh:  gyro completion powershell | Out-String | Invoke-Expression
Server, SSH key and other ids complete from the API.

```
gyro completion bash|zsh|fish|powershell [flags]
```

### Options

```
  -h, --help   help for completion
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

