# utils

### parser/range_parser

When we define an api as following format, we did not want add data limit filed into the api protocol

**request**
```json
{
	"user": "gogap",
	"start_time": "2015-10-10 22:11:33",
	"end_time": "2015-10-10 23:11:33"
	"max": "100",
	"offset": "100",
	"order_by": "end_time",
	"order": "desc"
	"want": "col1,col2,col3"
}
```

**response**
```json
[{
	col1:"",
	col2:"",
	col3:""
}]
```

So, you can see we need add `order` `order_by` `offset` `max` and `want` to every query like api protocol, when we use http protocol, we want add this options into request header, so it changed like as folloing


**header**

```
X-RANGE: order_by=end_time;order=desc;max=100;offset=100;want=col1 col2 col3
```

**body**

```json
{
	"user": "gogap",
	"start_time": "2015-10-10 22:11:33",
	"end_time": "2015-10-10 23:11:33"
}
```

It was clearly to define api protocol and use the option every where.
