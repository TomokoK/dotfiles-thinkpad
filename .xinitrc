#!/bin/usr
#
#Tomoko's .xinitrc

exec /home/asuka/.screencheck.sh &

xsetroot -cursor_name left_ptr &
xset +fp /home/asuka/.fonts &
xset fp rehash &
#Start programs here
mpd &

#Start WM
exec openbox-session
