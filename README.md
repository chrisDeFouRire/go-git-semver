# Git-semver

Git-semver is a CLI tool for easy SemVer versioning of your tags.

`go-git-semver help` returns usage information.

`go-git-semver -v` prints go-git-semver's own version.

`go-git-semver get` returns the latest semver tag.

`go-git-semver patch` increments the patch number, `inc semver minor` and `inc semver major` increment the minor and major number.

`go-git-semver patch --clear` increments the patch number even if current version has a pre-release/metadata part.

`go-git-semver patch --force` forces a new tag even if unnecessary or dirty repository.

`go-git-semver patch --yes` removes interactive y/N confirmation and assumes yes.

`go-git-semver patch --prerelease=rc1 -m "sounds about right"` creates a prerelease tag (`-rc1`) with a message (annotated tag)

`go-git-semver purge "<0.3-0"` removes all semver tags according to the semver spec. Interactive use only. This command uses `git` and pushes the tag deletion to "origin". Also, it's slow. Add `-0` to your spec to include prereleases. This command will show the tags about to be removed, and will ask for a confirmation.

## Download and Install as a git plugin

First grab a binary (Linux/MacOS) [release on github](https://github.com/chrisDeFouRire/go-git-semver/releases/tag/v1.0.0-rc1).

Extract it and simply copy `go-git-semver` to `git-semver` somewhere on your `$PATH` (if you know how to build and deploy on Windows, please submit an issue/pull request!).

```
sudo cp go-git-semver /usr/local/bin/git-semver
```

Then `git semver <command>` will work. For instance `git semver major` will tag with a major tag bump.

## License

Apache 2.0, Copyright Chris Hartwig

If you think I should have used another license, open an issue on github and lets talk about it.

## TODO

* [X] Allow setting an annotated version tag with a message
* [X] Generalize the patch command to minor and major ie. bumpWithBumper
* [X] Rename to git-semver to integrate as a git plugin
* [X] Force flag instead of quiet
* [X] Add --clear option to inc patch even with release information (clears release info + meta before incPatch)
