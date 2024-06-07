# mark-this-date

Mark this date with a message.

## Installation

```sh
% go install github.com/luka-hash/mark-this-date@latest
% # or, more manual approach
% git clone github.com/luka-hash/mark-this-date
% cd mark-this-date
% go build .
% mv mark-this-date "$XDG_BIN_HOME"/,mark-this-date
```

## Usage

```sh
% # first set env variable (using fish shell)
% set -gx MARKTHISDATEFILE="$HOME"/somewhere/there/history
% ,mark-this-date first message
```

And messages you wanted to mark will be in appropriate file, with date and everything.

## Licence

This code is licensed under the terms of the MIT licence (see LICENCE for details).
