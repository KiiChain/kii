# KII Chain
**KII Chain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Prerequisites
```
Node 18+
Go 1.19
bison (on the OS level)
build-essential (on the OS level)
jq (on the OS level)
Ignite CLI (https://docs.ignite.com/welcome/install)
```
## Run KII Node for Testnet
```shell
> git clone https://github.com/KiiBlockchain/kii.git
> cd kii
> nano config.yml
```

```yaml
//customize the config.yml similar to below
version: 1
accounts:
- name: name_of_your_wallet_account
  coins: [ 100ukii ] // this doesn't matter, but a value is needed for now
genesis:
  chain_id: kiiventador
  app_state:
    staking:
      params:
        bond_denom: ukii
validators:
- name: name_of_your_wallet_account
  bonded: 100ukii // this doesn't matter, but a vale is needed for now
  app:
    pruning: "nothing"
  config:
    moniker: "name_of_your_validator"
  client:
    output: "json"
```

```shell
> ignite chain init
// copy the mnemonic phrases and the address (long string with prefix "kii") from the output and make a note of that later
> cp genesis/genesis.json /home/ubuntu/.kiichain/config/genesis.json
> nano /home/ubuntu/.kiichain/config/config.toml
```
## Retrieve Node ID Peer
To connect to a peer on the network, you will need to get a peer's node id.  Any validator's node id that is connected to the network will do but if you don't have access to one, you can use one of KII's master node validator to connect to the network:

Visit: https://a.testnet.kiivalidator.com:26658/status

Extract the value from data object:
```
result.node_info.id
```

Then modify the config.toml file with the node id value

```toml
// search for persistent_peers = ""
// replace it with persistent_peers = "<node id>@3.129.207.228:26656"
// save the file
```

```shell
> kiichaind start
```

## Convert node to validator
Once you have the node running with the above instructions, you need to broadcast a transaction on the chain signaling you want to be a validator node.  To become a validator, you will need to have a minimum stake of 1 tkii.

Reach out to the Kii Team on discord and request for some test tkii tokens sent to your address (the generated address from the previous step).

Once you have the tokens in your address, in a separate terminal on the machine you're running your node on, execute the following command:

```
kiichaind tx staking create-validator \
  --amount=1tkii \
  --pubkey=$(kiichaind tendermint show-validator) \
  --moniker="<your validator name>" \
  --commission-rate=0.1 \
  --commission-max-rate=0.2 \
  --commission-max-change-rate=0.01 \
  --min-self-delegation=1 \
  --gas=auto --gas-adjustment=1.2 \
  --gas-prices=10.0tkii \
  --from <name of your wallet account>
```

Once the command is broadcasted successfully, the node should now be classified as a validator.

## Have your validator show up in the block explorer
For your validator to show up on the block explorer: https://app.kiiglobal.io/kii/staking

You will need to assign a domain name to your validator and ensure you enable HTTPS traffic for both port 1317 and 26657.
See more information about setting up app.toml and config.toml here: https://github.com/ping-pub/explorer/blob/master/installation.md

## Run KII Chain Node Locally

```
git clone https://github.com/KiiBlockchain/kii.git
cd kii
**customize the config.yml**
ignite chain serve -r
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).