## gyro lb cert letsencrypt

Ücretsiz Let's Encrypt sertifikası iste

### Synopsis

Ücretsiz Let's Encrypt sertifikası iste

A kaydı bu yük dengeleyiciye yönlenmiş en çok 10 ad. Hedefleriniz /.well-known/acme-challenge/<token> isteklerini GET certificates'taki adrese yönlendirmeli; `acme=waiting` yanıtı ad başına kontrolü verir, yönlendirmeyi düzeltip verify çağırın.

API: POST /cloud/load-balancers/{id}/certificates/letsencrypt (operationId typed_post_load_balancers__id__certificates_letsencrypt, scope services_write)

```
gyro lb cert letsencrypt <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --domains string           Comma separated names
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for letsencrypt
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

