## gyro ticket close

Talebi Kapatma

### Synopsis

Talebi Kapatma

Talebi çözüldü olarak kapatır. — puan verilebilir

API: POST /tickets/{id}/close (operationId coreCloseTicket, scope Tickets/CloseTicket)

```
gyro ticket close <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for close
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --rating int               Kapatırken verilen puan. Yalnız talep henüz puansızken işlenir.
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

