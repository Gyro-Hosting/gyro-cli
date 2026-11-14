## gyro vpc create

Özel ağ oluştur

### Synopsis

Özel ağ oluştur

Yalnız tr-ist-1. Sınırlar (GET /cloud/vpcs `limits`, destek talebiyle artırılır): hesap başına 3 VPC, VPC başına 4 subnet, sunucu başına 3 özel kart, VPC başına 20 port yönlendirme; bir sunucu aynı anda tek VPC'ye (o VPC'nin birden çok subnetine) bağlanır. Başka hesabın nesnesi 404 döner. 10.0.0.0/8, 172.16.0.0/12 ya da 192.168.0.0/16 içinde /16 - /24 özel aralık; ağ adresi yazılır (10.10.0.0/16, 10.10.0.5/16 değil). Platforma ayrılmış: 10.42.0.0/16, 10.43.0.0/16, 10.200.0.0/16, 10.250.0.0/16, 10.255.0.0/16. İlk subnet birlikte açılır: `subnet.cidr` ya da verilmezse aralığın ilk /24'ü (ad `default`). Ücretsiz.

API: POST /cloud/vpcs (operationId createVpc, scope cloud_write)

```
gyro vpc create [flags]
```

### Options

```
      --cidr string              cidr
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              name
      --project-id int           project_id
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --subnet string            subnet (JSON)
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

* [gyro vpc](gyro_vpc.md)	 - Özel ağları listele

