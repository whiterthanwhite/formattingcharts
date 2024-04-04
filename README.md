# formattingcharts

Application returns json based on interval date

## Flags

`-i` - interval type for graph formatting
- 0 - Hour
- 1 - Day (default)
- 2 - Week
- 3 - Month

## How to use

Application get data on os.Stdin and write result to os.Stdout.

`cat input.json | go run . > result.json`

```
input.json
[
	{
		"value": 4456,
		"timestamp": 1616026248
	},
	{
		"value": 4456,
		"timestamp": 1616026248
	},
	{
		"value": 4231,
		"timestamp": 1616022648
	},
	{
		"value": 5212,
		"timestamp": 1616019048
	},
	{
		"value": 4321,
		"timestamp": 1615889448
	},
	{
		"value": 4567,
		"timestamp": 1615871448
	}
]

result.json
[
    {
        "Value": 4456,
        "Timestamp": 1616025600
    },
    {
        "Value": 5212,
        "Timestamp": 1615939200
    },
    {
        "Value": 4567,
        "Timestamp": 1615852800
    }
]
```