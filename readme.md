# KII Chain
**KII Chain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Prerequisites
```
Node 18+

Go 1.19

bison (on the OS level)

build-essential (on the OS level)

jq (on the OS level)

Ignite CLI v0.27.1
(curl https://get.ignite.com/cli@v0.27.1! | bash)
```

## Testnet Validator Setup (Written Instructions)
Read more [here](./install.md)

## Bring up Testnet Node (not a validator yet) - Step 1 DEMO
```
Note: it is recommended to have a minimum of 100GB of root storage.  Preferred amount is 500GB-1TB of storage.  As the blockchain grows, node operators are responsible for determining the need of expanding their server storage to prevent insufficient storage space.

2nd Note: before you start, it is recommended that as soon as you ssh into your server, start a separate session using the "screen" command.  More info on the screen command here: https://linuxize.com/post/how-to-use-linux-screen/
```

[![IMAGE ALT TEXT](http://img.youtube.com/vi/k4cFlFxU6nE/0.jpg)](http://www.youtube.com/watch?v=k4cFlFxU6nE "KII Chain Testnet Node Setup - Step 1")

## Convert Node into Validator - Step 2 DEMO
```
Note: After step 1 has been completed, ensure the node has caught up to the current height of the blockchain.  You will notice that the logs for the node is a lot slower.  You can also check the current block height at: https://app.kiiglobal.io/.

2nd Note: Ensure you have enough KII coins in the validator wallet that you have created in Step 1.  It is recommended to NOT use your personal KII wallet.  Transfer KII coins from your personal wallet to your newly created Validator wallet from Step 1.
```
[![IMAGE ALT TEXT](http://img.youtube.com/vi/Vt0u9LdYz6I/0.jpg)](http://www.youtube.com/watch?v=Vt0u9LdYz6I "Convert Node into Validator - Step 2")

## (OPTIONAL but HIGHLY recommended) Making your Validator Secured and Visible - Step 3 DEMO
```
Note: Once your validator is up and running, we need it to be visible and secured on https://app.kiiglobal.io/kii/staking.  For this, it is required for the validator to have SSL enabled.

Requirements:
1) Register a domain name.  For example: myvalidator.com
2) get certificates for your domain OR generate self-signed certificates (shown in the video)
3) install NGINX on the node server
4) open inbound connection on port 443 and 26658 in your firewall (if you're on a cloud provider, you can normally do this through network security group)
```
[![IMAGE ALT TEXT](http://img.youtube.com/vi/dsgySMaGJAw/0.jpg)](http://www.youtube.com/watch?v=dsgySMaGJAw "Enable SSL on your node - Step 3")
