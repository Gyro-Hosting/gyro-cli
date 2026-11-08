## gyro server snapshot

Snapshot'ları ve kotayı listele

```
gyro server snapshot [flags]
```

### Options

```
  -h, --help   help for snapshot
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

* [gyro server](gyro_server.md)	 - Bulut sunucular: oluşturma, güç, snapshot, SSH ...
* [gyro server snapshot create](gyro_server_snapshot_create.md)	 - Snapshot al
* [gyro server snapshot delete](gyro_server_snapshot_delete.md)	 - Snapshot sil
* [gyro server snapshot list](gyro_server_snapshot_list.md)	 - Snapshot'ları ve kotayı listele
* [gyro server snapshot restore](gyro_server_snapshot_restore.md)	 - Snapshot'ı sunucuya geri yükle (diski ezer)

