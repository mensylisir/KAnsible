# ping_vip.sh
VIP_IF=ens160
MASTER_VIP=192.168.81.11
for node_ip in 192.168.80.60 192.168.80.61
  do
    echo ">>> ${node_ip}"
    ssh -p7122 ${node_ip} "/usr/sbin/ip addr show ${VIP_IF}"
    ssh -p7122 ${node_ip} "ping -c 5 ${MASTER_VIP}"
  done
