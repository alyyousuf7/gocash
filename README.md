# GoCash
A small utility to manage petty cash or a wallet.

## Install
Simply run `go install github.com/alyyousuf7/gocash/cmd/gocash`.

## Usage
Global help
```bash
$ gocash help
Petty cash summary

Usage:
  gocash [flags]
  gocash [command]

Available Commands:
  credit      Add a credit transaction
  debit       Add a debit transaction
  help        Help about any command
  rm          Remove a transaction

Flags:
  -c, --config string   Config file path (default "/home/ali/.gocash.yml")
  -h, --help            help for gocash

Use "gocash [command] --help" for more information about a command.
```

---

To record a new transaction you can use `debit` or `credit` command:
```bash
$ gocash credit 50 "7up"
Transaction added.
Your new balance: Rs. -50
```

or

```bash
$ gocash debit 50 "Paid"
Transaction added.
Your new balance: Rs. 0
```

---

If you mistakenly add a transaction, it can be removed using the `rm` command and providing the numeric transaction ID:
```bash
$ gocash rm 10
Transaction removed.
Your new balance: Rs. -50
```

---

Get all transaction history
```bash
$ gocash
  ID |  DATE  |   NOTE   |  AMOUNT    
+----+--------+----------+-----------+
   1 | Sep 07 | Biscuits | - Rs. 30   
   2 | Sep 08 | 7up      | - Rs. 50   
   3 | Sep 11 | Mirinda  | - Rs. 50   
   4 | Sep 13 | Paid     | + Rs. 230  
   5 | Sep 21 | 7up      | - Rs. 50   
   6 | Sep 22 | Mirinda  | - Rs. 50   
   7 | Oct 02 | 7up      | - Rs. 50   
   8 | Oct 11 | Paid     | + Rs. 50   
   9 | Oct 11 | 7up      | - Rs. 50   
+----+--------+----------+-----------+
                 TOTAL   | - RS  50   
              +----------+-----------+
``` 

## Configuration
GoCash loads configuration from `~/.gocash.yml` file. Configuration file stores information about which database to use.

GoCash supports SQLite and BoltDB, you have to choice to select one of them, but by default it uses SQLite to store all the transactions.

If configuration file doesn't exist, it initiates a new YAML file with default values which looks like this:

```yaml
storage: sqlite
storage-config:
  file: /home/ali/gocash.sqlite
```

To switch to BoltDB, `storage` value can be changed to `boltdb` and appropriate `storage-config.file`.

## Multiple wallets
Multiple wallets can be maintained by creating multiple configuration files and access each of them by providing relevant configuration file using `--config` global flag in each command.

Example:
```bash
$ gocash --config $HOME/.wallet1.yml credit 50 Lays
$ gocash --config $HOME/.bank.yml credit 100 Transport
``` 

You can create aliases in your shell for each wallet to make life easier. :)

## License
MIT