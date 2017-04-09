# labour

> simple configable spider core application with golang

### 0. Build
```shell
> go get github.com/ErosZy/singoriensis
> git clone https://github.com/ErosZy/labour
> cd labour
> go build main.go
```

### 1. Run(Windows x64)
```shell
> cd example
> labour.exe config/51job.json
```

### 2. Task Config
```javascript
{
    "targetUrl": "http://www.example.com",
    "threadNum": 4,
    "retryMaxCount": 10,
    "sleepTime": 500,
    "closeTime": 5,
    "requestTimeout": 2,
    "method": "GET",
    "headers": [{
        "key": "Host",
        "value": "www.example.com"
    }],
    "proxy": [],
    "schedulers": [{
        "route": "http://sou.example.com/jobs/searchresult.ashx",
        "text": {
            "regex": [],
            "xpath": [{
                "domStr": ".pagesDown ul li:not(.nextpagego-box):not(.clearfix) a",
                "key": "pageHref",
                "type": 1,
                "attrKey": "href"
            }],
            "json": []
        }
    }],
    "pages": [{
        "route": "http://sou.example.com/jobs/searchresult.ashx",
        "mainKey": "id",
        "text": {
            "regex": [],
            "xpath": [{
                "prefix": "#newlist_list_content_table",
                "arr": [{
                    "domStr": ".zwmc a",
                    "key": "name",
                    "type": 0,
                    "attrKey": ""
                }, {
                    "domStr": ".zwmc > input",
                    "key": "id",
                    "type": 1,
                    "attrKey": "value"
                }]
            }],
            "json": []
        }
    }]
}
```

