## gyro me security get

Güvenlik Özetini Okuma

### Synopsis

Güvenlik Özetini Okuma

Doğrulama durumunu, iki adımlı girişi ve son oturumları döndürür. — değiştirilemez

API: GET /me/security (operationId coreGetSecurity, scope Account/GetSecurity)

```
gyro me security get [flags]
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

* [gyro me security](gyro_me_security.md)	 - Güvenlik Özetini Okuma

