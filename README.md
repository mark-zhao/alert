#### 修改C:\Users\Think\Desktop\alert\goDemo2\src\common\xml.go
#### 信息再企业微信--> 应用管理--> api接受消息里
```
const corpID = "123456"
const key = "123456"
const token = "123456"
```
#### 修改C:\Users\Think\Desktop\alert\goDemo2\src\alert\alert.go
#### 信息在企业微信--> 应用管理里
```
const tokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
const msgUrl = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
//上面不要改
const weChatAgentId = 1000002
const corpid = "123456"
const secret = "123456"
```
#### 启动文件
```
# /etc/systemd/system/alert-control.service
[Unit]
Description=ceph cluster alert control service
After=network-online.target firewalld.service
Wants=network-online.target

[Service]
Type=simple
# the default is not to use systemd for cgroups because the delegate issues still
# exists and systemd currently does not support the cgroup feature set required
# for containers run by docker
ExecStart=/root/zzc/alert/alert-control
ExecStop=/bin/kill -s HUP $MAINPID
Restart=always

[Install]
WantedBy=multi-user.target
```
#### 配置文件
```
[root@host-211-159-204-110 alert]# cat /etc/ceph/cephAlert.json 
{
    "regions": "JDOS",
    "JDOS":{
        "addr": "1.1.1.1:1234"    
    },
    "instructionSet":{
        "help" : "01: ceph cluster status 
                  02: ceph cluster status detail
                  03: ceph pg repair
                  04: ceph osd set nobackfill
                  05: ceph osd set norecover
                  06: ceph osd unset nobackfill
                  07: ceph osd unset norecover
                  08: pool usage",

        "status":"show all cluster status",
        "add": "admin add member to whiteList",
        "adjust": "adjust alert check time(seconds)",
        "01": "ceph -s",
        "02": "ceph health detail",    
        "03": "ceph health detail | grep pg | grep -v pgs | awk '{ print $2 }' | xargs -n 1 -i ceph pg repair {}",
        "04": "ceph osd set nobackfill",
        "05": "ceph osd set norecover",
        "06": "ceph osd unset nobackfill",
        "07": "ceph osd unset norecover",
        "08": "cat /root/chengpeng/alert/poolusage.txt"
    },

    "whiteList": {
        "DingJiaoJiao": "admin",
        "zhaozhicheng": "member"
    }
}
[root@host-211-159-204-110 alert]# cat /etc/glog/glog.json 
{
  "glog":{
     "log_dir": "/var/log/glog/",
     "toStderr": false,
     "alsoToStderr": false,
     "verbosity": 0,
     "stderrThreshold": 1
    }   
}
mkdir /var/log/glog/
```
#### iptables
```
/sbin/iptables -F
/sbin/iptables  -A INPUT  -p tcp -m tcp --dport 2223 -j ACCEPT
/sbin/iptables -A INPUT -i lo -j ACCEPT
/sbin/iptables -A OUTPUT -o lo -j ACCEPT
/sbin/iptables  -A INPUT  -s 1.1.1.1/24 -p tcp -m tcp --dport 443 -j ACCEPT
```
#### nginx 配置
```
[root@host-211-159-204-110 alert]# cat /etc/nginx/nginx.conf
# For more information on configuration, see:
#   * Official English Documentation: http://nginx.org/en/docs/
#   * Official Russian Documentation: http://nginx.org/ru/docs/

user nginx;
worker_processes 2;
error_log /var/log/nginx/error.log;
pid /run/nginx.pid;

# Load dynamic modules. See /usr/share/nginx/README.dynamic.
include /usr/share/nginx/modules/*.conf;

events {
    worker_connections 1024;
}

http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile            on;
    tcp_nopush          on;
    tcp_nodelay         on;
    keepalive_timeout   65;
    types_hash_max_size 2048;

    include             /etc/nginx/mime.types;
    default_type        application/octet-stream;
    include /etc/nginx/conf.d/*.conf;

    server {
        listen       443;
        server_name  zzcweixin.ztgame.com.cn;
        #location / {
        #    root   html;
        #    index  index.html index.htm;
        #}
        ssl on;
        ssl_certificate /usr/local/nginx/conf/key/Server_Wildcard_com_cn_20171128.cer;
        ssl_certificate_key /usr/local/nginx/conf/key/Server_Wildcard_com_cn_20171128.key;
        ssl_session_timeout 5m;
        ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
        ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
        ssl_prefer_server_ciphers on;
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
        location / {
            proxy_pass   http://127.0.0.1:8080;
        }
    }
}
```
