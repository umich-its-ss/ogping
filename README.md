# ogping - a simple OpsGenie heartbeat ping

This simple tool will ping a OpsGenie heartbeat.

You might have some automated tools that either don't emit any status when they succeed, or they send out alerts saying everything is working. Neither is particularly good - silence coould also mean the tool isn't running at all, and alerts all the time can lead to mail rules to just delete them.

OpsGenie has an easy to use heartbeat function that you can leverage to let you know when a process has stopped working successfully.

## Usage

First, create an OpsGenie "API" integration. This API key can be used among all heartbeats on your team. Configure it as:
 - :green_check_mark: Read Access
 - :green_check_mark: Create and Update Access
 - :empty_check_box: Delete Access
 - :green_check_mark: Restrict Configuration Access
 - :green_check_mark: Enabled

Note the API key provided.

Then create a heartbeat, configure the deadline and priority as appropriate. For example, we have a service that runs twice a day, but gave it a deadline of 2 days. You'll want to factor in how critical that run is versus likelihood of minor failures.

In your automation, fire off this utility. For example:

```
ogping -api-key 1f53d805-562e-47a6-800c-26b64a77efaa -name 'DuplicateDetector'
```

**Note:** OpsGenie provides no feedback if the heartbeat name is incorrect. You should verify that the ping has been reflected in the UI.

Instead of either command line argument, `ogping` also will look for the `OPSGENIE_API_KEY` and `OPSGENIE_HEARTBEAT` environment variables. You may find that easier to use with kubernetes cron jobs.
