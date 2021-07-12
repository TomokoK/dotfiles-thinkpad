#!/bin/usr
#
#Tomoko's .xinitrc

export GTK_IM_MODULE=ibus
export XMODIFIERS=@im=ibus
export QT_IM_MODULE=ibus
xrdb /home/asuka/.Xresources
exec openbox-session
