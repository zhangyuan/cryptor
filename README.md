# cryptor

Encrypt a file with password and AES (GCM or CFB mode).

> This is a personal project that is in its very early stage and there may be breaking changes at any time.

## Usage

```
$ cryptor --help
NAME:
   cryptor - encrypt and decrypt file

USAGE:
   cryptor [global options] command [command options]

COMMANDS:
   encrypt, e  Encrypt a file
   decrypt, d  Decrypt a file
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### Encrypt a file

```bash
cryptor encrypt -p a -s b -m gcm -i tests/fixtures/data.txt -o tmp/tests/encrypted.gcm.txt
```

### Decrypt a file

```bash
cryptor decrypt -p a -s b -m gcm -i tmp/tests/encrypted.gcm.txt -o tmp/tests/decrypted.gcm.txt
```