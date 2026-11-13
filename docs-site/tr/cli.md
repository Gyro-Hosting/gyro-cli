# Gyro CLI — komut satırı aracı

gyro: Gyro Hosting hizmetlerini terminalden yönetin.

## Kurulum

(HTML: cli.html#install)

## Başlarken

(HTML: cli.html#start)

## Örnek oturum

(HTML: cli.html#session)

## Profiller ve kimlik

(HTML: cli.html#profiles)

## Çıktı ve betikler

(HTML: cli.html#output)

## Güvenlik: onay, deneme, idempotency

(HTML: cli.html#safety)

## Tüm ürünler

(HTML: cli.html#lines)

## Komut başvurusu

| Komut | Açıklama | API |
|---|---|---|
| `gyro activity list` | API anahtarlarınızla yapılan yazma çağrılarını listele | `GET /cloud/activity` |
| `gyro address create` | Adres Ekleme | `POST /addresses` |
| `gyro address default` | Varsayılanı Taşıma | `POST /addresses/{id}/default` |
| `gyro address delete` | Adresi Silme | `DELETE /addresses/{id}` |
| `gyro address get` | Tek Adresi Okuma | `GET /addresses/{id}` |
| `gyro address list` | Adresleri Listeleme | `GET /addresses` |
| `gyro address set` | Adresi Güncelleme | `PUT /addresses/{id}` |
| `gyro affiliate commission list` | Komisyonları Listeleme | `GET /affiliate/commissions` |
| `gyro affiliate enroll` | Programa Katılma | `POST /affiliate/enroll` |
| `gyro affiliate get` | Program Durumunu Okuma | `GET /affiliate` |
| `gyro affiliate payment-method set` | Ödeme Hedefini Kaydetme | `PUT /affiliate/payment-method` |
| `gyro affiliate withdraw` | Ödeme Talebi Açma | `POST /affiliate/withdraw` |
| `gyro api request DELETE /cloud/mail/{id}/aliases/{aid}` | aliases:delete (GyroMail) | `DELETE /cloud/mail/{id}/aliases/{aid}` |
| `gyro api request DELETE /cloud/mail/{id}/autoresponders/{local}` | autoresponder:delete (GyroMail) | `DELETE /cloud/mail/{id}/autoresponders/{local}` |
| `gyro api request DELETE /cloud/mail/{id}/mailboxes/{local}` | mailboxes:delete (GyroMail) | `DELETE /cloud/mail/{id}/mailboxes/{local}` |
| `gyro api request DELETE /cloud/mail/{id}/migrations/{jid}` | sync-jobs:delete (GyroMail) | `DELETE /cloud/mail/{id}/migrations/{jid}` |
| `gyro api request DELETE /cloud/storage/{id}/buckets/{bucket}` | buckets:delete (GyroDepolama) | `DELETE /cloud/storage/{id}/buckets/{bucket}` |
| `gyro api request DELETE /cloud/storage/{id}/keys/{key}` | access-keys:delete (GyroDepolama) | `DELETE /cloud/storage/{id}/keys/{key}` |
| `gyro api request GET /cloud/apps` | Kurulabilir uygulamaları listele | `GET /cloud/apps` |
| `gyro api request GET /cloud/billing/balance` | Balance, hourly cost of running prepaid services and estimated time left | `GET /cloud/billing/balance` |
| `gyro api request GET /cloud/billing/statement/{month}` | Monthly usage statement (not an invoice); JSON, PDF or CSV | `GET /cloud/billing/statement/{month}` |
| `gyro api request GET /cloud/billing/usage` | Hour-by-hour usage ledger (or daily/monthly totals). Times are Europe/Istanbul. | `GET /cloud/billing/usage` |
| `gyro api request GET /cloud/mail/{id}/aliases` | aliases (GyroMail) | `GET /cloud/mail/{id}/aliases` |
| `gyro api request GET /cloud/mail/{id}/autoresponders` | autoresponder (GyroMail) | `GET /cloud/mail/{id}/autoresponders` |
| `gyro api request GET /cloud/mail/{id}/catch-all` | catch-all (GyroMail) | `GET /cloud/mail/{id}/catch-all` |
| `gyro api request GET /cloud/mail/{id}/dns` | dns-setup (GyroMail) | `GET /cloud/mail/{id}/dns` |
| `gyro api request GET /cloud/mail/{id}/mailboxes` | mailboxes (GyroMail) | `GET /cloud/mail/{id}/mailboxes` |
| `gyro api request GET /cloud/mail/{id}/migrations` | sync-jobs (GyroMail) | `GET /cloud/mail/{id}/migrations` |
| `gyro api request GET /cloud/storage/{id}/buckets` | buckets (GyroDepolama) | `GET /cloud/storage/{id}/buckets` |
| `gyro api request GET /cloud/storage/{id}/keys` | access-keys (GyroDepolama) | `GET /cloud/storage/{id}/keys` |
| `gyro api request PATCH /cloud/mail/{id}/aliases/{aid}` | aliases:edit (GyroMail) | `PATCH /cloud/mail/{id}/aliases/{aid}` |
| `gyro api request PATCH /cloud/mail/{id}/mailboxes/{local}` | mailboxes:edit (GyroMail) | `PATCH /cloud/mail/{id}/mailboxes/{local}` |
| `gyro api request POST /cloud/billing/estimate` | Hourly / daily / monthly cost of prepaid plans and serverless containers, and how long the balance lasts | `POST /cloud/billing/estimate` |
| `gyro api request POST /cloud/licenses/{id}/change-ip` | method change_ip (LisansAPP) | `POST /cloud/licenses/{id}/change-ip` |
| `gyro api request POST /cloud/mail/{id}/aliases` | aliases:create (GyroMail) | `POST /cloud/mail/{id}/aliases` |
| `gyro api request POST /cloud/mail/{id}/mailboxes` | mailboxes:create (GyroMail) | `POST /cloud/mail/{id}/mailboxes` |
| `gyro api request POST /cloud/mail/{id}/migrations` | sync-jobs:create (GyroMail) | `POST /cloud/mail/{id}/migrations` |
| `gyro api request POST /cloud/storage/{id}/buckets` | Kova oluştur | `POST /cloud/storage/{id}/buckets` |
| `gyro api request POST /cloud/storage/{id}/keys` | access-keys:create (GyroDepolama) | `POST /cloud/storage/{id}/keys` |
| `gyro balance` | Hesap bakiyesi | `GET /balance` |
| `gyro balance setting update` | Cüzdan Ayarlarını Değiştirme | `PATCH /balance/settings` |
| `gyro balance top-up` | Bakiye yükleme faturası oluştur, ödeme linkini al | `POST /cloud/balance/top-up` |
| `gyro balance transaction list` | Bakiye hareketleri | `GET /balance/transactions` |
| `gyro billing topup-url get` | Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not) | `GET /cloud/billing/topup-url` |
| `gyro bulk` | Bir eylemi 50 sunucuya kadar toplu çalıştır | `POST /cloud/bulk` |
| `gyro container create` | Kapsayıcı ortamı oluştur (toplu) | `POST /cloud/containers` |
| `gyro container domain add` | Alan adı bağla | `POST /cloud/containers/{id}/services/{name}/domains` |
| `gyro container domain list` | Servisin alan adları | `GET /cloud/containers/{id}/services/{name}/domains` |
| `gyro container domain remove` | Alan adını kaldır | `DELETE /cloud/containers/{id}/services/{name}/domains/{domain}` |
| `gyro container domain verify` | Bekleyen alan adlarını kontrol et | `POST /cloud/containers/{id}/services/{name}/domains/verify` |
| `gyro container env set` | Ortam değişkenlerini ayarla | `PUT /cloud/containers/{id}/services/{name}/env` |
| `gyro container events` | Son uyarılar | `GET /cloud/containers/{id}/events` |
| `gyro container external-registry delete` | Özel depo bilgisini sil | `DELETE /cloud/containers/{id}/external-registry` |
| `gyro container external-registry set` | Harici özel depo bilgisi | `PUT /cloud/containers/{id}/external-registry` |
| `gyro container logs` | Servis logları | `GET /cloud/containers/{id}/services/{name}/logs` |
| `gyro container metrics` | Servis işlemci / bellek | `GET /cloud/containers/{id}/services/{name}/metrics` |
| `gyro container nodes` | Düğüm havuzu | `GET /cloud/containers/{id}/nodes` |
| `gyro container notifications get` | Bildirim ayarları ve süren uyarılar | `GET /cloud/containers/{id}/notifications` |
| `gyro container notifications set` | Bildirim kanallarını ayarla | `PUT /cloud/containers/{id}/notifications` |
| `gyro container notifications test` | Test bildirimi gönder | `POST /cloud/containers/{id}/notifications/test` |
| `gyro container notifications webhook-secret` | Webhook imza anahtarı | `POST /cloud/containers/{id}/notifications/webhook-secret` |
| `gyro container pipeline approve` | Dağıtımı onayla | `POST /cloud/containers/{id}/pipeline-runs/{run}/approve` |
| `gyro container pipeline cancel` | Çalıştırmayı durdur | `POST /cloud/containers/{id}/pipeline-runs/{run}/cancel` |
| `gyro container pipeline create` | Pipeline oluştur | `POST /cloud/containers/{id}/pipelines` |
| `gyro container pipeline delete` | Pipeline'ı sil | `DELETE /cloud/containers/{id}/pipelines/{pipeline}` |
| `gyro container pipeline deploy-key create` | SSH deploy key oluştur | `PUT /cloud/containers/{id}/pipelines/{pipeline}/deploy-key` |
| `gyro container pipeline deploy-key delete` | SSH deploy key'i sil | `DELETE /cloud/containers/{id}/pipelines/{pipeline}/deploy-key` |
| `gyro container pipeline get` | Pipeline'ı getir | `GET /cloud/containers/{id}/pipelines/{pipeline}` |
| `gyro container pipeline host-key scan` | SSH sunucu anahtarı | `GET /cloud/containers/{id}/pipelines/{pipeline}/host-key` |
| `gyro container pipeline host-key trust` | SSH sunucu anahtarını onayla (ya da kaldır) | `PUT /cloud/containers/{id}/pipelines/{pipeline}/host-key` |
| `gyro container pipeline list` | Pipeline'ları listele | `GET /cloud/containers/{id}/pipelines` |
| `gyro container pipeline logs` | Çalıştırma logu | `GET /cloud/containers/{id}/pipeline-runs/{run}/logs` |
| `gyro container pipeline redeploy` | Bu çalıştırmanın imajını yeniden dağıt (geri dönüş) | `POST /cloud/containers/{id}/pipeline-runs/{run}/redeploy` |
| `gyro container pipeline reject` | Dağıtımı reddet | `POST /cloud/containers/{id}/pipeline-runs/{run}/reject` |
| `gyro container pipeline rotate-secret` | Webhook gizli anahtarını yenile | `POST /cloud/containers/{id}/pipelines/{pipeline}/webhook/rotate` |
| `gyro container pipeline run` | Çalıştırma başlat | `POST /cloud/containers/{id}/pipelines/{pipeline}/runs` |
| `gyro container pipeline run-get` | Çalıştırmayı getir | `GET /cloud/containers/{id}/pipeline-runs/{run}` |
| `gyro container pipeline runs` | Pipeline çalıştırmaları | `GET /cloud/containers/{id}/pipelines/{pipeline}/runs` |
| `gyro container pipeline secret delete` | Derleme gizli değerini sil | `DELETE /cloud/containers/{id}/pipelines/{pipeline}/secrets/{secret}` |
| `gyro container pipeline secret set` | Derleme gizli değeri ayarla | `PUT /cloud/containers/{id}/pipelines/{pipeline}/secrets` |
| `gyro container pipeline test` | Depo erişimini test et | `POST /cloud/containers/{id}/pipelines/{pipeline}/test` |
| `gyro container pipeline update` | Pipeline'ı değiştir | `PATCH /cloud/containers/{id}/pipelines/{pipeline}` |
| `gyro container pipeline upload` | Arşiv derlemesi için yükleme adresi | `POST /cloud/containers/{id}/pipelines/{pipeline}/uploads` |
| `gyro container pipeline webhook` | Webhook adresi ve gizli anahtar | `GET /cloud/containers/{id}/pipelines/{pipeline}/webhook` |
| `gyro container plans` | Kapsayıcı ortam paketlerini listele | `GET /cloud/containers/plans` |
| `gyro container registry credentials` | Depo parolası | `POST /cloud/containers/{id}/registry/credentials` |
| `gyro container registry get` | Ortamın Gyro imaj deposu | `GET /cloud/containers/{id}/registry` |
| `gyro container registry image-delete` | İmaj etiketini sil | `DELETE /cloud/containers/{id}/registry/images` |
| `gyro container registry images` | Depodaki imajlar | `GET /cloud/containers/{id}/registry/images` |
| `gyro container registry import` | Başka bir depodan imaj içe aktar | `POST /cloud/containers/{id}/registry/imports` |
| `gyro container registry imports` | İmaj içe aktarmaları (ilerleme ve geçmiş) | `GET /cloud/containers/{id}/registry/imports` |
| `gyro container registry policy` | İmaj temizlik ayarı | `PUT /cloud/containers/{id}/registry/policy` |
| `gyro container registry rotate` | Depo parolasını yenile | `POST /cloud/containers/{id}/registry/rotate` |
| `gyro container secret delete` | Gizli değer sil | `DELETE /cloud/containers/{id}/services/{name}/secrets` |
| `gyro container secret set` | Gizli değer ayarla | `PUT /cloud/containers/{id}/services/{name}/secrets` |
| `gyro container service autoscale` | Servisin otomatik ölçekleme / uyarı eşiği | `PUT /cloud/containers/{id}/services/{name}/autoscaling` |
| `gyro container service create` | Kapsayıcı servisi oluştur | `POST /cloud/containers/{id}/services` |
| `gyro container service delete` | Kapsayıcı servisini sil | `DELETE /cloud/containers/{id}/services/{name}` |
| `gyro container service get` | Kapsayıcı servisini getir | `GET /cloud/containers/{id}/services/{name}` |
| `gyro container service list` | Kapsayıcı servislerini listele | `GET /cloud/containers/{id}/services` |
| `gyro container service restart` | Servisi yeniden başlat | `POST /cloud/containers/{id}/services/{name}/restart` |
| `gyro container service revisions` | Sürüm geçmişi | `GET /cloud/containers/{id}/services/{name}/revisions` |
| `gyro container service rollback` | Önceki sürüme dön | `POST /cloud/containers/{id}/services/{name}/rollback` |
| `gyro container service scale` | Servisi ölçekle | `POST /cloud/containers/{id}/services/{name}/scale` |
| `gyro container service update` | Yeni sürüm dağıt / ayar değiştir | `PATCH /cloud/containers/{id}/services/{name}` |
| `gyro container sizes` | Sunucusuz kapsayıcı boyutları (vCPU / bellek) ve fiyatları | `GET /cloud/billing/container-sizes` |
| `gyro core invoice pay` | Faturayı Ödeme | `POST /invoices/{id}/pay` |
| `gyro core product list` | Ürünleri Listeleme | `GET /products` |
| `gyro core service addon list` | Hizmet Eklentilerini Okuma | `GET /services/{id}/addons` |
| `gyro core service get` | Hizmeti Okuma | `GET /services/{id}` |
| `gyro core service list` | Hizmetleri Listeleme | `GET /services` |
| `gyro core service tool get` | Araç Verisini Okuma | `GET /services/{id}/tools/{tool}` |
| `gyro core service tool list` | Araçları Listeleme | `GET /services/{id}/tools` |
| `gyro core service tool run` | Araç Eylemi Çalıştırma | `POST /services/{id}/tools/{tool}/{action}` |
| `gyro dns acme-challenge add` | ACME doğrulama TXT değeri ekle | `POST /cloud/dns/acme-challenge` |
| `gyro dns acme-challenge add-in-service` | ACME doğrulama TXT değeri ekle | `POST /cloud/dns/{id}/acme-challenge` |
| `gyro dns acme-challenge remove` | ACME doğrulama TXT değerini sil | `DELETE /cloud/dns/acme-challenge` |
| `gyro dns acme-challenge remove-in-service` | ACME doğrulama TXT değerini sil | `DELETE /cloud/dns/{id}/acme-challenge` |
| `gyro dns acme-challenge status` | Doğrulama kaydı yayında mı? | `GET /cloud/dns/acme-challenge` |
| `gyro dns acme-challenge status-in-service` | Doğrulama kaydı yayında mı? | `GET /cloud/dns/{id}/acme-challenge` |
| `gyro dns acme-key create` | ACME anahtarı oluştur | `POST /cloud/dns/{id}/acme-keys` |
| `gyro dns acme-key delete` | ACME anahtarını sil | `DELETE /cloud/dns/{id}/acme-keys/{kid}` |
| `gyro dns acme-key list` | ACME anahtarlarını listele | `GET /cloud/dns/{id}/acme-keys` |
| `gyro dns dnssec disable` | DNSSEC kapat | `DELETE /cloud/dns/{id}/zones/{zone}/dnssec` |
| `gyro dns dnssec enable` | DNSSEC aç | `POST /cloud/dns/{id}/zones/{zone}/dnssec` |
| `gyro dns dnssec status` | DNSSEC durumu ve DS kayıtları | `GET /cloud/dns/{id}/zones/{zone}/dnssec` |
| `gyro dns health create` | Sağlık kontrollü kayıt oluştur | `POST /cloud/dns/{id}/health-checks` |
| `gyro dns health delete` | Sağlık kontrolünü sil | `DELETE /cloud/dns/{id}/health-checks/{hid}` |
| `gyro dns health list` | Sağlık kontrollü kayıtları listele | `GET /cloud/dns/{id}/health-checks` |
| `gyro dns rrset delete` | Kayıt kümesini sil | `DELETE /cloud/dns/{id}/zones/{zone}/records` |
| `gyro dns rrset list` | Kayıt kümelerini listele | `GET /cloud/dns/{id}/zones/{zone}/records` |
| `gyro dns rrset set` | Kayıt kümesi oluştur ya da değiştir | `PUT /cloud/dns/{id}/zones/{zone}/records` |
| `gyro dns zone create` | Zone ekle | `POST /cloud/dns/{id}/zones` |
| `gyro dns zone delete` | Zone'u sil | `DELETE /cloud/dns/{id}/zones/{zone}` |
| `gyro dns zone export` | Zone dosyasını dışa aktar | `GET /cloud/dns/{id}/zones/{zone}/export` |
| `gyro dns zone get` | Zone'u kayıtlarıyla getir | `GET /cloud/dns/{id}/zones/{zone}` |
| `gyro dns zone import` | BIND zone dosyası içe aktar | `POST /cloud/dns/{id}/zones/{zone}/import` |
| `gyro dns zone list` | DNS zone'larını listele | `GET /cloud/dns/{id}/zones` |
| `gyro dns zone verify` | Sahipliği şimdi doğrula | `POST /cloud/dns/{id}/zones/{zone}/verify` |
| `gyro domain addon buy` | Alan Adı Ek Hizmeti Alma | `POST /domains/{domain}/addons/{key}` |
| `gyro domain auth-code send` | Transfer Kodunu İsteme | `POST /domains/{domain}/auth-code` |
| `gyro domain auth-code set` | Gelen Transfer Kodunu Kaydetme | `PUT /domains/{domain}/auth-code` |
| `gyro domain check` | Alan Adının Boş Olup Olmadığını Sorma | `POST /domains/check` |
| `gyro domain child-nameserver create` | Kendi Ad Sunucusunu Tanımlama | `POST /domains/{domain}/child-nameservers` |
| `gyro domain child-nameserver delete` | Kendi Ad Sunucusunu Kaldırma | `DELETE /domains/{domain}/child-nameservers` |
| `gyro domain child-nameserver list` | Kendi Ad Sunucularını Listeleme | `GET /domains/{domain}/child-nameservers` |
| `gyro domain default-nameserver list` | Varsayılan Ad Sunucularını Okuma | `GET /domains/default-nameservers` |
| `gyro domain default-nameserver set` | Varsayılan Ad Sunucularını Kaydetme | `PUT /domains/default-nameservers` |
| `gyro domain dns-record create` | DNS Kaydı Ekleme | `POST /domains/{domain}/dns-records` |
| `gyro domain dns-record delete` | DNS Kaydını Silme | `DELETE /domains/{domain}/dns-records` |
| `gyro domain dns-record list` | DNS Kayıtlarını Listeleme | `GET /domains/{domain}/dns-records` |
| `gyro domain dns-record set` | DNS Kaydını Düzenleme | `PUT /domains/{domain}/dns-records` |
| `gyro domain dnssec create` | İmza Kaydı Ekleme | `POST /domains/{domain}/dnssec` |
| `gyro domain dnssec delete` | İmza Kaydını Silme | `DELETE /domains/{domain}/dnssec` |
| `gyro domain dnssec list` | İmza Kayıtlarını Listeleme | `GET /domains/{domain}/dnssec` |
| `gyro domain email-forward create` | E-posta Yönlendirmesi Ekleme | `POST /domains/{domain}/email-forwards` |
| `gyro domain email-forward delete` | E-posta Yönlendirmesini Silme | `DELETE /domains/{domain}/email-forwards` |
| `gyro domain email-forward list` | E-posta Yönlendirmelerini Listeleme | `GET /domains/{domain}/email-forwards` |
| `gyro domain forwarding delete` | Adres Yönlendirmesini Kaldırma | `DELETE /domains/{domain}/forwarding` |
| `gyro domain forwarding get` | Adres Yönlendirmesini Okuma | `GET /domains/{domain}/forwarding` |
| `gyro domain forwarding set` | Adres Yönlendirmesi Kurma | `PUT /domains/{domain}/forwarding` |
| `gyro domain get` | Alan Adını Okuma | `GET /domains/{domain}` |
| `gyro domain list` | Alan Adlarını Listeleme | `GET /domains` |
| `gyro domain nameserver list` | Ad Sunucularını Okuma | `GET /domains/{domain}/nameservers` |
| `gyro domain nameserver set` | Ad Sunucularını Değiştirme | `PUT /domains/{domain}/nameservers` |
| `gyro domain register` | Alan Adı Kaydetme | `POST /domains/register` |
| `gyro domain renew` | Yenileme Faturası Kesme | `POST /domains/{domain}/renew` |
| `gyro domain tld list` | Uzantı Fiyatlarını Okuma | `GET /domains/tlds` |
| `gyro domain transfer` | Alan Adı Transferi | `POST /domains/transfer` |
| `gyro domain transfer-lock get` | Transfer Kilidini Okuma | `GET /domains/{domain}/transfer-lock` |
| `gyro domain transfer-lock set` | Transfer Kilidini Değiştirme | `PUT /domains/{domain}/transfer-lock` |
| `gyro domain transfer-status get` | Transferin Nerede Olduğunu Sorma | `GET /domains/{domain}/transfer-status` |
| `gyro domain update` | Alan Adı Tercihlerini Değiştirme | `PATCH /domains/{domain}` |
| `gyro domain whois apply` | Kayıtlı Profili Uygulama | `POST /domains/{domain}/whois/apply` |
| `gyro domain whois get` | Alan Adı Kişilerini Okuma | `GET /domains/{domain}/whois` |
| `gyro domain whois set` | Alan Adı Kişilerini Yazma | `PUT /domains/{domain}/whois` |
| `gyro domain whois-privacy set` | Kayıt Gizliliğini Değiştirme | `PUT /domains/{domain}/whois-privacy` |
| `gyro domain whois-profile create` | Profil Oluşturma | `POST /domains/whois-profiles` |
| `gyro domain whois-profile default` | Profili Varsayılan Yapma | `POST /domains/whois-profiles/{pid}/default` |
| `gyro domain whois-profile delete` | Profili Silme | `DELETE /domains/whois-profiles/{pid}` |
| `gyro domain whois-profile list` | Kayıtlı Profilleri Listeleme | `GET /domains/whois-profiles` |
| `gyro domain whois-profile set` | Profili Güncelleme | `PUT /domains/whois-profiles/{pid}` |
| `gyro firewall confirm` | Güvenlik duvarı denemesini kalıcı yap | `POST /cloud/servers/{id}/firewall/policy/confirm` |
| `gyro firewall delete` | Güvenlik duvarı kurallarınızı kaldır | `DELETE /cloud/servers/{id}/firewall/policy` |
| `gyro firewall get` | Ağ güvenlik duvarı politikası (kurallar, varsayılan kural, durum) | `GET /cloud/servers/{id}/firewall/policy` |
| `gyro firewall set` | Ağ güvenlik duvarı politikasını değiştir | `PUT /cloud/servers/{id}/firewall/policy` |
| `gyro image list` | İşletim sistemi imajlarını listele | `GET /cloud/images` |
| `gyro invoice coupon` | Faturaya Kupon Uygulama | `POST /invoices/{id}/coupon` |
| `gyro invoice get` | Faturayı getir | `GET /invoices/{id}` |
| `gyro invoice list` | Faturaları listele | `GET /invoices` |
| `gyro invoice pay` | Faturayı hesap bakiyesinden öde | `POST /cloud/invoices/{id}/pay` |
| `gyro invoice pdf get` | Fatura PDF (JSON içinde base64) | `GET /cloud/invoices/{id}/pdf` |
| `gyro key create` | Alt hesap anahtarı oluştur (token bir kez döner) | `POST /cloud/keys` |
| `gyro key delete` | Alt hesap anahtarını sil | `DELETE /cloud/keys/{id}` |
| `gyro key list` | Alt hesap anahtarlarını listele (token göstermez) | `GET /cloud/keys` |
| `gyro lb algorithm set` | Dağıtım algoritmasını değiştir | `PUT /cloud/load-balancers/{id}/algorithm` |
| `gyro lb cert cancel` | Bekleyen Let's Encrypt isteğini iptal et | `DELETE /cloud/load-balancers/{id}/certificates/letsencrypt` |
| `gyro lb cert delete` | Sertifikayı sil | `DELETE /cloud/load-balancers/{id}/certificates/{cid}` |
| `gyro lb cert letsencrypt` | Ücretsiz Let's Encrypt sertifikası iste | `POST /cloud/load-balancers/{id}/certificates/letsencrypt` |
| `gyro lb cert list` | Sertifikaları listele | `GET /cloud/load-balancers/{id}/certificates` |
| `gyro lb cert upload` | Sertifika yükle | `POST /cloud/load-balancers/{id}/certificates` |
| `gyro lb cert verify` | Bekleyen Let's Encrypt isteğini sürdür | `POST /cloud/load-balancers/{id}/certificates/letsencrypt/verify` |
| `gyro lb create` | Yük dengeleyici sipariş et | `POST /cloud/load-balancers` |
| `gyro lb delete` | Yük dengeleyiciyi sil (iptal) | `DELETE /cloud/load-balancers/{id}` |
| `gyro lb get` | Yük dengeleyiciyi getir | `GET /cloud/load-balancers/{id}` |
| `gyro lb list` | Yük dengeleyicileri listele | `GET /cloud/load-balancers` |
| `gyro lb metrics` | Yük dengeleyici ölçümleri | `GET /cloud/load-balancers/{id}/metrics` |
| `gyro lb network attach` | Özel ağa bağla | `POST /cloud/load-balancers/{id}/network` |
| `gyro lb network detach` | Özel ağdan ayır | `DELETE /cloud/load-balancers/{id}/network/{nid}` |
| `gyro lb rdns set` | Ters DNS ayarla | `PUT /cloud/load-balancers/{id}/rdns` |
| `gyro lb service add` | Servis ekle | `POST /cloud/load-balancers/{id}/services` |
| `gyro lb service delete` | Servisi sil | `DELETE /cloud/load-balancers/{id}/services/{port}` |
| `gyro lb service list` | Yük dengeleyicinin servislerini (portlarını) listele | `GET /cloud/load-balancers/{id}/services` |
| `gyro lb service update` | Servisi güncelle | `PATCH /cloud/load-balancers/{id}/services/{port}` |
| `gyro lb settings` | Algoritma, ters DNS, özel ağ | `GET /cloud/load-balancers/{id}/settings` |
| `gyro lb target add` | Hedef ekle | `POST /cloud/load-balancers/{id}/targets` |
| `gyro lb target list` | Hedefleri sağlık durumuyla listele | `GET /cloud/load-balancers/{id}/targets` |
| `gyro lb target remove` | Hedefi çıkar | `DELETE /cloud/load-balancers/{id}/targets` |
| `gyro lb types` | Yük dengeleyici boyutları ve fiyatları | `GET /cloud/load-balancer-types` |
| `gyro license get` | method license_status (LisansAPP) | `GET /cloud/licenses/{id}` |
| `gyro location list` | Sipariş verilebilen lokasyonları listele | `GET /cloud/locations` |
| `gyro mail archive get` | archive (GyroMail) | `GET /cloud/mail/{id}/archive` |
| `gyro mail autoresponder set` | autoresponder:edit (GyroMail) | `PUT /cloud/mail/{id}/autoresponders/{local}` |
| `gyro mail catch-all set` | catch-all:edit (GyroMail) | `PUT /cloud/mail/{id}/catch-all` |
| `gyro mail deleted get` | deleted-mail (GyroMail) | `GET /cloud/mail/{id}/deleted` |
| `gyro mail delivery-log get` | delivery-log (GyroMail) | `GET /cloud/mail/{id}/delivery-log` |
| `gyro mail dns verify` | dns-setup:verify (GyroMail) | `POST /cloud/mail/{id}/dns/verify` |
| `gyro mail mailbox password` | mailboxes:password (GyroMail) | `POST /cloud/mail/{id}/mailboxes/{local}/password` |
| `gyro mail migration run` | sync-jobs:run (GyroMail) | `POST /cloud/mail/{id}/migrations/{jid}/run` |
| `gyro mail plan get` | plan-size (GyroMail) | `GET /cloud/mail/{id}/plan` |
| `gyro mail quarantine get` | quarantine (GyroMail) | `GET /cloud/mail/{id}/quarantine` |
| `gyro mail spam-filter get` | spam-filter (GyroMail) | `GET /cloud/mail/{id}/spam-filter` |
| `gyro mail two-factor get` | two-factor (GyroMail) | `GET /cloud/mail/{id}/two-factor` |
| `gyro me get` | Profili Okuma | `GET /me` |
| `gyro me notification list` | Bildirim Tercihlerini Okuma | `GET /me/notifications` |
| `gyro me notification set` | Bildirim Tercihlerini Yazma | `PUT /me/notifications` |
| `gyro me security get` | Güvenlik Özetini Okuma | `GET /me/security` |
| `gyro me update` | Profili Güncelleme | `PATCH /me` |
| `gyro monitor channel create` | channels:create (GyroMonitor) | `POST /cloud/monitors/{id}/channels` |
| `gyro monitor channel delete` | channels:delete (GyroMonitor) | `DELETE /cloud/monitors/{id}/channels/{cid}` |
| `gyro monitor channel list` | channels (GyroMonitor) | `GET /cloud/monitors/{id}/channels` |
| `gyro monitor heartbeat create` | heartbeats:create (GyroMonitor) | `POST /cloud/monitors/{id}/heartbeats` |
| `gyro monitor heartbeat get` | heartbeats (GyroMonitor) | `GET /cloud/monitors/{id}/heartbeats` |
| `gyro monitor monitor create` | monitors:create (GyroMonitor) | `POST /cloud/monitors/{id}/monitors` |
| `gyro monitor monitor delete` | monitors:delete (GyroMonitor) | `DELETE /cloud/monitors/{id}/monitors/{mid}` |
| `gyro monitor monitor list` | monitors (GyroMonitor) | `GET /cloud/monitors/{id}/monitors` |
| `gyro monitor monitor pause` | monitors:pause (GyroMonitor) | `POST /cloud/monitors/{id}/monitors/{mid}/pause` |
| `gyro monitor monitor resume` | monitors:resume (GyroMonitor) | `POST /cloud/monitors/{id}/monitors/{mid}/resume` |
| `gyro monitor monitor update` | monitors:edit (GyroMonitor) | `PATCH /cloud/monitors/{id}/monitors/{mid}` |
| `gyro monitor status-page get` | status-pages (GyroMonitor) | `GET /cloud/monitors/{id}/status-pages` |
| `gyro order create` | Sipariş Verme | `POST /orders` |
| `gyro order get` | Sipariş Detayını Okuma | `GET /orders/{id}` |
| `gyro order list` | Siparişleri Listeleme | `GET /orders` |
| `gyro order preview` | Siparişi Fiyatlandırma | `POST /orders/preview` |
| `gyro ping get` | Ping | `GET /ping` |
| `gyro plan list` | Sipariş verilebilen bulut sunucu ürünlerini listele | `GET /cloud/products` |
| `gyro product category list` | Kategorileri Listeleme | `GET /products/categories` |
| `gyro product get` | Ürün Detayı ve Sipariş Şeması | `GET /products/{id}` |
| `gyro project create` | Proje oluştur | `POST /cloud/projects` |
| `gyro project delete` | Projeyi sil | `DELETE /cloud/projects/{id}` |
| `gyro project get` | Projeyi hizmet numaralarıyla getir | `GET /cloud/projects/{id}` |
| `gyro project list` | Projeleri listele | `GET /cloud/projects` |
| `gyro project move` | Hizmetleri projeye taşı | `POST /cloud/projects/{id}/services` |
| `gyro project rename` | Projeyi yeniden adlandır / renklendir | `PATCH /cloud/projects/{id}` |
| `gyro reference country list` | Get countries | `GET /reference/countries` |
| `gyro reference country state list` | Get states | `GET /reference/countries/{code}/states` |
| `gyro reference state city list` | Get cities | `GET /reference/states/{id}/cities` |
| `gyro reseller assign` | Hizmeti uç-müşteriye ata (customer_id 0 = atamayı kaldır) | `POST /cloud/reseller/assign` |
| `gyro reseller customer create` | Uç-müşteri oluştur | `POST /cloud/reseller/customers` |
| `gyro reseller customer delete` | Uç-müşteriyi sil (anahtarlar ve webhook'lar iptal, hizmet ataması kalkar) | `DELETE /cloud/reseller/customers/{id}` |
| `gyro reseller customer get` | Uç-müşteriyi getir | `GET /cloud/reseller/customers/{id}` |
| `gyro reseller customer key create` | Uç-müşteri için anahtar üret (token bir kez döner) | `POST /cloud/reseller/customers/{id}/keys` |
| `gyro reseller customer key delete` | Uç-müşteri anahtarını iptal et | `DELETE /cloud/reseller/customers/{id}/keys/{kid}` |
| `gyro reseller customer key list` | Uç-müşterinin anahtarlarını listele | `GET /cloud/reseller/customers/{id}/keys` |
| `gyro reseller customer list` | Uç-müşterileri listele | `GET /cloud/reseller/customers` |
| `gyro reseller customer update` | Uç-müşteriyi güncelle (askıya alma = anahtarları çalışmaz) | `PATCH /cloud/reseller/customers/{id}` |
| `gyro reseller report get` | Uç-müşteri başına kullanım ve fatura raporu | `GET /cloud/reseller/report` |
| `gyro reseller setting get` | Bayi ayarları | `GET /cloud/reseller/settings` |
| `gyro reseller setting set` | Bayi ayarlarını kaydet (beyaz etiket) | `PUT /cloud/reseller/settings` |
| `gyro security-group attach` | Güvenlik grubunu sunucuya bağla | `POST /cloud/servers/{id}/security-groups/{sg}` |
| `gyro security-group confirm` | Kural denemesini kalıcı yap | `POST /cloud/security-groups/{sg}/confirm` |
| `gyro security-group create` | Güvenlik grubu oluştur | `POST /cloud/security-groups` |
| `gyro security-group delete` | Güvenlik grubunu sil | `DELETE /cloud/security-groups/{sg}` |
| `gyro security-group detach` | Güvenlik grubunu sunucudan ayır | `DELETE /cloud/servers/{id}/security-groups/{sg}` |
| `gyro security-group get` | Güvenlik grubunu getir | `GET /cloud/security-groups/{sg}` |
| `gyro security-group list` | Güvenlik gruplarını listele | `GET /cloud/security-groups` |
| `gyro security-group rules set` | Güvenlik grubunun kurallarını değiştir | `PUT /cloud/security-groups/{sg}/rules` |
| `gyro security-group server list` | Sunucuya bağlı güvenlik grupları | `GET /cloud/servers/{id}/security-groups` |
| `gyro security-group server set` | Bağlı grupları tam liste olarak ayarla | `PUT /cloud/servers/{id}/security-groups` |
| `gyro security-group update` | Güvenlik grubunun adını / açıklamasını değiştir | `PATCH /cloud/security-groups/{sg}` |
| `gyro server activity list` | Sunucudaki son platform işlemleri | `GET /cloud/servers/{id}/activity` |
| `gyro server app create` | Uygulama kur (diski yeniden kurar; confirm zorunlu; sunucu başına günde 3) | `POST /cloud/servers/{id}/app` |
| `gyro server app get` | Uygulama kurulum durumu (hazır olunca kimlik bir kez) | `GET /cloud/servers/{id}/app` |
| `gyro server app log get` | Son uygulama kurulumunun canlı ilerleme günlüğü (adımlar, sunucu günlüğü) | `GET /cloud/servers/{id}/app/log` |
| `gyro server app retry` | Hata veren ya da yanıt vermeyen kurulumu yeniden dene (diski yeniden kurar; confirm zorunlu; sunucu başına saatte 3) | `POST /cloud/servers/{id}/app/retry` |
| `gyro server backup create` | Elle yedek al (yedekleme ek hizmeti gerekir) | `POST /cloud/servers/{id}/backups` |
| `gyro server backup list` | Otomatik yedekleri listele | `GET /cloud/servers/{id}/backups` |
| `gyro server console` | Tek seferlik VNC konsol bileti al | `POST /cloud/servers/{id}/console` |
| `gyro server create` | Bulut sunucu oluştur (sipariş) | `POST /cloud/servers` |
| `gyro server delete` | Sunucuyu sil (iptal) | `DELETE /cloud/servers/{id}` |
| `gyro server disk add` | Ek disk ekle | `POST /cloud/servers/{id}/disks` |
| `gyro server disk list` | Sunucunun ek diskleri | `GET /cloud/servers/{id}/disks` |
| `gyro server disk remove` | Ek diski sil (verisiyle) | `DELETE /cloud/servers/{id}/disks/{disk}` |
| `gyro server disk resize` | Ek diski büyüt | `PUT /cloud/servers/{id}/disks/{disk}` |
| `gyro server firewall list` | Gelen trafik güvenlik duvarı kuralları | `GET /cloud/servers/{id}/firewall` |
| `gyro server firewall rule add` | Gelen trafik kuralı ekle | `POST /cloud/servers/{id}/firewall/rules` |
| `gyro server firewall rule delete` | Güvenlik duvarı kuralı sil (son kural confirm ister) | `DELETE /cloud/servers/{id}/firewall/rules/{n}` |
| `gyro server get` | Sunucuyu canlı durumuyla getir | `GET /cloud/servers/{id}` |
| `gyro server ip add` | Ek IPv4 adresi ekle | `POST /cloud/servers/{id}/ips` |
| `gyro server ip list` | Sunucunun IPv4 adresleri | `GET /cloud/servers/{id}/ips` |
| `gyro server ip remove` | Ek IPv4 adresini bırak | `DELETE /cloud/servers/{id}/ips/{ip}` |
| `gyro server list` | Bulut sunucuları listele | `GET /cloud/servers` |
| `gyro server metrics` | CPU, disk ve ağ zaman serileri | `GET /cloud/servers/{id}/metrics` |
| `gyro server power off` | Sunucu işlemi çalıştır | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server power on` | Sunucu işlemi çalıştır | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server power reboot` | Sunucu işlemi çalıştır | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server power reset` | Sunucu işlemi çalıştır | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server power shutdown` | Sunucu işlemi çalıştır | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server rdns list` | Ters DNS kayıtlarını listele | `GET /cloud/servers/{id}/rdns` |
| `gyro server rdns set` | Bir adresin PTR kaydını ayarla (veya sıfırla) | `PUT /cloud/servers/{id}/rdns` |
| `gyro server rebuild` | Sunucu işlemi çalıştır | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server rescue disable` | Sunucu işlemi çalıştır | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server rescue enable` | Sunucu işlemi çalıştır | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server reset-password` | Sunucu işlemi çalıştır | `POST /cloud/servers/{id}/actions/{action}` |
| `gyro server snapshot create` | Snapshot al | `POST /cloud/servers/{id}/snapshots` |
| `gyro server snapshot delete` | Snapshot sil | `DELETE /cloud/servers/{id}/snapshots/{sid}` |
| `gyro server snapshot list` | Snapshot'ları ve kotayı listele | `GET /cloud/servers/{id}/snapshots` |
| `gyro server snapshot restore` | Snapshot'ı sunucuya geri yükle (diski ezer) | `POST /cloud/servers/{id}/snapshots/{sid}/restore` |
| `gyro server storage-mount command` | Yeni kurulum komutu (bağlantı anahtarı yenilenir) | `POST /cloud/servers/{id}/storage-mounts/{mid}/command` |
| `gyro server storage-mount create` | Kovayı sunucuya bağla | `POST /cloud/servers/{id}/storage-mounts` |
| `gyro server storage-mount get` | Bağlantı ayrıntısı | `GET /cloud/servers/{id}/storage-mounts/{mid}` |
| `gyro server storage-mount list` | Sunucuya bağlı kovalar | `GET /cloud/servers/{id}/storage-mounts` |
| `gyro server storage-mount remove` | Bağlantıyı kaldır | `DELETE /cloud/servers/{id}/storage-mounts/{mid}` |
| `gyro server update | rename` | Sunucu adını değiştir veya otomatik yenilemeyi aç/kapat | `PATCH /cloud/servers/{id}` |
| `gyro service addon available get` | Bu hizmete alınabilecek ek hizmetler | `GET /cloud/services/{id}/addons/available` |
| `gyro service addon create` | Ek hizmet satın al (bakiyeden ödenir) | `POST /cloud/services/{id}/addons` |
| `gyro service addon delete` | Ön ödemeli ek hizmeti hemen kaldır | `DELETE /cloud/services/{id}/addons/{aid}` |
| `gyro service addon list` | Hizmetin ek hizmetleri | `GET /cloud/services/{id}/addons` |
| `gyro service cancellation create` | İptal Talebi Açma | `POST /services/{id}/cancellation` |
| `gyro service cancellation delete` | İptal Talebini Geri Çekme | `DELETE /services/{id}/cancellation` |
| `gyro service dashboard get` | Panel Özetini Okuma | `GET /services/{id}/dashboard` |
| `gyro service get` | Hizmeti getir (yenileme, iptal, erişim bilgileri) | `GET /cloud/services/{id}` |
| `gyro service invoice list` | Hizmetin Faturalarını Okuma | `GET /services/{id}/invoices` |
| `gyro service label get` | Hizmetin takma adını ve etiketlerini getir | `GET /cloud/services/{id}/label` |
| `gyro service label set` | Hizmetin takma adını ve/veya etiketlerini ayarla | `PUT /cloud/services/{id}/label` |
| `gyro service list` | Tüm ürün hatlarındaki hizmetleri listele | `GET /cloud/services` |
| `gyro service method call` | Müşteri-çağrılabilir modül metodunu çağır (örn. lisans change_ip) | `POST /cloud/services/{id}/methods/{method}` |
| `gyro service metrics billing list` | Kullanım Fatura Geçmişi | `GET /services/{id}/metrics/billing` |
| `gyro service metrics list` | Ölçümleri Okuma | `GET /services/{id}/metrics` |
| `gyro service metrics update` | Ölçümü Açma ve Kapatma | `PATCH /services/{id}/metrics/{key}` |
| `gyro service module-method` | Modül Metodu Çalıştırma | `POST /services/{id}/module-method` |
| `gyro service password` | Panel Parolasını Değiştirme | `POST /services/{id}/password` |
| `gyro service project set` | Hizmetin projesini ayarla | `PUT /cloud/services/{id}/project` |
| `gyro service renew` | Yenileme Faturası Kesme | `POST /services/{id}/renew` |
| `gyro service resize` | Paketi yükselt (bakiyeden ödenir) | `POST /cloud/services/{id}/resize` |
| `gyro service resize-option list` | Hizmetin yükseltilebileceği ürünler ve fiyatları | `GET /cloud/services/{id}/resize-options` |
| `gyro service sso` | Panele Tek Oturum Açma | `POST /services/{id}/sso` |
| `gyro service tool get` | Bir aracı oku (posta kutuları, snapshot'lar, DNS kayıtları, ...) | `GET /cloud/services/{id}/tools/{tool}` |
| `gyro service tool list` | Hizmetin araçlarını, eylemlerini ve güvenlik politikasını keşfet | `GET /cloud/services/{id}/tools` |
| `gyro service tool run` | Bir araç eylemi çalıştır (posta kutusu oluştur, snapshot'tan dön, ...) | `POST /cloud/services/{id}/tools/{tool}/{action}` |
| `gyro service update` | Hizmet Tercihlerini Değiştirme | `PATCH /services/{id}` |
| `gyro sms message get` | Tek Gönderimi Okuma | `GET /sms/messages/{id}` |
| `gyro sms message list` | Gönderim Geçmişi | `GET /sms/messages` |
| `gyro sms quote` | Fiyat Alma | `POST /sms/quote` |
| `gyro sms send` | Toplu Gönderim | `POST /sms/send` |
| `gyro sms sender list` | Başlıkları Listeleme | `GET /sms/senders` |
| `gyro ssh-key create` | SSH anahtarı ekle | `POST /cloud/ssh-keys` |
| `gyro ssh-key delete` | SSH anahtarını sil | `DELETE /cloud/ssh-keys/{id}` |
| `gyro ssh-key get` | SSH anahtarını getir | `GET /cloud/ssh-keys/{id}` |
| `gyro ssh-key list` | SSH anahtarlarını listele | `GET /cloud/ssh-keys` |
| `gyro storage bucket rename` | Kova adını değiştir | `POST /cloud/storage/{id}/buckets/{bucket}/rename` |
| `gyro storage connect get` | connect (GyroDepolama) | `GET /cloud/storage/{id}/connect` |
| `gyro storage cors delete` | Kovanın tüm CORS kurallarını kaldır | `DELETE /cloud/storage/{id}/buckets/{bucket}/cors` |
| `gyro storage cors get` | Kovanın CORS kuralları | `GET /cloud/storage/{id}/buckets/{bucket}/cors` |
| `gyro storage cors set` | Kovanın CORS kurallarını değiştir | `PUT /cloud/storage/{id}/buckets/{bucket}/cors` |
| `gyro storage folder count` | Klasördeki nesneleri say | `GET /cloud/storage/{id}/buckets/{bucket}/folders/count` |
| `gyro storage folder create` | Klasör oluştur | `POST /cloud/storage/{id}/buckets/{bucket}/folders` |
| `gyro storage folder delete` | Klasörü içindekilerle sil | `DELETE /cloud/storage/{id}/buckets/{bucket}/folders` |
| `gyro storage key rotate` | access-keys:rotate (GyroDepolama) | `POST /cloud/storage/{id}/keys/{key}/rotate` |
| `gyro storage lifecycle noncurrent get` | Kovanın eski sürüm silme kuralı | `GET /cloud/storage/{id}/buckets/{bucket}/lifecycle/noncurrent` |
| `gyro storage lifecycle noncurrent set` | Eski sürüm silme süresini ayarla | `PUT /cloud/storage/{id}/buckets/{bucket}/lifecycle/noncurrent` |
| `gyro storage object delete` | Nesneleri sil | `POST /cloud/storage/{id}/buckets/{bucket}/objects/delete` |
| `gyro storage object download-url` | İmzalı indirme adresi | `POST /cloud/storage/{id}/buckets/{bucket}/objects/download-url` |
| `gyro storage object legal-hold get` | Nesne sürümünün legal hold durumu | `GET /cloud/storage/{id}/buckets/{bucket}/objects/legal-hold` |
| `gyro storage object legal-hold set` | Legal hold koy / kaldır | `PUT /cloud/storage/{id}/buckets/{bucket}/objects/legal-hold` |
| `gyro storage object list` | Kova klasöründeki nesneleri listele | `GET /cloud/storage/{id}/buckets/{bucket}/objects` |
| `gyro storage object retention get` | Nesne sürümünün saklaması | `GET /cloud/storage/{id}/buckets/{bucket}/objects/retention` |
| `gyro storage object retention set` | Nesne sürümünü kilitle | `PUT /cloud/storage/{id}/buckets/{bucket}/objects/retention` |
| `gyro storage object upload-url` | İmzalı yükleme adresleri | `POST /cloud/storage/{id}/buckets/{bucket}/objects/upload-url` |
| `gyro storage object-lock get` | Kovanın varsayılan saklaması | `GET /cloud/storage/{id}/buckets/{bucket}/object-lock` |
| `gyro storage object-lock scan` | Kilitli sürümleri şimdi tara | `POST /cloud/storage/{id}/object-lock/scan` |
| `gyro storage object-lock set` | Kovanın varsayılan saklamasını ayarla | `PUT /cloud/storage/{id}/buckets/{bucket}/object-lock` |
| `gyro storage object-lock status` | Korumalı Depolama: nesne kilidi özeti | `GET /cloud/storage/{id}/object-lock` |
| `gyro storage usage get` | usage (GyroDepolama) | `GET /cloud/storage/{id}/usage` |
| `gyro storage versions delete` | Sürümü kalıcı sil | `DELETE /cloud/storage/{id}/buckets/{bucket}/versions` |
| `gyro storage versions list` | Nesne sürümlerini listele | `GET /cloud/storage/{id}/buckets/{bucket}/versions` |
| `gyro storage versions restore` | Eski sürümü geri yükle | `POST /cloud/storage/{id}/buckets/{bucket}/versions/restore` |
| `gyro storage versions undelete` | Silme işaretini kaldır | `POST /cloud/storage/{id}/buckets/{bucket}/versions/undelete` |
| `gyro storage website disable` | Statik web sitesini kapat | `DELETE /cloud/storage/{id}/buckets/{bucket}/website` |
| `gyro storage website get` | Kovanın statik web sitesi ayarları | `GET /cloud/storage/{id}/buckets/{bucket}/website` |
| `gyro storage website set` | Statik web sitesini aç / güncelle | `PUT /cloud/storage/{id}/buckets/{bucket}/website` |
| `gyro template create` | Sunucudan şablon oluştur | `POST /cloud/servers/{id}/templates` |
| `gyro template delete` | Şablonu sil | `DELETE /cloud/templates/{id}` |
| `gyro template deploy` | Şablondan yeni sunucu | `POST /cloud/templates/{id}/servers` |
| `gyro template get` | Şablonu getir | `GET /cloud/templates/{id}` |
| `gyro template list` | Şablonlarınızı listele | `GET /cloud/templates` |
| `gyro template rebuild` | Sunucuyu şablonla yeniden kur | `POST /cloud/templates/{id}/rebuild` |
| `gyro template update` | Şablonu yeniden adlandır | `PATCH /cloud/templates/{id}` |
| `gyro ticket access-group list` | Erişim Gruplarını Listeleme | `GET /tickets/access-groups` |
| `gyro ticket attachment get` | Eki İndirme | `GET /tickets/{id}/attachments/{aid}` |
| `gyro ticket close` | Talebi Kapatma | `POST /tickets/{id}/close` |
| `gyro ticket create` | Talep Açma | `POST /tickets` |
| `gyro ticket department list` | Departmanları Listeleme | `GET /tickets/departments` |
| `gyro ticket get` | Talebi ve Mesajları Okuma | `GET /tickets/{id}` |
| `gyro ticket list` | Talepleri Listeleme | `GET /tickets` |
| `gyro ticket rating` | Talebi Puanlama | `POST /tickets/{id}/rating` |
| `gyro ticket reopen` | Talebi Yeniden Açma | `POST /tickets/{id}/reopen` |
| `gyro ticket reply` | Yanıt Yazma | `POST /tickets/{id}/reply` |
| `gyro ticket reply rating` | Personel Yanıtını Puanlama | `POST /tickets/{id}/replies/{rid}/rating` |
| `gyro user create` | Alt kullanıcı davet et | `POST /cloud/users` |
| `gyro user delete` | Alt kullanıcıyı sil | `DELETE /cloud/users/{id}` |
| `gyro user list` | Alt kullanıcıları listele | `GET /cloud/users` |
| `gyro user resend-invite` | Daveti yeniden gönder | `POST /cloud/users/{id}/resend-invite` |
| `gyro user update` | Alt kullanıcıyı güncelle | `PATCH /cloud/users/{id}` |
| `gyro webhook create` | Webhook oluştur (gizli anahtar bir kez döner) | `POST /cloud/webhooks` |
| `gyro webhook delete` | Webhook'u sil | `DELETE /cloud/webhooks/{id}` |
| `gyro webhook delivery list` | Son 100 teslim | `GET /cloud/webhooks/{id}/deliveries` |
| `gyro webhook get` | Webhook'u getir | `GET /cloud/webhooks/{id}` |
| `gyro webhook list` | Webhook'ları listele | `GET /cloud/webhooks` |
| `gyro webhook test` | Şimdi bir api.test olayı gönder | `POST /cloud/webhooks/{id}/test` |
| `gyro webhook update` | Adres, olay, açıklama ya da durumu güncelle | `PATCH /cloud/webhooks/{id}` |
| `gyro whoami` | Çağıran API anahtarını göster | `GET /whoami` |
