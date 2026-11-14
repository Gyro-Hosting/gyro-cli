## gyro container service autoscale

Servisin otomatik ölçekleme / uyarı eşiği

### Synopsis

Servisin otomatik ölçekleme / uyarı eşiği

Gövde `{"autoscaling": {...}, "alerts": {"min_ready": 2}}`. Ortalama kullanım hedefi aşınca kopya eklenir (30 sn'de en çok iki katı), 3 dakika sakin kalınca dakikada bir azaltılır. En çok kopya paket sınırını aşamaz; havuza sığmayan kopyalar bekler. Sürüm oluşturmaz. Kapatınca mevcut kopya sayısı (ya da `replicas`) kalır. Sayı servis başına 0..en fazla kopya aralığı dışındaysa 422 `invalid_replicas` (`details.max` = planınızın sınırı).

API: PUT /cloud/containers/{id}/services/{name}/autoscaling (operationId typed_put_containers__id__services__name__autoscaling, scope services_write)

```
gyro container service autoscale <ID> <NAME> [flags]
```

### Options

```
      --alerts string               Uyarı eşiği: hazır kopya bu sayının altına düşerse bildirim (0 = kapalı). (JSON)
      --autoscaling enabled:false   Otomatik ölçekleme: kopya sayısı ayrılan boyutun ortalama CPU (isteğe bağlı bellek) kullanımına göre değişir. enabled:false = kapalı. (JSON)
      --data string                 istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                     hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                        help for autoscale
      --idempotency-key string      tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --replicas int                Fixed replicas when switching autoscaling off
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

