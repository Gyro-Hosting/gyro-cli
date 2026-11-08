## gyro domain update

Alan Adı Tercihlerini Değiştirme

### Synopsis

Alan Adı Tercihlerini Değiştirme

Otomatik yenilemeyi ve fatura profilini değiştirir. — iki alan

API: PATCH /domains/{domain} (operationId coreUpdateDomain, scope Domains/UpdateDomain)

```
gyro domain update <DOMAIN> [flags]
```

### Options

```
      --auto-renew               Alan adı bazlı otomatik yenileme. Yalnız etkin alan adlarında ve gerçek bir mantıksal değerle.
      --billing-profile-id int   Fatura profilinin numarası. Sıfır hesabın varsayılanına döner.
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for update
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

