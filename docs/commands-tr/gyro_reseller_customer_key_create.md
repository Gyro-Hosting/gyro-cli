## gyro reseller customer key create

Uç-müşteri için anahtar üret (token bir kez döner)

### Synopsis

Uç-müşteri için anahtar üret (token bir kez döner)

API: POST /cloud/reseller/customers/{id}/keys (operationId createResellerCustomerKey, scope reseller)

```
gyro reseller customer key create <ID> [flags]
```

### Options

```
      --actions strings          Optional action allow list: server actions (reboot) or tool:action / tool:* / method:name
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --expires string           Anahtarın son kullanma tarihi: never (varsayılan), 7d, 30d, 90d, 1y, <n>d (en çok 1830) ya da tarih YYYY-MM-DD (gün sonu). Süresi dolan anahtar 401 key_expired ile reddedilir ve iptal edilir; 7 gün önce hesap sahibine e-posta gider.
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --ips strings              ips
      --name string              name
      --rate-limit int           Requests per minute (<= your own limit; 0 = default)
      --scopes strings           scopes
      --services int64Slice      Optional service allow list (default [])
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --template string          Hazır kapsam kümesi, scopes listesine eklenir: readonly (cloud_read, services_read), dns (services_read, services_write, dns_acme), cloud (cloud_*, services_*), containers / storage (services_read, services_write, cloud_read), full (devredilebilen tüm kapsamlar). (readonly|dns|cloud|containers|storage|full)
```

### Options inherited from parent commands

```
      --account gyro account list   panel kullanıcısı olduğunuz başka bir hesapta işlem yap (gyro account list id'si; "self" = kendi hesabınız; ortam GYRO_ACCOUNT, profil varsayılanı: `gyro account use`)
      --api-url string              API taban adresi (ortam GYRO_API_URL)
      --debug                       HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string                 mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color                    renkleri kapat (NO_COLOR da olur)
      --no-header                   tablo başlığını yazma
  -o, --output string               çıktı biçimi: table, json, yaml, text
  -p, --profile string              kullanılacak profil (ortam GYRO_PROFILE)
      --query string                sonuca uygulanacak JMESPath sorgusu
  -q, --quiet                       yalnız kimlikleri yaz
      --timeout duration            toplam zaman aşımı, ör. 30s, 10m
      --wide                        tüm tablo sütunlarını göster
  -y, --yes                         onay sorma
```

### SEE ALSO

* [gyro reseller customer key](gyro_reseller_customer_key.md)	 - Uç-müşterinin anahtarlarını listele

