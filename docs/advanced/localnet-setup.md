# Setup of Local Net in Docker

- Install Oracle VM VirtualBox
- Download Ubuntu 18.04.3 VDI for VirtualBox
- In VirtualBox create a VM using the downloaded VDI
- Increase processors count and RAM amount in the created VM settings (set them to a half of physical facilities)
- Start and log in to the created VM using the following credentials: osboxes/osboxes.org
- In Ubuntu VM:
  - Install Go as described at <https://golang.org/doc/install>
    - Ensure that the following line has been appended to `/etc/profile`:

    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```

    - Ensure that the following line has been appended to `~/.profile`:

    ```bash
    export PATH=$PATH:~/go/bin
    ```

  - Install Docker as described at <https://docs.docker.com/engine/install/ubuntu/>
    - In `Installation methods` section follow `Install using the repository` method
  - Check whether your user of Ubuntu has been added to `docker` group using the following command:

  ```bash
  getent group docker | awk -F: '{print $4}'
  ```

  - If it has not been added, add it using `Manage Docker as a non-root user` section from <https://docs.docker.com/engine/install/linux-postinstall/>
  - Install Docker Compose as described at <https://docs.docker.com/compose/install/>
  - In any location that you prefer within your user home directory (`~`) execute the following command to clone DC Ledger project:

  ```bash
  git clone https://github.com/zigbee-alliance/distributed-compliance-ledger.git
  ```

  - In the the root directory of the cloned project execute the following commands to create and start DC Ledger network in Docker:

  ```bash
  make install
  make localnet_init
  make localnet_start
  ```

  - Execute the following command:

  ```bash
  dcld
  ```

  - This command runs DC Ledger CLI. Without any arguments it shows the help on the command usage
  - To stop the network execution:

  ```bash
  make localnet_stop
  ```

  - Then you can start the network again with the existing data using `make localnet_start`
  - If you need to start a new clean network then do the following steps prior to executing `make localnet_start`:
    - Remove the existing network data using `make localnet_clean`
    - Initialize the new network data using `make localnet_init`
