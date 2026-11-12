## gyro container service update

Yeni sürüm dağıt / ayar değiştir

### Synopsis

Yeni sürüm dağıt / ayar değiştir

Yalnız değişecek alanları gönderin (ör. `image`). `replicas` dışındaki her değişiklik yeni sürüm oluşturur ve kesintisiz dağıtılır.

API: PATCH /cloud/containers/{id}/services/{name} (operationId typed_patch_containers__id__services__name, scope services_write)

```
gyro container service update <ID> <NAME> [flags]
```

### Options

```
      --alerts string               Uyarı eşiği: hazır kopya bu sayının altına düşerse bildirim (0 = kapalı). (JSON)
      --args strings                İmajın komutuna argümanlar
      --autoscaling enabled:false   Otomatik ölçekleme: kopya sayısı ayrılan boyutun ortalama CPU (isteğe bağlı bellek) kullanımına göre değişir. enabled:false = kapalı. (JSON)
      --command strings             İmajın giriş komutunu değiştir
      --cpu-m int                   Örnek başına işlemci (milicore, 250 = 0,25 vCPU)
      --data string                 istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                     hiçbir şeyi değiştirmeden teklifi / isteği göster
      --env string                  Ortam değişkenleri (güncellemede tamamı değişir) (JSON)
      --health-path string          HTTP sağlık yolu (boş = TCP)
  -h, --help                        help for update
      --idempotency-key string      tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --image string                İmaj, ör. ghcr.io/firma/web:1.4
      --memory-mi int               Örnek başına bellek (MiB)
      --port int                    Kapsayıcı portu (0 = port yok). Kök olmayan kullanıcı: 1024 üstü port
      --public                      HTTPS adres <ad>-<kod>.gyroapp.net
      --reason string               Free text shown in the revision history (e.g. a git sha)
      --replicas int                Örnek sayısı (varsayılan 1)
      --run-as-user int             Kapsayıcı içi kullanıcı kimliği (varsayılan 10001)
      --set stringArray             ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro container service](gyro_container_service.md)	 - Kapsayıcı servislerini listele

