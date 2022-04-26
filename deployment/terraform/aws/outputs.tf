output "nodes" {
    value = local.nodes
}

output "ansible_inventory" {
    value = { 
        all = {
            children = {

                genesis = {
                    hosts = { for host in local.nodes.validator.public_ips: host => null } 
                }
                
                validators = {
                    hosts = { for host in local.nodes.validator.public_ips: host => null } 
                }

                private_sentries = {
                    hosts = { for host in local.nodes.private_sentries.public_ips: host => null } 
                }

                public_sentries = {
                    hosts = { for host in local.nodes.public_sentries.public_ips: host => null }  
                }

                seeds = { 
                    hosts = { for host in local.nodes.seeds.public_ips: host => null }
                }

                observers = {
                    hosts = { for host in local.nodes.observers.public_ips: host => null }
                }
            }
        }
    }
}