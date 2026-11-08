## gyro reseller customer update

Uç-müşteriyi güncelle (askıya alma = anahtarları çalışmaz)

### Synopsis

Uç-müşteriyi güncelle (askıya alma = anahtarları çalışmaz)

API: PATCH /cloud/reseller/customers/{id} (operationId updateResellerCustomer, scope reseller)

```
gyro reseller customer update <ID> [flags]
```

### Options

```
      --allow-order              May this customer's keys order servers (paid from YOUR balance)?
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --email string             email
      --external-ref string      Your own customer id (unique per reseller)
  -h, --help                     help for update
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              name
      --note string              note
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --status string            status (active|suspended)
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

* [gyro reseller customer](gyro_reseller_customer.md)	 - Uç-müşterileri listele

