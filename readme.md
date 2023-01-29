# chat
**chat** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Solution Explanation

The message text data consists of 3 field:
- Sender
- Receiver
- Body

### CLI Commands and Field

- The command to send a text is 
```
send-text [address] [body]
```
When sending message, the user can specify the address of the receiver (or any string, hence enabling group chat).

An example of a tx sending text looks like this:
```
chatd tx chat send-text cosmos1k5l32jxsxsu7c90cnyq7shrxlgq6ddha45g0x2 "hi bob" --from alice
```
- The command to list all messages sent to their wallets/group
```
get-received [address]
```
- The command to list all messages sent from their wallets
```
get-sent [address]
```

### logic of the methods

When a user send the data, the text data that consists of the sender, receiver and text body is added to the store.

During the retrieval of the text, we can query the store and match the sender or receiver depending on the method.

### Example run
set up
```
limkaraik@Lims-MacBook-Pro chat % ignite c serve                     
  Blockchain is running
  
  👤 alice's account address: cosmos1jax77wjnvyuffrjlc204zrad8m7p63tx5ndyd9
  👤 bob's account address: cosmos1k5l32jxsxsu7c90cnyq7shrxlgq6ddha45g0x2
  
  🌍 Tendermint node: http://0.0.0.0:26657
  🌍 Blockchain API: http://0.0.0.0:1317
  🌍 Token faucet: http://0.0.0.0:4500
```

send messages
```
chatd tx chat send-text cosmos1k5l32jxsxsu7c90cnyq7shrxlgq6ddha45g0x2 "hi bob" --from alice

```
```
chatd tx chat send-text cosmos1jax77wjnvyuffrjlc204zrad8m7p63tx5ndyd9 "hi alice" --from bob
```

retrieve messages
```
limkaraik@Lims-MacBook-Pro chat % chatd q chat get-sent cosmos1jax77wjnvyuffrjlc204zrad8m7p63tx5ndyd9                                              
pagination:
  next_key: null
  total: "2"
text:
- body: hi bob
  id: "0"
  receiver: cosmos1k5l32jxsxsu7c90cnyq7shrxlgq6ddha45g0x2
  sender: cosmos1jax77wjnvyuffrjlc204zrad8m7p63tx5ndyd9
```
```
limkaraik@Lims-MacBook-Pro chat % chatd q chat get-received cosmos1jax77wjnvyuffrjlc204zrad8m7p63tx5ndyd9
pagination:
  next_key: null
  total: "2"
text:
- body: hi alice
  id: "1"
  receiver: cosmos1jax77wjnvyuffrjlc204zrad8m7p63tx5ndyd9
  sender: cosmos1k5l32jxsxsu7c90cnyq7shrxlgq6ddha45g0x2
```





