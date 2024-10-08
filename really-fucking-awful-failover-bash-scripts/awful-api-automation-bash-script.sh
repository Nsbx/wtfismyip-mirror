#!/bin/bash
#host myip.wtf
# Backup with no recovery
bash ../nsupdate/check.sh 

sleep 5

source init.sh
TTL=60

# Hetzner backup
BA4=65.108.75.112
BA6=2a01:4f9:6b:4b55::acab

# OVH primary
NA4=15.204.2.228
NA6=2604:2dc0:200:1014::acab

echo "  wtfismyip.com -  a"

curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189570&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189570&host=&record=$NA4&ttl=$TTL" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189570&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=0&main_ip=$NA4&backup_ip_1=$BA4&&check_period=60&backup_ip_2=$BA42" ; echo " ; echo "

echo " wtfismyip.com -  aaaa"
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json"?auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189575&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189575&host=&record=$NA6&ttl=$TTL&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189575&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=0&main_ip=$NA6&backup_ip_1=$BA6&&check_period=60&backup_ip_2=$BA62" ; echo " ; echo "

echo "ipv4.wtfismyip.com - a"
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189596&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189596&host=ipv4&record=$NA4&ttl=$TTL&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189596&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=0&main_ip=$NA4&backup_ip_1=$BA4&&check_period=60&backup_ip_2=$BA42" ; echo " ; echo "

echo "ipv6.wtfismyip.com - aaaa"
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189599&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189599&host=ipv6&record=$NA6&ttl=$TTL&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=wtfismyip.com&record-id=256189599&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=0&main_ip=$NA6&backup_ip_1=$BA6&&check_period=60&backup_ip_2=$BA62" ; echo " ; echo "

echo " myip.wtf - a"
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189643&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189643&host=&record=$NA4&ttl=$TTL&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189643&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=0&main_ip=$NA4&backup_ip_1=$BA4&&check_period=60backup_ip_2=$BA42&" ; echo " ; echo "

echo " myip.wtf - aaaa"
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-deactivate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189645&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/mod-record.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189645&host=&record=$NA6&ttl=$TTL&" ; echo " ; echo "
curl -4 --user-agent "wtfdnsapi v0.1" https://api.cloudns.net/dns/failover-activate.json?"auth-id=$CLOUDNS_AUTH_ID&auth-password=$CLOUDNS_AUTH_PASSWORD&domain-name=myip.wtf&record-id=256189645&check_type=4&host=wtfismyip.com&port=80&path=test&down_event_handler=2&up_event_handler=0&main_ip=$NA6&backup_ip_1=$BA6&&check_period=60&backup_ip_2=$BA62" ; echo " ; echo "
