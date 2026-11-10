## gyro security-group detach

Güvenlik grubunu sunucudan ayır

### Synopsis

Güvenlik grubunu sunucudan ayır

API: DELETE /cloud/servers/{id}/security-groups/{sg} (operationId detachSecurityGroup, scope cloud_write)

```
gyro security-group detach <ID> <SG> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for detach
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --wait                     sunucu meşgul olmayana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
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

* [gyro security-group](gyro_security-group.md)	 - Güvenlik gruplarını listele

