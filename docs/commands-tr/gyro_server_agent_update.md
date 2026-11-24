## gyro server agent update

Misafir ajanını güncelle

### Synopsis

Misafir ajanını güncelle

Eski ajanın sabit Gyro indirme kanalıyla otomatik güncellemeyi başlatır (ajan 3 ya da önce sorgulanan bilinmeyen sürüm); sunucu çalışmaya devam eder, yaklaşık 30 sn, ardından özel ağ ayarı yeniden uygulanır. Ret (409): `manual_update_required` (ajan 1-2: `update_command` çalıştırın), `already_latest`, `not_running`, `not_supported`, `busy`. Sunucu internete çıkamıyorsa görev başarısız olur; o zaman `update_command`'ı web konsoldan çalıştırın. Gyro misafir ajanı özel ağ kartlarını ve public IP kapalıyken rotayı sunucu içinde yapılandırır. 15 Kasım 2026 öncesi kurulan sunucularda eski ajan vardır: kart takılır ama `applied` false kalır ve public IP kapatma 409 `agent_update_required` döner. `update_command`, sunucuda root olarak çalıştırılacak tek satır komuttur (Gyro'nun sürümlü güncelleyicisini indirir ve SHA-256 özetini doğrular); yalnız ajanı değiştirir, yeniden başlatma gerekmez, verilere dokunmaz. Karta adresi kendiniz verdiyseniz (netplan) sonra o dosyayı silip `netplan apply` çalıştırın.

API: POST /cloud/servers/{id}/agent/update (operationId updateServerAgent, scope cloud_write)

```
gyro server agent update <ID> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for update
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --wait                     sunucu meşgul olmayana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
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

* [gyro server agent](gyro_server_agent.md)	 - Misafir ajanı sürümü

