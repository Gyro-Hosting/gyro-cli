## gyro service label get

Hizmetin takma adını ve etiketlerini getir

### Synopsis

Hizmetin takma adını ve etiketlerini getir

Hizmetlerinizden birinin (her ürün hattı) müşteri tanımlı takma adı ve renkli etiketleri. Etiketsiz hizmet `"nickname": null` ve `"tags": []` döner. Etiketler yalnız hizmet sahibine görünür; devirden sonra eski sahibin etiketi gösterilmez.

API: GET /cloud/services/{id}/label (operationId getServiceLabel, scope services_read)

```
gyro service label get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro service label](gyro_service_label.md)	 - Hizmetin takma adını ve etiketlerini getir

