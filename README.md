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
slack-status [--away] [--group group] [--workspace workspace] [:emoji:] [status_message]
```

```shell
# set away status
slack-status --away :car: on the road

# set status for a group of workspaces
slack-status -g work :hamburger: lunch

# clear all status and set to active
slack-status
```

configuration
=============
Configuration is defined in a toml file at `~/.slack-status`

```
[my_org]
token = "SLACK_TOKEN"
groups = ["work"]

[my_other_org]
token = "SLACK_TOKEN"
groups = ["work"]

[my_other_other_org]
token = "SLACK_TOKEN"
```
