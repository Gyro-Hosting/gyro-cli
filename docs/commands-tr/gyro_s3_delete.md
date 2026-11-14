## gyro s3 delete

Depolama hizmetini iptal et

### Synopsis

Depolama hizmetini iptal et

`confirm=true` zorunlu. `immediate=true` hemen iptal eder: erişim anahtarları, web sitesi ve özel alan adları yaklaşık bir dakikada durur, tüm kovalar ve nesneler dakikalar içinde kalıcı olarak silinir (geri alınamaz; Korumalı Depolama'da COMPLIANCE kilidi ya da legal hold altındaki sürümler kilit bitene kadar kalır ve ücretlenir). `immediate` olmadan iptal talebi açılır, işlenene kadar hizmet çalışır.

API: DELETE /cloud/storage/{id} (operationId deleteStorageService, scope cloud_delete)

```
gyro s3 delete <ID> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for delete
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --immediate                
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

* [gyro s3](gyro_s3.md)	 - S3 nesne depolama: kovalar ve erişim anahtarları

