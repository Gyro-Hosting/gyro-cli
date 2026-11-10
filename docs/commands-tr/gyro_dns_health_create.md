## gyro dns health create

Sağlık kontrollü kayıt oluştur

### Synopsis

Sağlık kontrollü kayıt oluştur

`name` kaydını `primary`'ye yönlendirir; kontrol çoğunluk lokasyondan düşerse `backup`'a geçer, 3 başarılı turdan sonra geri döner. Bir sağlık kontrolü ek hizmet birimi kullanır.

API: POST /cloud/dns/{id}/health-checks (operationId typed_post_dns__id__health_checks, scope services_write)

```
gyro dns health create <ID> [flags]
```

### Options

```
      --backup string            backup
      --check string             check (http|https|tcp|ping)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              name
      --path string              path
      --port int                 port
      --primary string           primary
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --ttl int                  30-300 s
      --type string              type (A|AAAA)
      --zone string              zone
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

* [gyro dns health](gyro_dns_health.md)	 - Sağlık kontrollü kayıtları listele

