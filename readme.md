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

NOTE: You need to wait for your node to catch up to the current blockchain height first. Once it's caught up, you can execute the create validator command. To execute this command, you will need test kii coins from our discord faucet.

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

You will need this set up on your node server:
1) nginx
2) self-signed certificates (or purchased certificates from domain)
3) open inbound connection on port 443 and 26658 in your firewall (if you're on a cloud provider, you can normally do this through network security group)
4) signed up for a domain name for your validator

If you are using self signed certificates, you can use cerbot to automatically generate self-signed certificates for you: https://certbot.eff.org/instructions?ws=nginx&os=ubuntufocal
(NOTE: on step 6, ensure you select the `sudo certbot --nginx` option since we want to get the certificates and also install them in your nginx installation on your server).  Follow the prompts, it will eventually ask you for your domain name.  You need to provide a VALID domain name otherwise this will not work.

Once you've ensured all 4 items have been set up on your server, modify the nginx configuration file (no extension) in:
```
/etc/nginx/sites-available/default
/etc/nginx/sites-enabled/default
```
If there are no default files in those directories, create a file for each of the directory called `default`.
In both the files, add the following content:

Note: you will need to change the following:
```
YOUR_DOMAIN_WITHOUT_THE_PROTOCOL - your domain name without the protocol.  For example, mydomain.com

PATH_TO_FULLCHAIN_CERTIFICATE - path to your ssl certificates.  Ensure this is the full chain for your certificates (including root).  There are many guides out there on how to combine your certificates if they come separately)

PATH_TO_CERTIFICATE_KEY - path to your certificate private key
```

```
upstream lcd_url {
    server 127.0.0.1:1317;
}
upstream rpc_url {
    server 127.0.0.1:26657;
}

server  {
    listen                       443 ssl;
    server_name                  YOUR_DOMAIN_WITHOUT_THE_PROTOCOL;
    ssl_certificate              PATH_TO_FULLCHAIN_CERTIFICATE/fullchain.pem;
    ssl_certificate_key          PATH_TO_CERTIFICATE_KEY/privkey.pem;
    ssl_session_cache            builtin:1000  shared:SSL:10m;
    ssl_protocols                TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers                  HIGH:!aNULL:!eNULL:!EXPORT:!CAMELLIA:!DES:!MD5:!PSK:!RC4;
    ssl_prefer_server_ciphers    on;
    access_log                   /var/log/nginx/access.log;
    error_log                    /var/log/nginx/error.log;
    location / {

#                add_header Access-Control-Allow-Origin *;
        add_header Access-Control-Max-Age 3600;
        add_header Access-Control-Expose-Headers Content-Length;
        proxy_pass                          http://lcd_url;
        proxy_read_timeout                  90;
    }

}
server  {
    listen                       26658 ssl;
    server_name                  YOUR_DOMAIN_WITHOUT_THE_PROTOCOL;
    ssl_certificate              PATH_TO_FULLCHAIN_CERTIFICATE/fullchain.pem;
    ssl_certificate_key          PATH_TO_CERTIFICATE_KEY/privkey.pem;
    ssl_session_cache            builtin:1000  shared:SSL:10m;
    ssl_protocols                TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers                  HIGH:!aNULL:!eNULL:!EXPORT:!CAMELLIA:!DES:!MD5:!PSK:!RC4;
    ssl_prefer_server_ciphers    on;
    access_log                   /var/log/nginx/access.log;
    error_log                    /var/log/nginx/error.log;
    location / {
proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "upgrade";
            proxy_set_header                    Host $host;
            proxy_set_header                    X-Real-IP $remote_addr;
            proxy_set_header                    X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_http_version 1.1;
            proxy_bind                          $server_addr;
            proxy_pass                          http://rpc_url;
            proxy_read_timeout                90;
    }

}
```

Once you've set up both `default` files, save them then restart nginx `sudo systemctl restart nginx`.
Test your setup by visiting your browsers with the urls:

```
https://YOUR_DOMAIN - this should lead to a swagger API implementation of endpoints available on your node
https://YOUR_DOMAIN:26658 - this should lead to a page with all the rpc available endpoints on your node
```

You should now also see your validator in https://app.kiiglobal.io/kii/staking

## What if I can't see my validator on the explorer after enabling SSL?
To debug this issue, ensure the following:

1) you've executed the `create-validator` command without any errors and ensured that your node IS a validator
2) you've tested your domain on the browser and it shows properly
```
https://YOUR_DOMAIN - this should lead to a swagger API implementation of endpoints available on your node
https://YOUR_DOMAIN:26658 - this should lead to a page with all the rpc available endpoints on your node
```
3) your validator is not jailed.  If it is jailed, you can unjail your validator by executing the following command:
   `kiichaind tx slashing unjail --from=<your_validator_account_name> --chain-id=kiiventador`

## Run KII Chain Node Locally

```
git clone https://github.com/KiiBlockchain/kii.git
cd kii
**customize the config.yml**
ignite chain serve -r
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).
