# Deploying DC Ledger UI

1. Run `dclcli` rest service on the client machine. See `deployment/ansible` folder in the `dc-ledger` repository.
2. Build the app by running `ng build --prod`.
3. Execute playbooks in the following order: `content`, `nginx`, `config`, `reload`.
4. Use certbot to obtain a certificate and don't forger to open `443` port.
