# go-example-ollama2-client

This wants be one gist for describe a client Ollama.

### Step 1: For example you need orca2 because it is sentiment analyzer
```
ollama list
NAME                            ID              SIZE    MODIFIED   
codellama:13b                   9f438cb9cd58    7.4 GB  6 days ago
codeup:latest                   54289661f7a9    7.4 GB  5 days ago
llama2:13b                      b3f03629d9a6    7.4 GB  6 days ago
llama2:chat                     fe938a131f40    3.8 GB  6 days ago
llama2:latest                   fe938a131f40    3.8 GB  6 days ago
llama2:text                     f398b0c61128    3.8 GB  6 days ago
llama2-uncensored:latest        44040b922233    3.8 GB  6 days ago
magicoder:latest                8007de06f5d9    3.8 GB  6 days ago
mistral:latest                  d364aa8d131e    4.1 GB  6 days ago
mistral-openorca:latest         12dc6acc14d0    4.1 GB  5 days ago
neural-chat:latest              f4c6a8e532e8    4.1 GB  6 days ago
orca-mini:latest                2dbd9f439647    2.0 GB  5 days ago
orca2:latest                    ea98cc422de3    3.8 GB  5 days ago
sqlcoder:latest                 77ac14348387    4.1 GB  5 days ago
stablelm-zephyr:latest          7c596e78b1fc    1.6 GB  6 days ago
starling-lm:latest              ff4752739ae4    4.1 GB  6 days ago
vicuna:latest                   370739dc897b    3.8 GB  6 days ago
zephyr:latest                   03af36d860cc    4.1 GB  6 days ago
```

### Step 2: Run server

```
ollama serve
2023/12/13 23:21:05 images.go:734: total blobs: 75
2023/12/13 23:21:05 images.go:741: total unused blobs removed: 2
2023/12/13 23:21:05 routes.go:787: Listening on 127.0.0.1:11434 (version 0.1.13)
```

### Step 3: Try example
```
go run main.go  -prompt "good product"
POSITIVE

go run main.go  -prompt "bad product"
NEGATIVE

go run main.go  -prompt "this is the  product"
NEUTRAL

```

### Step 4: Custom by line console

```
go run main.go  -prompt "my name is Arturo" -model vicuna -system "you are mario bros" -template ""

Hi Arturo! It's nice to meet you. Is there anything you would like to chat about or ask me a question? I'm here to help with any information or advice you may need.
```


### Step 5: Custom by JSON config (properties.json)


```json
{
  "url": "http://localhost:11434/api/generate",
  "system": "./prompt.system.txt",
  "prompt": "./prompt.user.txt",
  "template": "./prompt.template.txt",
  "model": "orca2",
  "temperature": 1.0
}
```

