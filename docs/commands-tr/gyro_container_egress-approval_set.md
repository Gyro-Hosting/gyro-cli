## gyro container egress-approval set

Çıkış IP değişikliği onay adımı

### Synopsis

Çıkış IP değişikliği onay adımı

`hours` 1-72 onay adımını açar (önerilen 24), 0 kapatır. Ayrılmış ortamın düğüm takımı her değiştiğinde (paket değişikliği, ek düğüm ekleme / çıkarma) değişiklik eski düğümler kapanmadan ÖNCE duyurulur: yeni düğümlerin adresleri belli olur olmaz (genellikle bir dakikadan kısa) ortam sahibine bildirim gider (Uyarılar kanalları, `egress` olayı), `egress_change` / `pending_egress_ips` alanlarında görünür. Onay adımı açıksa (`PUT /cloud/containers/{id}/egress-approval`, 1-72 saat, önerilen 24) yeni düğümler trafik almaz, eski düğümler çalışmaya devam eder; `POST /cloud/containers/{id}/egress-change/approve` ya da süre dolunca geçiş devam eder. Yalnız düğüm çıkarma onay beklemez.

API: PUT /cloud/containers/{id}/egress-approval (operationId setContainerEgressApproval, scope services_write)

```
gyro container egress-approval set <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --hours int                hours
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

* [gyro container egress-approval](gyro_container_egress-approval.md)	 - Çıkış IP değişikliği onay adımı

