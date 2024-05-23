# Dijkstra's Algorithm Solver using Golang

# Usage:

This is a CLI tool for illustrating Dijkstra's shortest path algorithm in routing between computer network layer routers

The links information is provided inside link.json. This tell us from which node to which node what is the cost.
```bash
~ $ ./dijkstra --help   
Implementation of Dijkstra's Algorithm in go

Usage:
  Dijkstra's Algorithm [flags]

Flags:
  -e, --end string     Enter the End point of the route (default "B")
  -h, --help           help for Dijkstra's
  -s, --start string   Enter the starting point of the route (default "A")
  -t, --toggle         Help message for toggle
```  

## Steps:

1. Clone the Repo using `git clone https://github.com/Vinayaks439/Dijkstra-Algorithm-Go`
2. build the binary using `go build -o dijkstra`
3. Use -s to indicate starting node and -e to indicate ending node `./dijkstra -s "A" -e "J"` for shorthands
4. Use --start and --end to provide long hand cli args
5. For help use: `./dijkstra --help`


# Example 1

```bash
~ $ ./dijkstra --start "A" --end "J"
{
    "TotalCost": 6,
    "Path": [
        "A",
        "B",
        "C",
        "J"
    ]
}
# The total Cost to reach J node from A is 6 and the route is mentioned inside the PATH array one after the other (A -> B -> C -> J) 
```

# Example 2

```bash
~ $ ./dijkstra --start "A" --end "F"
{
    "TotalCost": 5,
    "Path": [
        "A",
        "B",
        "C",
        "F"
    ]
}
# The total Cost to reach F node from A is 5 and the route is mentioned inside the PATH array one after the other (A -> B -> C -> F)
```

# Example 3

```bash
~ $ ./dijkstra --start "K" --end "B"
{
    "TotalCost": 5,
    "Path": [
        "K",
        "F",
        "C",
        "B"
    ]
}
# The total Cost to reach B node from K is 5 and the route is mentioned inside the PATH array one after the other (K -> F -> C -> B)
```
