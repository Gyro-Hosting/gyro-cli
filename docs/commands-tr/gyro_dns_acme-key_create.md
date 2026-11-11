## gyro dns acme-key create

ACME anahtarı oluştur

### Synopsis

ACME anahtarı oluştur

Sertifika otomasyonu için yalnız `_acme-challenge` TXT kaydı ekleyip silebilen anahtar (kapsam `dns_acme`); isteğe bağlı yalnız `zones` içinde ve yalnız `ips` adreslerinden. Zone okuyamaz, başka kayıt değiştiremez, başka uca erişemez. Anahtar bir kez döner. Hizmet başına en çok 10.

API: POST /cloud/dns/{id}/acme-keys (operationId createAcmeKey, scope services_write)

```
gyro dns acme-key create <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --ips strings              Allowed client IPs (empty = any)
      --name string              Label (e.g. web-01)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --zones strings            Allowed zones (empty = all zones of the service)
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

* [gyro dns acme-key](gyro_dns_acme-key.md)	 - ACME anahtarlarını listele

