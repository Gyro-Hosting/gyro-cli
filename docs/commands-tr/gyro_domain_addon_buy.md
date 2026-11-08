## gyro domain addon buy

Alan Adı Ek Hizmeti Alma

### Synopsis

Alan Adı Ek Hizmeti Alma

**Para harcar.** Alan adı yönetimi, gizlilik ya da yönlendirme ek hizmetini satın alır. — oranlı ilk dönem

API: POST /domains/{domain}/addons/{key} (operationId coreBuyDomainAddon, scope Domains/BuyDomainAddon)

```
gyro domain addon buy <DOMAIN> <KEY> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for buy
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

* [gyro domain addon](gyro_domain_addon.md)	 - Alan Adı Ek Hizmeti Alma

