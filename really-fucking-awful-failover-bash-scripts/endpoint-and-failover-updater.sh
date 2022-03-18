source init.sh
TTL=60

# OVH Primary
#NA4=54.39.106.25
#NA6=2607:5300:203:3c19::2

# OVH Backup
BA4=54.39.106.25
BA6=2607:5300:203:3c19::2

# Hetzner 1 Primary
NA4=95.217.228.176
NA6=2a01:4f9:4b:4c8f::2

# Hetzner 1 Backup 
#BA4=95.217.228.176
#BA6=2a01:4f9:4b:4c8f::2

# Hetzner 2 Backup
#BA4=195.201.86.186
#BA6=2a01:4f8:13b:4285::2

# Hetner 2 Primary
#NA4=195.201.86.186
#NA6=2a01:4f8:13b:4285::2


echo "  wtfismyip.com -  a"

curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189570&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189570&host=&record=$NA4&ttl=$TTL" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189570&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=1&main_ip=$NA4&backup_ip_1=$BA4&&check_period=60&" ; echo " ; echo "

echo " wtfismyip.com -  aaaa"
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json"?auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189575&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189575&host=&record=$NA6&ttl=$TTL&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189575&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=1&main_ip=$NA6&backup_ip_1=$BA6&&check_period=60&" ; echo " ; echo "

echo "ipv4.wtfismyip.com - a"
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189596&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189596&host=ipv4&record=$NA4&ttl=$TTL&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189596&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=1&main_ip=$NA4&backup_ip_1=$BA4&&check_period=60&" ; echo " ; echo "

echo "ipv6.wtfismyip.com - aaaa"
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189599&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189599&host=ipv6&record=$NA6&ttl=$TTL&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189599&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=1&main_ip=$NA6&backup_ip_1=$BA6&&check_period=60&" ; echo " ; echo "

echo " myip.wtf - a"
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189643&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189643&host=&record=$NA4&ttl=$TTL&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189643&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=1&main_ip=$NA4&backup_ip_1=$BA4&&check_period=60&" ; echo " ; echo "

echo " myip.wtf - aaaa"
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189645&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189645&host=&record=$NA6&ttl=$TTL&" ; echo " ; echo "
curl --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189645&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=1&main_ip=$NA6&backup_ip_1=$BA6&&check_period=60&" ; echo " ; echo "
