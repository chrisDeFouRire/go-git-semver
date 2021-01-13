# Git-semver

Git-semver is a CLI tool for easy SemVer versioning of your tags.

`go-git-semver help` returns usage information.

`go-git-semver get` returns the latest semver tag.

`go-git-semver patch` increments the patch number, `inc semver minor` and `inc semver major` increment the minor and major number.

`go-git-semver patch --clear` increments the patch number even if current version has a pre-release/metadata part.

`go-git-semver patch --force` forces a new tag even if unnecessary or dirty repository.

`go-git-semver patch --yes` removes interactive y/N confirmation and assumes yes.

`go-git-semver prerelease rc1 -m "sounds about right"` creates a prerelease tag (`-rc1`) with a message (annotated tag)

## Install as a git plugin

Simply copy go-git-semver to git-semver somewhere on your PATH.
```
sudo cp go-git-semver /usr/local/bin/git-semver
```

Then `git semver get` will work.

## License

Apache 2.0, Copyright Chris Hartwig

If you think I should have used another license, open an issue on github

## TODO

* [X] Allow setting an annotated version tag with a message
* [X] Generalize the patch command to minor and major ie. bumpWithBumper
* [X] Rename to git-semver to integrate as a git plugin
* [X] Force flag instead of quiet
* [X] Add --clear option to inc patch even with release information (clears release info + meta before incPatch)
