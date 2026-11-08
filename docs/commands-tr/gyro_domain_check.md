## gyro domain check

Alan Adının Boş Olup Olmadığını Sorma

### Synopsis

Alan Adının Boş Olup Olmadığını Sorma

Bir alan adının alınabilirliğini canlı sorar. — kayıt kuruluşuna sorar

API: POST /domains/check (operationId coreCheckDomain, scope Domains/CheckDomain)

```
gyro domain check [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --domain string            Sorulacak tam alan adı.
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for check
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

* [gyro domain](gyro_domain.md)	 - Alan adları: sorgu, kayıt, transfer, ad sunucuları, DNS, WHOIS

