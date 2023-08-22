# KII Chain
**KII Chain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Prerequisites
```
Node 18+
Go 1.19
Ignite CLI (https://docs.ignite.com/welcome/install)
```
## Run KII Chain Locally

```
git clone --recurse-submodules -j8 https://github.com/KiiBlockchain/kii.git
cd kii
ignite chain serve -r -v

**NOTE** The recursive flag is optional if you want to get the frontend projects.
**NOTE 2** if you would like to pull all updates including the submodules use this command:

git pull && git submodule foreach git pull origin master
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).



## Install only binary


```
curl https://get.ignite.com/KiiBlockchain/kii@latest! | sudo bash
```
`username/kii` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
