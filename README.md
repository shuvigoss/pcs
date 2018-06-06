`./pcs -config=./pcs.json -command=\"whoami\""`

针对pcs.json 下配置的所有主机执行`whoami`命令

``` json
{
  "hosts": [
    {
      "host": "192.168.111.11",
      "username": "root",
      "password": "123456",
      "port": 22
    },
    {
      "host": "192.168.111.12"
    },
    {
      "host": "192.168.111.13"
    }
  ],
  "globalPwd": "123",
  "globalName": "root",
  "globalPort": 22
}
```