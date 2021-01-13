# Git-semver

Git-semver is a CLI tool for easy SemVer versioning of your tags.

`git-semver help` returns usage information.

`git-semver get` returns the latest semver tag.

`git-semver patch` increments the patch number, `inc semver minor` and `inc semver major` increment the minor and major number.

`git-semver patch --clear` increments the patch number even if current version has a pre-release/metadata part.

`git-semver patch --force` forces a new tag even if unnecessary or dirty repository.

`git-semver patch --yes` removes interactive y/N confirmation and assumes yes.

`git-semver prerelease rc1 -m "sounds about right"` creates a prerelease tag (`-rc1`) with a message (annotated tag)

## TODO

* [X] Allow setting an annotated version tag with a message
* [X] Generalize the patch command to minor and major ie. bumpWithBumper
* [ ] Rename to git-semver to integrate as a git plugin
* [X] Force flag instead of quiet
* [X] Add --clear option to inc patch even with release information (clears release info + meta before incPatch)
