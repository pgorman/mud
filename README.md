# MUD

This is a dumb little from-scratch MUD (multi-user dungeon) just for me to noodle around with. Goals don't necessarily include producing a playable game.


## Connecting to the MUD

Clients connect like:

```
$ openssl s_client -quiet -connect example.com:2323
```

…where `:2323` is the MUD's port number.


## Running the MUD Server

When running the MUD server, a a self-signed SSL/TLS cert can be generated like:

```
openssl req -x510 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 365 -nodes -subj "/C=US/ST=Michigan/L=Detroit/O=MUD/OU=MUD/CN=localhost"
```


## License

Copyright (C) 2023 Paul Gorman

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

