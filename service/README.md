# How to set it up as a system service

0. `cp useProxy /usr/local/bin/useProxy`
1. `vim /etc/systemd/system/useProxy.service` and add the `useProxy.service` file
2. `systemctl daemon-reload`
3. `systemctl restart useProxy`

If the `useProxy` executable changed, please run step 2 and step 3

