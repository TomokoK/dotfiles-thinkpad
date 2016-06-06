#!/bin/usr
#
# Tomoko's .xinitrc file

xsetroot -cursor_name left_ptr &

# Start stuff here
mpd  & transmission-daemon &

# Start Openbox

exec openbox-session
