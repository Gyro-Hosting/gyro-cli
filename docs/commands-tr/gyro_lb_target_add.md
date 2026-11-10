## gyro lb target add

Hedef ekle

### Synopsis

Hedef ekle

type=server (server = bulut sunucu hizmet no, aynı bölge), type=all (tüm bulut sunucularınız, yeniler otomatik) ya da type=ip (fiziksel sunucu IP'niz). use_private_ip için özel ağ bağlı olmalı.

API: POST /cloud/load-balancers/{id}/targets (operationId typed_post_load_balancers__id__targets, scope services_write)

```
gyro lb target add <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for add
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --ip string                Dedicated server IP (type=ip)
      --server int               Cloud server id (type=server)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --type string              type (server|all|ip)
      --use-private-ip           use_private_ip
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

* [gyro lb target](gyro_lb_target.md)	 - Hedefleri sağlık durumuyla listele

