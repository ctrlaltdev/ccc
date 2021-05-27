# Conventional Commit Checker

Will parse the commit message and will fail if it doesn't follow the convention: [Conventional Commits v1.0.0](https://www.conventionalcommits.org/en/v1.0.0/)

## Installation

```sh
brew install ctrlaltdev/tap/ccc
```
or 
```sh
brew tap ctrlaltdev/tap
brew install ccc
```

## Automatic Git Hook Installation

From within the repository you want to install the hook in:
```sh
ccc -init
```

## Manual Git Hook Installation

```sh
echo ccc -f $@ >> .git/hooks/commit-msg
chmod a+x .git/hooks/commit-msg
```
