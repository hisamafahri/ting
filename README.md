# Ting CLI

A Go CLI for everyday developer needs.

---

## Installation

```bash
go get github.com/hisamafahri/ting
```

After that, you chould be able to access `ting` command in your terminal.

```bash
ting <command>
```

---

## Available commands

### 1. Generate

```bash
ting generate <command>
# or
ting g <command>

```

- Generate UUID:
Generate UUID string. Currently only support [v1](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_1_(date-time_and_MAC_address)) and [v4 (default)](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) version.

```bash
ting g uuid

# Sample output:
# 6a350fa6-0db9-40ed-80f8-5336f8ac0420

ting g uuid -v="v1" -t=5 # or: ting g uuid --uuidVersion="v1" --total=5

# Sample Output:
# 1adf271c-7cf3-11ec-b21d-c2f0362b0d18
# 1adf3a18-7cf3-11ec-b21d-c2f0362b0d18
# 1adf3a72-7cf3-11ec-b21d-c2f0362b0d18
# 1adf3a90-7cf3-11ec-b21d-c2f0362b0d18
# 1adf3aae-7cf3-11ec-b21d-c2f0362b0d18

```

- Generate JWT Token:
Generate [JSON Web Token (JWT)](https://en.wikipedia.org/wiki/JSON_Web_Token) string. The token are available in both valid (30 min) and invalid/expired (30 min ago) form.

```bash
ting g jwt

# Sample output:
# eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMwMTY1NTksImlhdCI6MTY0MzAxNDc1OSwiaXNzIjoidGluZy1jbGktdG9vbHMiLCJqdGkiOiI0MWI5YzI0Ni0zMWIyLTQ2ODItYjk1Zi1hZGQ1MjNmZmM2OGIifQ.sCHUVBjDuxJEdcm14cWzAhl-oDJxNvACI9ThEe6VqSU

ting g jwt --isValid=false -t=3 # or: ting g uuid --isValid=false --total=3

# Sample Output:
# eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMwMTMwMDcsImlhdCI6MTY0MzAxNDgwNywiaXNzIjoidGluZy1jbGktdG9vbHMiLCJqdGkiOiJiMjRhZGRlNS03MTdlLTRiY2UtYTI0Zi03Yzc4OGZhMjBiYjcifQ.LKgxTUi_e7Svzqv8IWDpI6xvR204mlt195ojZySNU6s
# eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMwMTMwMDcsImlhdCI6MTY0MzAxNDgwNywiaXNzIjoidGluZy1jbGktdG9vbHMiLCJqdGkiOiJlYWUyMjAzNi01NTQ5LTRiZDktYjJjZC0wNDY2N2MwZDM1ZmYifQ.W6X5o8OzEAAS56oF8DhuldEB41xKPyb6smYtslsg32k
# eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMwMTMwMDcsImlhdCI6MTY0MzAxNDgwNywiaXNzIjoidGluZy1jbGktdG9vbHMiLCJqdGkiOiI2NjRkMDc3Mi02NDFiLTQzNzUtOWYxYi1hMzI0NzYwMGI0ZWYifQ.uB4pUXVEIqOv_ff-jarvkrBwQoYpJ9LFVV5yCsQ1N7Y

```


### 2. Help

CLI helper are available in general command or for each specific command.

```bash
ting help
# or
ting <command> --help # example: ting g uuid --help

```

---

*Other commands are coming soon*

## Author

[Hisam A Fahri](https://hisamafahri.com): [@hisamafahri](https://github.com/hisamafahri)

## License

[GNU GPLv3](LICENSE)