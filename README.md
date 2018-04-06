slack-status
============

Update presence and status in multiple Slack workspaces.

install
=======

```shell
go get github.com/dstokes/slack-status
```

usage
=====

```shell
slack-status [--away] [:emoji:] [status_message]
```

```shell
# set away status
slack-status --away :car: on the road

# clear all status and set to active
slack-status
```

configuration
=============
Configuration is defined in a toml file at `~/.slack-status`

```
[my_org]
token = "SLACK_TOKEN"

[my_other_org]
token = "SLACK_TOKEN"
```
