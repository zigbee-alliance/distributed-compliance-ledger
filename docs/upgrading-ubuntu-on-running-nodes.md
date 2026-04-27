# Node OS Upgrade: Upgrading DCL Cloud Instance from Ubuntu 20.04 to Ubuntu 24.04

This document outlines the procedure for upgrading an active Distributed Compliance Ledger (DCL) node from Ubuntu 20.04 to Ubuntu 24.04.

> **Note:** The upgrade from 20.04 to 24.04 is a two-step process: first to 22.04, then to 24.04.

1.  **Take a Snapshot**: Before proceeding, take a snapshot of your running instance. Follow the guidelines of your cloud provider.

    Below are the steps for AWS Lightsail:

    1.1 Open the `AWS Lightsail` console and select your running instance.

    1.2 From the bottom navigation bar, select the `Snapshots` tab.

    1.3 Click `Create snapshot`, provide a name, and click `Create`.

    1.4 Wait until the snapshotting process finishes.

2. **Connect Using SSH**: Connect to your instance via SSH. You can use the AWS Lightsail Browser Console, or an SSH client (Terminal on Linux/macOS, PuTTY on Windows).

   2.1 Open the `AWS Lightsail` console and select your running instance.

   2.2 From the bottom navigation bar, select the `Connect` tab.

   2.3 Click on `Connect using SSH`.

3.  **OS Preparation**: Update existing packages.

    3.1 Switch to the root user:
    ```bash
    sudo -i
    ```

    3.2 Update existing packages:
    ```bash
    apt update 
    apt upgrade -y
    apt autoremove -y
    ```
    *Note: During the upgrade, your SSH session may disconnect. If that happens, reconnect using your SSH client.*

    3.3 After the package upgrade, reboot the instance:
    ```bash
    reboot
    ```

4.  **First Release Upgrade (20.04 to 22.04)**: Initiate the upgrade to Ubuntu 22.04.
    
    4.1 Connect to the instance using SSH:
    ```bash
    ssh -i "private-key.pem" ubuntu@{{ip or hostname}}
    ```

    4.2 Start the release upgrade:
    ```bash
    sudo do-release-upgrade 
    ```

    4.3 During the upgrade, you will be asked to open an additional SSH daemon on port `1022`. Enter `y` to continue.
 
    4.4 You will also be asked to open the port in the firewall. This refers to the OS-level firewall. Press `Enter` to continue. 
    
    **Important:** You must also open port `1022` in the AWS Lightsail firewall:
      * From the `Networking` tab of your instance, click `Add rule`.
      * Select `Custom` application, `TCP` protocol, and enter `1022` as the port.
      * Click `Create`.

    *Note: Port 1022 is a fallback to ensure you can connect if the default SSH port (22) fails during the upgrade. Remember to remove this rule after the upgrade is complete.*

    4.5 When prompted "Installing the upgrade can take several hours", press `y` to continue.

    4.6 While the upgrade is downloading packages, it is recommended to connect to the `1022` SSH port as a backup connection:
    ```bash
    ssh -p 1022 -i "private-key.pem" ubuntu@{ip or hostname}
    ```

    4.7 For the prompt regarding restarting services (libpam, libc, etc.), select `Yes` and press `Enter`.

    4.8 For the `Postfix Configuration` prompt, select `No configuration` and press `Enter`.

    4.9 For the `Configuration file 'etc/monit/monitrc'` prompt, press `Y` to use the new version (or your preferred choice if you have custom configurations).

    4.10 For the `Remove obsolete packages` prompt, enter `y`.

    4.11 When prompted "System upgrade complete. Restart required", enter `y` to reboot.

5.  **Verification (22.04)**: After rebooting, ensure the instance is running Ubuntu 22.04 and `dcld` is functional. 

    5.1 Verify the Ubuntu version:
    ```bash
    lsb_release -a
    ```
    
    5.2 Verify that the `dcld` node is running:
    ```bash
    dcld status
    ```
    
6. **Second Release Upgrade (22.04 to 24.04)**: To upgrade to the final `24.04` release, repeat steps 2 through 5.
   - You can skip step 1 (Snapshot).
   - You can skip step 4.4 about opening new firewall rule for port `1022`.

7. **Post-Upgrade Cleanup**: Remove the firewall rule for port `1022` from your Cloud Provider's network settings.