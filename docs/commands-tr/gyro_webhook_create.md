## gyro webhook create

Webhook oluştur (gizli anahtar bir kez döner)

### Synopsis

Webhook oluştur (gizli anahtar bir kez döner)

Teslimler `POST` JSON'dur; başlıklar `X-Gyro-Event`, `X-Gyro-Delivery`, `X-Gyro-Timestamp` ve `X-Gyro-Signature: sha256=HMAC_SHA256(secret, timestamp + "." + body)` (beyaz etiketli bayi: `X-Webhook-*`). Zaman damgası 5 dakikadan eskiyse reddedin. Tekrar: üstel geri çekilme, 8 deneme; 20 ardışık hata webhook'u kapatır.

API: POST /cloud/webhooks (operationId createWebhook, scope webhooks)

```
gyro webhook create [flags]
```

### Options

```
      --customer-id int          Reseller main key only
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --description string       description
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --events strings           empty = all
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

