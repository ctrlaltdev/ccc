# Conventional Commit Checker

Will parse the commit message and will fail if it doesn't follow the convention: [Conventional Commits v1.0.0](https://www.conventionalcommits.org/en/v1.0.0/)

# Installation

```sh
echo <path to ccc> $@ >> .git/hooks/commit-msg
chmod a+x .git/hooks/commit-msg
```
