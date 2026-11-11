## gyro container domain add

Alan adı bağla

### Synopsis

Alan adı bağla

Önce CNAME kaydını servis adresine (`<ad>-<kod>.gyroapp.net`) yönlendirin; DNS bize gösterince ücretsiz sertifikayla etkinleşir (saatlik ya da .../domains/verify ile).

API: POST /cloud/containers/{id}/services/{name}/domains (operationId typed_post_containers__id__services__name__domains, scope services_write)

```
gyro container domain add <ID> <NAME> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --domain string            domain
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for add
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

* [gyro container domain](gyro_container_domain.md)	 - Servisin alan adları

