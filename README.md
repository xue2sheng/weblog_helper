# WebLog Helper 

Little commandline utility to parse a log file.

## Description

In the file [public_access.log.txt](public_access_log.txt), you'll find a standard HTTP access log ... something that we all need to parse from time to time, usually in the heat of an outage.

Create a command line script called weblog_helper that can provide the two features listed below. Each feature should build on the previous implementation.
 
Feature 1: Return all log lines that correspond to a given source IP address

Add a switch (--ip <IP>) that restricts the output to the given IP address

Example:

    $ ./weblog_helper --ip 178.93.28.59
    178.93.28.59 - - [02/Jun/2015:17:06:06 -0700] "GET /logs/access_150122.log HTTP/1.1" 200 3240056 "http://fruit.fm/20487/blog/1327873/" "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.124 Safari/537.36" "redlug.com"
    178.93.28.59 - - [02/Jun/2015:17:06:09 -0700] "GET /logs/access_150122.log HTTP/1.1" 200 3240056 "http://fruit.fm/20487/blog/1327873/" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.124 YaBrowser/14.10.2062.12061 Safari/537.36" "redlug.com"
    ...

Feature 2: Return all log lines that correspond to a given IP CIDR network ( e.g. 180.76.15.0/24)

Expand the --ip switch to handle CIDR ranges. Network address libraries are permitted.

Example:

    $ ./weblog_helper --ip 180.76.15.0/24
    180.76.15.135 - - [02/Jun/2015:17:05:23 -0700] "GET /logs/access_140730.log HTTP/1.1" 200 979626 "-" "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)" "www.redlug.com"
    180.76.15.137 - - [02/Jun/2015:17:05:28 -0700] "GET /logs/access_140730.log HTTP/1.1" 200 7849856 "-" "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)" "www.redlug.com"
    180.76.15.17 - - [02/Jun/2015:17:20:23 -0700] "GET /logs/access_141026.log HTTP/1.1" 200 45768 "-" "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)" "www.redlug.com"

