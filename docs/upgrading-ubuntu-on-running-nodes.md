# Node OS Upgrade: Upgrading DCL Cloud Instance from Ubuntu 20.04 to Ubuntu 24.04

This document outlines the procedure for upgrading an active Distributed Compliance Ledger (DCL) node from Ubuntu 20.04 to Ubuntu 24.04.

> **Note:** The upgrade from 20.04 to 24.04 is a two-step process: first to 22.04, then to 24.04.

> **Provider scope:** The release-upgrade procedure itself (steps 3–6) is provider-agnostic. The provider-specific actions are **snapshot** (step 1), **SSH connection** (step 2), and **opening port 1022** (step 4.4). Expand the section for your cloud provider in each step.

1.  **Take a Snapshot**: Before proceeding, take a snapshot of your running instance so you can roll back if anything goes wrong.

    <details>
    <summary> AWS Lightsail </summary>

    1.1 Open the `AWS Lightsail` console and select your running instance.

    1.2 From the bottom navigation bar, select the `Snapshots` tab.

    1.3 Click `Create snapshot`, provide a name, and click `Create`.

    1.4 Wait until the snapshotting process finishes.

    </details>

    <details>
    <summary> AWS EC2 </summary>

    1.1 Open the EC2 console and go to **Elastic Block Store → Snapshots**.

    1.2 Choose **Create snapshot**. For **Resource type**, select **Volume** (or **Instance** to snapshot all attached volumes at once).

    1.3 Select the target volume (the instance's root volume), add a description/tags, and choose **Create snapshot**.

    Please see also AWS [docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-creating-snapshot.html).

    </details>

    <details>
    <summary> GCP </summary>

    1.1 In the Cloud Console, go to **Compute Engine → Snapshots**.

    1.2 Click **Create snapshot**.

    1.3 Enter a name; under **Source disk**, select the VM's boot disk.

    1.4 Choose location (regional / multi-regional) and click **Create**.

    Please see also GCP [docs](https://cloud.google.com/compute/docs/disks/create-snapshots).

    </details>

    <details>
    <summary> Azure </summary>

    1.1 Azure Portal → **Virtual machines** → select the VM → **Disks**.

    1.2 Click the **OS disk** name to open the managed disk blade.

    1.3 Select **Create snapshot** from the top command bar.

    1.4 Provide name, resource group, and region; set **Snapshot type = Incremental** (recommended). Click **Review + create**.

    Please see also Azure [docs](https://learn.microsoft.com/azure/virtual-machines/snapshot-copy-managed-disk).

    </details>

2. **Connect Using SSH**: Connect to your instance via SSH. You can use your cloud provider's browser console or a standard SSH client (Terminal on Linux/macOS, PuTTY on Windows).

    <details>
    <summary> AWS Lightsail </summary>

    2.1 Open the `AWS Lightsail` console and select your running instance.

    2.2 From the bottom navigation bar, select the `Connect` tab.

    2.3 Click on `Connect using SSH`.

    </details>

    <details>
    <summary> AWS EC2 </summary>

    Pick one of:

    2.1 **EC2 Instance Connect (browser):** EC2 console → **Instances** → select instance → **Connect** → **EC2 Instance Connect** tab → **Connect**.

    2.2 **Session Manager:** requires SSM Agent on the instance and an attached IAM role with `AmazonSSMManagedInstanceCore`. Console → **Connect** → **Session Manager** tab → **Connect**.

    2.3 **Local SSH client** (default user is `ubuntu` for Ubuntu AMIs; key file must have `chmod 400`):
    ```bash
    ssh -i /path/to/key.pem ubuntu@<public-dns-or-ip>
    ```

    Please see also AWS [docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/connect-to-linux-instance.html).

    </details>

    <details>
    <summary> GCP </summary>

    Pick one of:

    2.1 **Cloud Console SSH (browser):** **Compute Engine → VM instances** → click **SSH** next to the instance (uses OS Login or metadata-managed keys automatically).

    2.2 **gcloud CLI:**
    ```bash
    gcloud compute ssh INSTANCE_NAME --zone=ZONE
    ```

    2.3 **Local SSH client** (after adding your public key via OS Login or instance/project metadata):
    ```bash
    ssh -i PATH_TO_PRIVATE_KEY USERNAME@EXTERNAL_IP
    ```
    With OS Login, `USERNAME` is derived from your Google account (e.g., `user_example_com`). With metadata SSH keys, `USERNAME` is the user you specified when uploading the key. GCP Ubuntu images do **not** use a fixed `ubuntu` default.

    Please see also GCP [docs](https://cloud.google.com/compute/docs/instances/ssh).

    </details>

    <details>
    <summary> Azure </summary>

    Pick one of:

    2.1 **Azure Bastion (browser):** VM → **Connect** → **Bastion** → enter username / SSH key (or password) → **Connect**. No public IP or local SSH client required.

    2.2 **Azure CLI with Microsoft Entra ID** (requires the AAD-login VM extension and the *Virtual Machine Administrator/User Login* role):
    ```bash
    az ssh vm --name <vm-name> --resource-group <rg>
    ```

    2.3 **Local SSH client:**
    ```bash
    ssh -i ~/.ssh/<private-key> <admin-username>@<public-ip-or-fqdn>
    ```

    > **_Note:_** Azure Ubuntu Marketplace images do **not** have a fixed default user (unlike AWS's `ubuntu`); the username is whatever was entered as the administrator account during VM creation.

    Please see also Azure [docs](https://learn.microsoft.com/azure/virtual-machines/linux-vm-connect).

    </details>

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
    
    4.1 Connect to the instance using SSH (replace `{{user}}` with your instance's SSH user — `ubuntu` is the default for AWS Ubuntu AMIs):
    ```bash
    ssh -i "private-key.pem" {{user}}@{{ip or hostname}}
    ```

    4.2 Start the release upgrade:
    ```bash
    sudo do-release-upgrade 
    ```

    4.3 During the upgrade, you will be asked to open an additional SSH daemon on port `1022`. Enter `y` to continue.
 
    4.4 You will also be asked to open the port in the firewall. This refers to the OS-level firewall. Press `Enter` to continue. 
    
    **Important:** You must also open port `1022` in your cloud provider's network firewall.

    <details>
    <summary> AWS Lightsail </summary>

    4.4.1 From the `Networking` tab of your instance, click `Add rule`.

    4.4.2 Select `Custom` application, `TCP` protocol, and enter `1022` as the port.

    4.4.3 Click `Create`.

    </details>

    <details>
    <summary> AWS EC2 </summary>

    4.4.1 EC2 console → **Instances** → select the instance → **Security** tab.

    4.4.2 Click the linked **Security group** ID.

    4.4.3 Choose **Edit inbound rules** → **Add rule**.

    4.4.4 Set **Type** = *Custom TCP*, **Port range** = `1022`, **Source** = *My IP* (or a specific CIDR).

    4.4.5 Choose **Save rules**.

    Please see also AWS [docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/working-with-security-groups.html#adding-security-group-rule).

    </details>

    <details>
    <summary> GCP </summary>

    4.4.1 **VPC network → Firewall → Create firewall rule.**

    4.4.2 Set **Network** to the VM's VPC; **Direction** = *Ingress*; **Action** = *Allow*.

    4.4.3 **Targets:** Specified target tags (e.g., `ssh-1022`).

    4.4.4 **Source IPv4 ranges:** your CIDR (e.g., `203.0.113.4/32`).

    4.4.5 **Protocols and ports:** Specified, TCP = `1022`. Click **Create**.

    4.4.6 Attach the tag to the VM: **Compute Engine → VM instances** → edit the VM → add `ssh-1022` under **Network tags** → **Save**.

    Please see also GCP [docs](https://cloud.google.com/firewall/docs/using-firewalls).

    </details>

    <details>
    <summary> Azure </summary>

    4.4.1 Azure Portal → the VM → **Networking** (or **Networking → Network settings**).

    4.4.2 **Inbound port rules** → **Add inbound port rule**.

    4.4.3 Set: **Source** = *My IP address* (or specific CIDR); **Source port ranges** = `*`; **Destination** = *Any*; **Service** = *Custom*; **Destination port ranges** = `1022`; **Protocol** = *TCP*; **Action** = *Allow*; **Priority** = `1010`; **Name** = `AllowSSHFallback`.

    4.4.4 Click **Add**.

    Please see also Azure [docs](https://learn.microsoft.com/azure/virtual-network/manage-network-security-group).

    </details>

    *Note: Port 1022 is a fallback to ensure you can connect if the default SSH port (22) fails during the upgrade. Remember to remove this rule after the upgrade is complete.*

    4.5 When prompted "Installing the upgrade can take several hours", press `y` to continue.

    4.6 While the upgrade is downloading packages, it is recommended to connect to the `1022` SSH port as a backup connection:
    ```bash
    ssh -p 1022 -i "private-key.pem" {{user}}@{{ip or hostname}}
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