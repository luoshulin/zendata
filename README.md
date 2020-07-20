# zendata
zendata是一款通用的数据生成工具，您可以使用yaml文件来定义您的数据格式，然后交由zendata生成。

## 参数：
```shell
-d  --default    默认的数据格式配置文件。
-c  --config     当前场景的数据格式配置文件，可以覆盖默认文件里面的设置。
-o  --output     生成的数据的文件名。可通过扩展名指定输出json|xml|sql格式的数据。默认输出原始格式的文本数据。
-n  --lines      要生成的记录条数，默认为10条。
-r  --recursive  递归模式。如不指定，默认为平行模式。平行模式下各个字段独立循环。
                 递归模式下每个字段的取值依赖于前一字段。可增强数据的随机性。

-F  --field      可通过该参数指定要输出的字段列表，用逗号分隔。 默认是所有的字段。
-t  --table      输出格式为sql时，需通过该参数指定要插入数据的表名。
-H  --human      输出可读格式，打印字段名，并使用tab键进行分割。

-b  --bind       监听的ip地址，默认监听所有的ip地址。
-p  --port       在指定端口上运行HTTP服务。可通过http://ip/接口获得JSON格式的数据。服务模式下只支持数据生成。
-R  --root       运行HTTP服务时根目录。客户端可调用该根目录下面的配置文件。如果不指定，取zd可执行文件所在目录。

-i  --input      指定一个schema文件，输出每个表的yaml配置文件。需通过-o参数指定一个输出的目录。
-s  --server     数据库服务器类型，支持mysql|oracle|sqlite|sqlserver，默认为mysql。可用于解析yaml文件或者生成SQL。
-D  --decode     根据指定的配置文件，将通过-i参数指定的数据文件解析成json格式，可通过-H参数输出可读格式。

-e  --example    打印示例的数据格式配置文件。
-l  --list       列出所有支持的数据格式。
-v  --view       查看某一个数据格式的详细定义。
-h  --help       打印帮助。
```

## 命令行模式举例：
```shell
$>zd.exe -d demo\default.yaml 根据-d参数指定的配置文件生成10条记录。
$>zd.exe -c demo\default.yaml 根据-c参数指定的配置文件生成10条记录。
$>zd.exe -d demo\default.yaml -c demo\test.yaml -n 100  -c和-d两个文件的配置合并，输出100条记录。

$>zd.exe -d demo\default.yaml -c demo\test.yaml -n 100 -o test.txt   输出原始格式的数据。
$>zd.exe -d demo\default.yaml -c demo\test.yaml -n 100 -o test.json  输出json格式的数据。
$>zd.exe -d demo\default.yaml -c demo\test.yaml -n 100 -o test.xml   输出xml格式的数据。
$>zd.exe -d demo\default.yaml -n 100 -o test.sql -t user -s mysql    输出插入到user表里面的sql。

$>zd.exe -i db.sql -s mysql -o db  根据db.sql的定义生成每个表的yaml文件，存储到db目录里面。
$>zd.exe -c demo\default.yaml -i test.txt --decode  将-i指定的文件根据-d参数的配置进行解析。
```
## 服务模式举例：
```shell
$zd.exe -p 80 -R d:\zd\config  监听80端口，以d:\zd\config为根目录。
```

## 客户端调用：
```shell
$curl http://loclahost/?d=default.yaml&c=config.yaml&n=100&o=test.sql&t=user  通过GET方式指定服务器端配置文件。
$curl http://loclahost/?default=default.yamloutput=test.sql&table=user        参数名可以用全拼。
$curl -d "default=...&config=...&lines=10" http://localhost/                  可以通过POST方式上传配置。
```
