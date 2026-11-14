## gyro container nodes set

Ek düğüm sayısını ayarla

### Synopsis

Ek düğüm sayısını ayarla

Paketin düğüm boyutunda ek düğüm ekler ya da çıkarır (mutlak `extra_nodes`, 0 = yalnız paket; ortam başına toplam en çok 20 düğüm, GyroKota düşürebilir). Ek düğüm paketteki düğümle aynı fiyattır, ön ödemeli bakiyeden saatlik düşer. Çıkarma önce boşaltır: kopyalar kalan düğümlere taşınır (servisler sırayla, sağlık kontrollü yeniden başlar), sonra düğüm silinir; boşaltılırken ücretlenmez. Kalan düğümler çalışan servislere yetmiyorsa 409 `quota_in_use`. Sunucusuz ortamlarda düğüm yoktur. `"confirm": true` gerekir. Ayrılmış ortamın düğüm takımı her değiştiğinde (paket değişikliği, ek düğüm ekleme / çıkarma) değişiklik eski düğümler kapanmadan ÖNCE duyurulur: yeni düğümlerin adresleri belli olur olmaz (genellikle bir dakikadan kısa) ortam sahibine bildirim gider (Uyarılar kanalları, `egress` olayı), `egress_change` / `pending_egress_ips` alanlarında görünür. Onay adımı açıksa (`PUT /cloud/containers/{id}/egress-approval`, 1-72 saat, önerilen 24) yeni düğümler trafik almaz, eski düğümler çalışmaya devam eder; `POST /cloud/containers/{id}/egress-change/approve` ya da süre dolunca geçiş devam eder. Yalnız düğüm çıkarma onay beklemez.

API: PUT /cloud/containers/{id}/nodes (operationId setContainerNodes, scope services_write)

```
gyro container nodes set <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --extra-nodes int          extra_nodes
  -h, --help                     help for set
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

* [gyro container nodes](gyro_container_nodes.md)	 - Düğüm havuzu

