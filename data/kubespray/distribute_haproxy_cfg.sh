# distribute_haproxy_cfg.sh
for node_ip in 192.168.80.60 192.168.80.61
  do
    echo ">>> ${node_ip}"
    scp -P 7122 /etc/haproxy/haproxy.cfg root@${node_ip}:/etc/haproxy
  done
