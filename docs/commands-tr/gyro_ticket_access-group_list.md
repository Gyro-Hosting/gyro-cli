## gyro ticket access-group list

Erişim Gruplarını Listeleme

### Synopsis

Erişim Gruplarını Listeleme

Mesaja eklenebilecek erişim bilgisi gruplarını şemalarıyla döndürür. — şifreli alanlar

API: GET /tickets/access-groups (operationId coreGetTicketAccessGroups, scope Tickets/GetTicketAccessGroups)

```
gyro ticket access-group list [flags]
```

### Options

```
  -h, --help   help for list
```

### Options inherited from parent commands

```
      --api-url string     API taban adresi (ortam GYRO_API_URL)
      --debug              HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string        mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color           renkleri kapat (NO_COLOR da olur)
      --no-header          tablo başlığını yazma
  -o, --output string      çıktı biçimi: table, json, yaml, text
  -p, --profile string     kullanılacak profil (ortam GYRO_PROFILE)
      --query string       sonuca uygulanacak JMESPath sorgusu
  -q, --quiet              yalnız kimlikleri yaz
      --timeout duration   toplam zaman aşımı, ör. 30s, 10m
      --wide               tüm tablo sütunlarını göster
  -y, --yes                onay sorma
```

### SEE ALSO

* [gyro ticket access-group](gyro_ticket_access-group.md)	 - Erişim Gruplarını Listeleme

