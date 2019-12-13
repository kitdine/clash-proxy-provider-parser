# clash-proxy-provider-parser

针对`clash`的`proxy-provider`提供clash 配置文件转换为provider格式。

逻辑为通过http请求拉取远端的clash配置文件，然后抽取出`Proxy`节点内容，然后将其转为成`proxies`，其余内容全部删除。

## 使用

### 编译

```bash
go build .
```

目前可配置性极低，只能做转换。方式为：

```bash
curl http://127.0.0.1:55533/clash/parse?source=xxx
# xxx 为远端请求地址，并且经过UrlEncode
```

### CFW

如果你使用的是CFW，可以将该工具作为一个子进程启动：

```yaml
cfw-child-process:
  - command: clash-proxy-provider-parser.exe
    options:
      cwd: C:\xxx\ # clash-proxy-provider-parser.exe所在目录
```

并且cfw最新版本还支持ruleset，因此你的配置文件可以写成：

```yaml
mode: Rule
default: &default
  type: http
  interval: 3600
  health-check:
    enable: true
    url: http://www.gstatic.com/generate_204
    interval: 200
proxy-provider:
  A:
    path: D:/xxx/a.yaml
    url: http://127.0.0.1:55533/clash/parse?source=xxxa
    <<: *default
  B:
    path: D:/xxx/b.yaml
    url: http://127.0.0.1:55533/clash/parse?source=xxxb
    <<: *default
Proxy Group:
  - name: Proxy
    type: select
    use:
      - A
      - B
Rule:
  - RULE-SET,https://raw.githubusercontent.com/ConnersHua/Profiles/master/Surge/Ruleset/Unbreak.list,DIRECT
  - RULE-SET,https://raw.githubusercontent.com/ConnersHua/Profiles/master/Surge/Ruleset/Global.list,Proxy
  - MATCH,DIRECT
```

