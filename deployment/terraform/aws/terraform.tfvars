region_1 = "us-west-1"
region_2 = "us-east-2"

private_sentries_config = {
    enable = true
    nodes_count = 2
}

public_sentries_config = {
    enable = true
    enable_ipv6 = true
    nodes_count = 2

    regions = [
        1,
        2
    ]
}

observers_config = {
    enable = true
    nodes_count = 3
    root_domain_name = "matterprotocol.com"
    enable_tls = true

    regions = [
        1,
        2
    ]
}