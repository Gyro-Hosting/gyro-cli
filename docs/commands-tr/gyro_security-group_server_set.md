## gyro security-group server set

Bağlı grupları tam liste olarak ayarla

### Synopsis

Bağlı grupları tam liste olarak ayarla

Tek değişiklikte bağlar / ayırır (sıra = değerlendirme sırası, en çok 5). Sunucunun önceki değişikliği uygulanırken 422 err_busy döner: birkaç saniye sonra tekrar deneyin.

API: PUT /cloud/servers/{id}/security-groups (operationId setServerSecurityGroups, scope cloud_write)

```
gyro security-group server set <ID> [flags]
```

### Options

```
      --data string               istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                   hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                      help for set
      --idempotency-key string    tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --security-groups strings   security_groups
      --set stringArray           ek gövde alanı anahtar=değer (tekrarlanabilir)
      --wait                      sunucu meşgul olmayana kadar bekle
      --wait-timeout duration     --wait için en uzun süre (default 15m0s)
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

* [gyro security-group server](gyro_security-group_server.md)	 - Sunucuya bağlı güvenlik grupları

