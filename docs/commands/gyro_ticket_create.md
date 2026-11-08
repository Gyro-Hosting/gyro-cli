## gyro ticket create

Create ticket

### Synopsis

Create ticket

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/destek-isteme

API: POST /tickets (operationId coreCreateTicket, scope Tickets/CreateTicket)

```
gyro ticket create [flags]
```

### Options

```
      --attachments strings      Ekler. Kodlanmış içerik olarak gönderilir; adres reddedilir.
      --credentials string       Erişim bilgileri. Her satır bir grup numarası ve alan haritası taşır; mesaja şifreli işlenir.
      --custom-fields string     Departman alanlarının yanıtları. Alan numarasını değere eşler; kutu tipli alanlar dizi alır. (JSON)
      --data string              request body as JSON, @file or @- (stdin)
      --department-id int        Departmanın numarası.
      --dry-run                  show the price quote / request without changing anything
      --encrypt-message          Mesaj gövdesi şifreli saklansın mı.
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --message string           Açılış mesajı. En az beş karakter; düz metindir ve işaretleme ayıklanır.
      --priority int             Öncelik. Verilmezse orta alınır.
      --service-id int           Talebi bir hizmete bağlar.
      --set stringArray          extra body field key=value (repeatable)
      --subject string           Konu.
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

