# tweetvps
Service for vps server which listens for encypted commands to execute, from your twitter account. If you blocked yourself with iptables or deleted your ssh key, or user this is your last hope.


## How to use
1. Create twitter account
1. clone
1. run `sudo crontab -e` and insert 
`cd /opt/rebootclient/;
sudo ./rebootclient daemon --user=$twitteruser --expire=3500 --secret=$secretcode --mydomain=$target --period=$seconds`
` rebootclient.go encode  --secret=xxx --targetdomain=quad.node.consul --message=reboot:robot.node.consul`
