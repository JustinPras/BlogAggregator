# BlogAggregator

### How to set up the blog aggregator
1. Ensure Postgres and Go is installed
2. Run: ``` go install github.com/JustinPras/BlogAggregator ```

4. The blog aggregator requires a config file at `home/.gatorconfig.json`. The following should be the contents of the file:
```
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}
```

### Commands
Example usage of commands:
``` go run . command_name ```

#### List of commands:
- ``` register ```
- ``` login ```
- ``` reset ```
- ``` users ```
- ``` agg ```
- ``` addfeed ```
- ``` feeds ```
- ``` follow ```
- ``` following ```
- ``` unfollow ```
- ``` browse ```
