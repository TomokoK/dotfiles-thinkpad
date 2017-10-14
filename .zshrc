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

export PATH="/usr/local/bin:/usr/bin:/bin:/opt/bin:/usr/x86_64-pc-linux-gnu/gcc-bin/4.9.4:/sbin:/home/asuka/.bin"
export MPD_HOST="/home/asuka/.cache/mpd.sock"
export ANDROID_HOME="/home/asuka/.Android/sdk"
export PATH="$PATH:/opt/android-studio/gradle/gradle-3.2/bin"
#alias neofetch='neofetch --ascii_distro gentoo'
alias wgetdir='wget -r --no-parent --reject "index.html*"'
alias maim='maim ~/$(date +%s).png'
alias xbps-search='sudo xbps-query -R -s'
alias neofetch='neofetch --colors 6 0 0 6 --ascii_colors 6 0'
