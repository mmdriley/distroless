## Extra root filesystem contents

### tmp

Programs expect this to exist.

Must be `chmod 1777` so `Dockerfile`s with different `USER` values can use it ([ref](https://askubuntu.com/a/432703))
