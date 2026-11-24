## gyro vpc delete

Özel ağı sil

### Synopsis

Özel ağı sil

Kalıcıdır. Önce sunucu kartlarını çıkarın ve port yönlendirmelerini silin (409 `vpc_in_use`); açık kapı kapatılır. `"confirm": true` gerekir (gövde ya da sorgu).

API: DELETE /cloud/vpcs/{vpc_id} (operationId deleteVpc, scope cloud_write)

```
gyro vpc delete <VPC-ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for delete
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro vpc](gyro_vpc.md)	 - Özel ağları listele

