# 🦆 `duck-gen`

Generate [**DuckDuckGo Email Protection**](https://spreadprivacy.com/protect-your-inbox-with-duckduckgo-email-protection/) addresses from the command-line.

```console
foo@bar:~$ duck-gen
Enter your Duck address: johnny@duck.com
Enter the one-time passphrase sent to your email: heftiness slate captivate cornmeal 
bzpkash3@duck.com

foo@bar:~$ duck-gen
embargo-frisk-rack@duck.com
```

`duck-gen` will initiate a login flow the first time you run it, but will save your authentication tokens at `~/.duck_token` for future runs. 

## Installation 

Download pre-built binaries from the [Releases](https://github.com/chowder/duck-gen/releases) page

## Usage
If you are logged in, you get four options:

```
Usage: duck-gen [generate|history|verbose_history|generate_all]
```

- `generate` makes an alias.
- `history` shows last generated aliases without timestamp.
- `verbose_history` shows history with timestamp.
- `generate_all` makes 30 aliases.
