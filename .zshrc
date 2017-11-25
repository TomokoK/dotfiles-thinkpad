# Created by newuser for 5.2
# Lines configured by zsh-newuser-install
HISTFILE=~/.histfile
HISTSIZE=1000
SAVEHIST=1000
bindkey -e
# End of lines configured by zsh-newuser-install
# The following lines were added by compinstall
zstyle :compinstall filename '/home/asuka/.zshrc'

autoload -Uz compinit
compinit
# End of lines added by compinstall

export PATH="$PATH:/home/asuka/.bin"
export MPD_HOST="/home/asuka/.cache/mpd.sock"
alias wgetdir='wget -r --no-parent --reject "index.html*"'
alias maim='maim ~/$(date +%s).png'
