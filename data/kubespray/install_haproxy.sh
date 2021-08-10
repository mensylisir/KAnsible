# install_haproxy.sh
for node_ip in 192.168.80.60 192.168.80.61
  do
    echo ">>> ${node_ip}"
    ssh -p7122 root@${node_ip} "yum install -y keepalived haproxy"
  done
