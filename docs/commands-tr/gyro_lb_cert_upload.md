## gyro lb cert upload

Sertifika yükle

### Synopsis

Sertifika yükle

PEM sertifika zinciri + şifresiz özel anahtar. Anahtar yalnız yük dengeleyiciye gönderilir; panelde saklanmaz, loglanmaz.

API: POST /cloud/load-balancers/{id}/certificates (operationId typed_post_load_balancers__id__certificates, scope services_write)

```
gyro lb cert upload <ID> [flags]
```

### Options

```
      --certificate string       PEM chain
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for upload
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              name
      --private-key string       PEM key (unencrypted)
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

* [gyro lb cert](gyro_lb_cert.md)	 - Sertifikaları listele

