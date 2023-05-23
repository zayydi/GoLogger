# GoLogger
GoLogger is a simple logging solution for your golang applications and other apps in general as well. GoLogger saves all your logs in ElasticSearch and you can then visualize these logs in Grafana.

## Using GoLogger as a service
GoLogger can easily be used as a service by running this command:
```
go install https://github.com/zayydi/GoLogger
```
This command will install GoLogger.exe globally in your system then, You can simply open cmd and type this command:
```
gologger --watcherPath=LOGS_FOLDER_TO_WATCH
```
gologger will start watching to the folderpath given for new log files and whenever a new log file gets added or an old one gets modified it will read the contents of the file line by line and push that to elastic search.
Typically, gologger would want the data to be formatted in json like this:
```
{"appName": "AppX", "level": "error", "description": "Null pointer exception occurred.", "trace": "at com.example.MyClass.method1(MyClass.java:25)"}
```
If the data is not formatted like this, gologger will consider the `level: unknown` and put the whole line in `description`, Trace is optional so if not provided it will be considered `nil`.

## Using GoLogger as a Library
If you're looking to use gologger as a library you can do that as well. For that, You have to import the whole project using
```
go get https://github.com/zayydi/GoLogger
```
Afterwards, You can use the submodule `logmon` to log your errors, warnings and messages using gologger.
