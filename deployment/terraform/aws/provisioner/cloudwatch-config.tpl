{
    "agent": {
      "metrics_collection_interval": 10,
      "logfile": "/opt/aws/amazon-cloudwatch-agent/logs/amazon-cloudwatch-agent.log",
      "run_as_user": "root"
    },
    "logs": {
      "logs_collected": {
        "files": {
          "collect_list": [
            {
              "file_path": "/opt/aws/amazon-cloudwatch-agent/logs/amazon-cloudwatch-agent.log",
              "log_group_name": "/apps/CloudWatchAgentLog/",
              "log_stream_name": "{ip_address}_{instance_id}",
              "timezone": "Local"
            },
            {
              "file_path": "/var/log/messages",
              "log_group_name":  "/apps/cosmovisor/",
              "log_stream_name": "{ip_address}_{instance_id}",
              "timestamp_format": "%b %d %H:%M:%S",
              "timezone": "Local"
            }
          ]
        }
      }
    }
  }