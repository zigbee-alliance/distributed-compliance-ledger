global:
  scrape_interval: 15s
  external_labels:
    monitor: '${monitor}'

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: [${targets}]

remote_write:
  -
    url: ${prometheus_endpoint}api/v1/remote_write
    queue_config:
        max_samples_per_send: 1000
        max_shards: 200
        capacity: 2500
    sigv4:
        region: ${region}
