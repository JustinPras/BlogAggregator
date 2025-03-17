# BlogAggregator (gator)

## How to set up the blog aggregator
1. Ensure Postgres and Go is installed
2. Run: ``` go install github.com/JustinPras/BlogAggregator ```

4. The blog aggregator requires a config file at `home/.gatorconfig.json`. The following should be the contents of the file:
```
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}
```

## Commands
Example usage of commands:
``` BlogAggregator <command_name> <optional_args> ```

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

## To-Do List
- [ ] Add sorting and filtering options to the browse command
- [ ] Add pagination to the browse command
- [ ] Add concurrency to the agg command so that it can fetch more frequently
- [ ] Add a search command that allows for fuzzy searching of posts
- [ ] Add bookmarking or liking posts
- [ ] Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)
- [ ] Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
- [ ] Write a service manager that keeps the agg command running in the background and restarts it if it crashes

