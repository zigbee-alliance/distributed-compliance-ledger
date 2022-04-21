output "nodes" {
    value = {
        validator = {
            private_ips = module.validator.private_ips
            public_ips = module.validator.public_ips
        },

        private_sentries = {
            private_ips = var.private_sentries_config.enable ? module.private_sentries[0].private_ips : []
            public_ips  = var.private_sentries_config.enable ? module.private_sentries[0].public_ips : []
        },

        public_sentries = {
            public_ips  = concat(
                                    (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 1)) ? module.public_sentries_1[0].public_ips : [], 
                                    (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 2)) ? module.public_sentries_2[0].public_ips : [],
                                )
        },

        seeds = {
            public_ips  = concat(
                                    (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 1)) ? module.public_sentries_1[0].seed_public_ips : [], 
                                    (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 2)) ? module.public_sentries_2[0].seed_public_ips : [],
                                )
        },

        observers = {
            public_ips  = concat(
                                    (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.regions, 1)) ? module.observers_1[0].public_ips : [], 
                                    (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.regions, 2)) ? module.observers_2[0].public_ips : [],
                                )
        },
    }
}