## gyro security-group rules set

Güvenlik grubunun kurallarını değiştir

### Synopsis

Güvenlik grubunun kurallarını değiştir

Tam değiştirme (gelen + giden), en çok 60 kural. Bağlı tüm sunuculara uygulanır. Bağlı bir sunucuda herhangi bir kaynağın SSH (22) / RDP (3389) erişimini kaldırıyor ya da daraltıyorsa veya çağıran adresi kesiyorsa 5 dakikalık deneme olur (`trial: true`, `trial_seconds_left`) - panel, CLI ve Terraform ile aynı kural: .../confirm ile onaylayın, yoksa önceki kurallar döner. Sunucunun kuralları (kendi + bağlı tüm gruplar) hakkına sığmalı (422 err_sg_quota).

API: PUT /cloud/security-groups/{sg}/rules (operationId putSecurityGroupRules, scope cloud_write)

```
gyro security-group rules set <SG> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --rules strings            rules
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

* [gyro security-group rules](gyro_security-group_rules.md)	 - Güvenlik grubunun kurallarını değiştir

