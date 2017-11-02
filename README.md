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

% gtrello -h
Usage of gtrello:
  -c string
        config file path (default "$HOME/.config/gTrello/config.toml")
  -d    day report (default true)
  -o string
        output dir path (default "$HOME/.config/gTrello/output")
  -t string
        template file path (default "$HOME/.config/gTrello/template/template.md")
  -w    week report
```

Config file
```
% tree ~/.config/gTrello
/Users/soichiro-taga/.config/gTrello
├── config.toml
├── output
│   ├── day
│   │   └── 2017-11-02.md
│   └── week
│       └── 2017-11-02.md
└── template
    └── template.md
```
