# Ledger Nano Support

## Prepare Ledger Nano Device to Use with DCL

### Step 1: Install Ledger Live

Download and install Ledger Live: Go to the [Ledger Live download page](https://www.ledger.com/ledger-live) and select the version suitable for your operating system. Follow the on-screen instructions to install it.

<!-- markdown-link-check-disable -->
### Step 2: Set Up Your Ledger Device

Initialize Your Ledger Device: If this is your first time using your Ledger, you will need to configure it. Follow the instructions to choose a PIN and securely write down the mnemonic seed phrase that appears. This seed is crucial for recovery and is compatible with DCL. Detailed instructions are available [here](https://support.ledger.com/hc/en-us/articles/4416927988625-Set-up-your-Ledger-Nano-S-Plus?docs=true).

### Step 3: Install the Cosmos (ATOM) app on your Ledger device

To use the Ledger Nano with DCL, you first need to install the Cosmos (ATOM) application. Connect your Ledger device to your computer, open Ledger Live, and navigate to the 'Manager' tab. Allow the manager on your Ledger device, find the Cosmos (ATOM) app in the catalog, and click 'Install'. For detailed instructions, visit the [Cosmos (ATOM) documentation page](https://support.ledger.com/hc/en-us/articles/360013713840-Cosmos-ATOM?docs=true).
<!-- markdown-link-check-enable -->

## Using Ledger Device with DCL CLI

To use the DCL (dcld) command line interface, it must first be installed on your computer. You can find the installation instructions in the [Quick start guide](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/50ef77243b49764f474e545cc4be2beee4793ed0/docs/quickStartGuide.adoc#L3).

### Add your Ledger Key

1. Connect and unlock your Ledger device.
2. Open the Cosmos app on your Ledger.
3. Add a new account in dcld using your Ledger key. Enter the following command:

    ```bash
    dcld keys add <keyName> --ledger
    ```

4. Approve the addition of the account on your Ledger device within the Cosmos app.
5. Verify the account has been added by using the following command:

    ```bash
    dcld keys show <keyName>
    ```

### Signing and Sending Transactions

To sign and send transactions using your Ledger Nano device, follow these steps:

1. Ensure your Ledger device is unlocked and that the Cosmos app is open.
2. Initiate the transaction using the DCL CLI. Replace <transaction_command> with the specific command for your transaction:

    ```bash
    dcld <transaction_command> --from <keyName> --ledger
    ```

3. Review and confirm the transaction details on your Ledger device. Carefully inspect each aspect of the transaction JSON displayed on the screen.

## Using DCL with a Web Browser

This section explains how to use a Ledger device to sign DCL transactions via the web UI using the Keplr wallet.

- Install the Keplr Wallet as a browser extension from the official Keplr website, available for Chrome or other supported browsers.
- Connect your Ledger device to your computer and ensure the Cosmos app is open on the device.
- Click on the Keplr extension icon and select "Connect Hardware Wallet." If you have previously created a wallet in Keplr, you will need to add a new wallet before selecting "Connect Hardware Wallet" again.
- Confirm that your Ledger device is unlocked with the Cosmos app active. Then, follow the on-screen instructions provided by the Keplr pop-up.

Now, you are ready to sign transactions using your Ledger Nano and the Keplr wallet. Each time you initiate a transaction, you will need to confirm it on your Ledger device following prompts from the Keplr interface.
