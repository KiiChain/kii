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
cp /home/ubuntu/.kiichain/config/genesis.json genesis/genesis.json
> cp genesis/genesis.json /home/ubuntu/.kiichain/config/genesis.json
> nano /home/ubuntu/.kiichain/config/config.toml
```

```toml
// search for persistent_peers = ""
// replace it with persistent_peers = "ad4379b36ab6e13fc2604d9d4de9e280385141a1@3.129.207.228:26656"
// save the file
```

```shell
> kiichaind start
```

## Convert node to validator
Once you have the node running with the above instructions, you need to broadcast a transaction on the chain signaling you want to be a validator node.  To become a validator, you will need to have a minimum stake of 1 ukii.

Reach out to the Kii Team on discord and request for some test ukii tokens sent to your address (the generated address from the previous step).

Once you have the tokens in your address, in a separate terminal on the machine you're running your node on, execute the following command:

```
kiichaind tx staking create-validator \
  --amount=1tkii \
  --pubkey=$(kiichaind tendermint show-validator) \
  --moniker="KiiPagani" \
  --commission-rate=0.1 \
  --commission-max-rate=0.2 \
  --commission-max-change-rate=0.01 \
  --min-self-delegation=1 \
  --gas=auto --gas-adjustment=1.2 \
  --gas-prices=0.01tkii \
  --from kiipaganiwallet
```

Once the command is broadcasted successfully, the node should now be classified as a validator.

## Run KII Chain Node Locally

```
git clone https://github.com/KiiBlockchain/kii.git
cd kii
**customize the config.yml**
ignite chain serve -r
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).