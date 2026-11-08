## gyro ticket reply

Reply ticket (path: {id})

### Synopsis

Reply ticket (path: {id})

Reply ticket (path: {id}). Support ticket of the calling account. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/destek-isteme).

API: POST /tickets/{id}/reply (operationId coreReplyTicket, scope Tickets/ReplyTicket)

```
gyro ticket reply <ID> [flags]
```

### Options

```
      --attachments strings      Ekler.
      --credentials string       Erişim bilgileri.
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --encrypt-message          Gövde şifreli saklansın mı.
  -h, --help                     help for reply
      --idempotency-key string   idempotency key (automatic for money operations)
      --message string           Yanıt metni. En az beş karakter.
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro ticket](gyro_ticket.md)	 - Support tickets
* [gyro ticket reply rating](gyro_ticket_reply_rating.md)	 - Rate ticket reply (path: {id}, {rid})

