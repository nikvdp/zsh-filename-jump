# zsh-filename-jump

small go utility and accompanying zsh glue to let you use B and W keys to jump
by "shell words" (read: filenames/paths) in zsh's vi mode

see: https://twitter.com/ArghZero/status/1550719990453911552 for original idea

# status 

weekend hack, it works on my machine! ymmv

# install

- `go build` (you'll need golang installed)
- put the new file somewhere on your path (`sudo cp zsh-filename-jump
  /usr/local/bin` should do the trick)
- do a `source zsh/vanilla.zsh` (for best results do this in your zshrc)
  - (if you're using [zsh-vi-mode](https://github.com/jeffreytse/zsh-vi-mode) use
    `source zsh/zsh-vi-mode.zsh` instead)

# todo

- [ ] clean this up and figure out how to package it as an actual zsh plugin
