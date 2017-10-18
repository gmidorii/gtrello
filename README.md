# gtrello

## Overview
gtrello create daily report from Trello cards.

## Usage
Install
```sh
% go get github.com/midorigreen/gtrello
```

Build
```sh
% go build
```

Run
```sh
% ./gtrello

% ./gtello -h
Usage of ./gtrello:
  -c string
        config file path (default "/Users/midori/.config/gTrello/config.toml")
  -o string
        output dir path (default "/Users/midori/.config/gTrello/output")
  -t string
        template file path (default "/Users/midori/.config/gTrello/template/template.md")
```

Config file
```
% tree ~/.config/gtrello
$HOME/.config/gtrello
├── config.toml
├── output
│   ├── 2017-10-15-gtrello.md
│   └── ...
└── template
    └── template.md
```
