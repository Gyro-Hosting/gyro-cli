## gyro lb cert

Sertifikaları listele

```
gyro lb cert [flags]
```

### Options

```
  -h, --help   help for cert
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

* [gyro lb](gyro_lb.md)	 - Yük dengeleyicileri listele
* [gyro lb cert cancel](gyro_lb_cert_cancel.md)	 - Bekleyen Let's Encrypt isteğini iptal et
* [gyro lb cert delete](gyro_lb_cert_delete.md)	 - Sertifikayı sil
* [gyro lb cert letsencrypt](gyro_lb_cert_letsencrypt.md)	 - Ücretsiz Let's Encrypt sertifikası iste
* [gyro lb cert list](gyro_lb_cert_list.md)	 - Sertifikaları listele
* [gyro lb cert upload](gyro_lb_cert_upload.md)	 - Sertifika yükle
* [gyro lb cert verify](gyro_lb_cert_verify.md)	 - Bekleyen Let's Encrypt isteğini sürdür

