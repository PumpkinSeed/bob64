# bob64
Interactive cli base64 encoder/decoder

### Installation

```
    go install github.com/PumpkinSeed/bob64
```

### Usage

**Encoder**

```
    $ bob64
    bob64 - Interactive base64 encoder
    To leave the interactive mode type 'exit'
    ---
    -> test123
    dGVzdDEyMw==
    -> exit
    Bye :)
```

**Decoder**

```
    $ bob64 -d
    bob64 - Interactive base64 decoder
    To leave the interactive mode type 'exit'
    ---
    -> dGVzdDEyMw==
    test123
    -> exit
    Bye :)
```
