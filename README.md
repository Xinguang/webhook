webhook
=====
[![Build Status](https://travis-ci.org/starboychina/webhook.svg)](https://travis-ci.org/starboychina/webhook)

---

## What is `webhook`?

`webhook` is a little webserver written in go. He waits for webhook calls by github (or bitbucket, gogs) to run little shell commands.

## How to use

Just edit the config.json to your needs. A short example:
You want to track the status of your Repository URL and the branch master. If there is an update to this branch you want to execute your shell script.

```json
{
    "Hooks":[
        {
          "Repo":"Repo URL",
          "Branch":"development",
          "Shell":"cd /var/ && git pull && make && make install"
        },
        {
          "Repo":"Repo URL",
          "Branch":"staging",
          "Shell":"docker restart app-node"
        },
        {
          "Repo":"Repo URL",
          "Branch":"production",
          "Shell":"script.sh"
        }
    ]
}
```

Now start the server with

```sh
    go run *.go -p 9000
```

and add a git-webhook for your.domain.com:9000/github. Everytime you push to master, your script gets executed.
or add your.domain.com:9000/bitbucket. for bitbucket Repository, add your.domain.com:9000/gogs. for gogs Repository.

## Using Docker

    docker run -d --name webhook -p 9000:9000 -v $(pwd)/config.json:/config.json starboychina/webhook

With Docker support (access host docker inside container)

    docker run -d --name webhook -p 9000:9000 -v $(pwd)/config.json:/config.json -v /var/run/docker.sock:/var/run/docker.sock starboychina/webhook


## License

This Software is licensed under the MIT License.

Copyright (c) 2016 starboychina

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
