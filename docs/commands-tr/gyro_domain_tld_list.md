## gyro domain tld list

Uzantı Fiyatlarını Okuma

### Synopsis

Uzantı Fiyatlarını Okuma

Satıştaki uzantıları süreleri ve fiyatlarıyla döndürür. — cüzdan para birimi

API: GET /domains/tlds (operationId coreGetTlds, scope Domains/GetTlds)

```
gyro domain tld list [flags]
```

### Options

```
  -h, --help         help for list
      --tld string   Tek bir uzantı. Listeyi daraltır ve süre başına fiyat tablosunu ekler.
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

* [gyro domain tld](gyro_domain_tld.md)	 - Uzantı Fiyatlarını Okuma

