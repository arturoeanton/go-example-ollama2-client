# go-example-ollama2-client


```
ollama serve
2023/12/13 23:21:05 images.go:734: total blobs: 75
2023/12/13 23:21:05 images.go:741: total unused blobs removed: 2
2023/12/13 23:21:05 routes.go:787: Listening on 127.0.0.1:11434 (version 0.1.13)
```

```
go run main.go  -prompt "good product"
POSITIVE

go run main.go  -prompt "bad product"
NEGATIVE

go run main.go  -prompt "this is the  product"
NEUTRAL

```

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

