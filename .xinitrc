#!/bin/usr
#
# Tomoko's xinitrc file

xsetroot -cursor_name left_ptr &

# Start Programs Here
mpd & compton --config ~/.config/compton.conf & dunst & mpdscribble &
feh --bg-fill ~/Pictures/Wallpapers/6.jpg &

# Start bspwm
sxhkd &
exec bspwm -c /home/asuka/.config/bspwm/bspwmrc

# Start i3
#exec i3

# Start Openbox
#exec openbox-session
