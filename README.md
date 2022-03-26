#Github Action Note

### Github Action Components
* Events - `An event that trigers worksflow`
* Jobs - `A set of step in a workflow that execute on the same runner`
* Runner - `Runner is a server that runs ower workflows when they're trigger - Each Runner can run a single job at a time(window,linux,macos...)`

### Github workflow
Each workflow is stored as a separate TAML file in our code repository called `.github/workflows`

### Workflow file Structure(yaml file)
```yaml
name: learn-github-actions
on: [push]
jobs:
  check-bats-version:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-node@v2
        with:
          node-version: '14'
      - run: npm install -g bats
      - run: bats -v
```
What's this workflow does?
* This workflow name is `learn-github-actions`
* The workflow will trigger on push event
* There is a job called `check-bats-version`
  * `check-bats-version` will run on `ubuntu` runner
  * `check-bats-version` are total 4 steps
    * run `v2` of `actions/checkout` action
    * run `v2` of `actions/setp-node` action to install Nodejs with node version 14
    * run a command `npm install -g bats` to install bats
    * run a command `bats -v` to check bats version