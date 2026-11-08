## gyro webhook update

Adres, olay, açıklama ya da durumu güncelle

### Synopsis

Adres, olay, açıklama ya da durumu güncelle

API: PATCH /cloud/webhooks/{id} (operationId updateWebhook, scope webhooks)

```
gyro webhook update <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --description string       description
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --events strings           events
  -h, --help                     help for update
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --status string            status (active|disabled)
      --url string               url
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

* [gyro webhook](gyro_webhook.md)	 - Webhook'lar (imzalı olay bildirimleri)

