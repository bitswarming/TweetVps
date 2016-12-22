# tweetvps 
### Drone [![Build Status](https://drone.io/github.com/bitswarming/TweetVps/status.png)](https://drone.io/github.com/bitswarming/TweetVps/latest) Travis [![Build Status](https://travis-ci.org/bitswarming/TweetVps.svg?branch=master)](https://travis-ci.org/bitswarming/TweetVps)

Service for vps which listens twitter for commands to execute,which are ecrypted with symmetric encyption. If you blocked yourself with iptables or deleted your ssh publick key, or even user, this tool is your last chance, before server resetting.


## How to use
1. Create twitter account
1. clone build, or download
1. at vps run `sudo crontab -e` and insert 
`cd /opt/rebootclient/;
sudo ./rebootclient daemon --user=$twitteruser --expire=3500 --secret=$secretcode --mydomain=$target --period=$seconds`
1. generete encrypted command` rebootclient.go encode  --secret=xxx --targetdomain=quad.node.consul --message=reboot:robot.node.consul`
1. post command to twitter.
