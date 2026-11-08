## gyro ticket reply

Yanıt Yazma

### Synopsis

Yanıt Yazma

Açık bir talebe yanıt ekler. — ek kabul eder

API: POST /tickets/{id}/reply (operationId coreReplyTicket, scope Tickets/ReplyTicket)

```
gyro ticket reply <ID> [flags]
```

### Options

```
      --attachments strings      Ekler.
      --credentials string       Erişim bilgileri.
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --encrypt-message          Gövde şifreli saklansın mı.
  -h, --help                     help for reply
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --message string           Yanıt metni. En az beş karakter.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro ticket](gyro_ticket.md)	 - Destek talepleri
* [gyro ticket reply rating](gyro_ticket_reply_rating.md)	 - Personel Yanıtını Puanlama

