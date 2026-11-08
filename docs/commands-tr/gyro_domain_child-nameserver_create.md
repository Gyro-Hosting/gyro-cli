## gyro domain child-nameserver create

Kendi Ad Sunucusunu Tanımlama

### Synopsis

Kendi Ad Sunucusunu Tanımlama

Alan adının altında bir ad sunucusu tanımlar. — sicile yazar

API: POST /domains/{domain}/child-nameservers (operationId coreCreateChildNameserver, scope Domains/CreateChildNameserver)

```
gyro domain child-nameserver create <DOMAIN> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --host string              Sunucunun tam adı. Küçük harfe çevrilir.
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --ip string                Sunucunun adresi. Biçimi doğrulanır.
      --ns string                ns
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

* [gyro domain child-nameserver](gyro_domain_child-nameserver.md)	 - Kendi Ad Sunucularını Listeleme

