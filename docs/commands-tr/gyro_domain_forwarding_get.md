## gyro domain forwarding get

Adres Yönlendirmesini Okuma

### Synopsis

Adres Yönlendirmesini Okuma

Alan adının başka bir adrese yönlendirilip yönlendirilmediğini döndürür. — ek hizmet gerekir

API: GET /domains/{domain}/forwarding (operationId coreGetUrlForwarding, scope Domains/GetUrlForwarding)

```
gyro domain forwarding get <DOMAIN> [flags]
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

* [gyro domain forwarding](gyro_domain_forwarding.md)	 - Adres Yönlendirmesini Okuma

