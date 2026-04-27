# Node OS Upgrade: Upgrading DCL cloud instance from Ubuntu 20.04 to Ubuntu 24.04

This document outlines the procedure for upgrading the active Distributed Compliance Ledger (DCL) node to Ubuntu 24.04.

1.  **Take a snapshot**: Take a snapshot from a running instance. Follow-up guidelines of your cloud provider to take a snapshot.

    Bellow are the steps for AWS Lightsail.

    1.1 Open `AWS Lightsail` console and select running instance.

    1.2 From bottom navigation bar, select `Snapshots` tab.

    1.3 Click `Create snapshot` and give a name then click to `Create`.

    1.4 Wait until Snapshotting finishes.

2. **Connect using SSH**. Below are the steps for AWS Lightsail Browser Console. You can also use an SSH client through Terminal on Linux/macOS or PuTTY on Windows.

   2.1 Open `AWS Lightsail` console and select running instance.

   2.2 From bottom navigation bar, select `Connect` tab.

   2.3 Click on `Connect using SSH`.

3.  **OS Preparation**: Update existing packages.

    3.1 Switch to root user
    ```bash
    sudo su
    ```
    3.2 Update existing packages.
    ```bash
    sudo apt update 
    sudo apt upgrade -y
    sudo apt autoremove -y
    ```
    *Note: During the upgrade, this terminal may fail. In that case, you have to use an SSH client like Terminal on Linux or PuTTY on Windows to connect to the server.*

    3.3 After package upgrade, reboot the instance
    ```bash
    reboot
    ```

4.  **Release Upgrade**: Initiate the Ubuntu release upgrade.
    
    4.1 Connect to the instance using SSH.
    ```bash
    ssh -i "private-key.pem" root@{{ip or hostname}}
    ```

    4.2 Execute release upgrade command
    ```bash
    do-release-upgrade 
    ```
    4.3 During the upgrade, it will ask to open an additional ssh daemon at `1022` port. Enter `y` to continue.
 
    4.4 It will also ask to open the port in the firewall. This is a firewall inside the instance. Press `Enter` button to continue
        Let's open the port in the AWS Lightsail firewall.

      * From bottom navigation bar of selected instance, select `Networking` tab.
      * Click `Add rule`
      * Default rule is `Custom` and protocol is `TCP` and enter `1022` port
      * Click on `Create`

    *Note: It is a reserved port to connect to the instance if the default `22` port will fail during the upgrade. Remember to remove this open port after upgrade.

    4.5 For `Installing the upgrade can take several hours` prompt, press `y` button to continue.

    4.6 While the upgrade is downloading these packages, connect to `1022` SSH port as a second safe connection.
    ```bash
    ssh -p 1022 -i "private-key.pem" root@{ip or hostname}
    ```

    4.7 For `There are services installed on your system which need to be restarted when certain libraries such as libpam, libc,.. Restart services during package upgrade without asking` prompt, select `yes` and press Enter.

    4.8 For `Postix Configuration` prompt, select default `No configuration` and press Enter.

    4.9 For `Configuration file 'etc/monit/monitrc` prompt, press `Y`.

    4.10 For `Remove obsolete packages` prompt, enter `y`.

    4.11 For `System upgrade complete. Restart required` prompt, enter `y` to reboot the instance.

5.  **Verification**: After reboot, ensure the instance upgraded  to Ubuntu `22.04` and the `dcld` is running. 

    5.1 After connecting using ssh, execute below command to verify current version of Ubuntu:
    ```bash
    sudo lsb_release -a
    ```
    
    5.2 Verify that the node is running.
    ```bash
    dcld status
    ```
    
   6. **Upgrade to 24.04**. To upgrade to next `24.04` release, please repeat steps from 2 to 5, you can skip `1`(Taking snapshot) and `4.4`(Opening `1022` port on instance) steps.
7. **Post upgrade**. Remove the firewall rule added in step 4.4 from your Network settings of your instance.