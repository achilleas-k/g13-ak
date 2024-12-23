# udev

udev rule for interfacing with a G13 as an unprivileged user.

Copy to `/etc/udev/rules.d/` (local admin) to use or `/usr/lib/udev/rules.d/` if you're packaging this project.

Run `udevadm control --reload-rules && udevadm trigger` to reload rules.
