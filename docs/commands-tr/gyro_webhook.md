## gyro webhook

Webhook'lar (imzalı olay bildirimleri)

```
gyro webhook [flags]
```

### Options

```
  -h, --help   help for webhook
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

* [gyro](gyro.md)	 - Gyro Hosting komut satırı aracı
* [gyro webhook create](gyro_webhook_create.md)	 - Webhook oluştur (gizli anahtar bir kez döner)
* [gyro webhook delete](gyro_webhook_delete.md)	 - Webhook'u sil
* [gyro webhook delivery](gyro_webhook_delivery.md)	 - Son 100 teslim
* [gyro webhook get](gyro_webhook_get.md)	 - Webhook'u getir
* [gyro webhook list](gyro_webhook_list.md)	 - Webhook'ları listele
* [gyro webhook test](gyro_webhook_test.md)	 - Şimdi bir api.test olayı gönder
* [gyro webhook update](gyro_webhook_update.md)	 - Adres, olay, açıklama ya da durumu güncelle

