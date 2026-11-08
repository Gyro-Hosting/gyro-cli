## gyro domain list

Alan Adlarını Listeleme

### Synopsis

Alan Adlarını Listeleme

Hesabın alan adlarını döndürür. — anahtarın sahibi

API: GET /domains (operationId coreGetDomains, scope Domains/GetDomains)

```
gyro domain list [flags]
```

### Options

```
      --all             tüm sayfaları getir
  -h, --help            help for list
      --limit int       sayfa başına kayıt (en çok 100)
      --page int        sayfa numarası (default 1)
      --search string   Alan adında arar.
      --status string   Durum süzgeci. Rozet gruplarını alır.
      --tld string      Yalnız bu uzantının alan adları.
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

* [gyro domain](gyro_domain.md)	 - Alan adları: sorgu, kayıt, transfer, ad sunucuları, DNS, WHOIS

