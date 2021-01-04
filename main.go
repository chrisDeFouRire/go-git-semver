package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

const (
	// SemverMajor represents the major part of a semver
	SemverMajor = iota
	// SemverMinor represents the minor part of a semver
	SemverMinor
	// SemverPatch represents the patch part of a semver
	SemverPatch
	// SemverRest represents the rest part of a semver
	SemverRest
)

// Version represents a semver version
type Version struct {
	Major int
	Minor int
	Patch int
	Rest  string
}

// NewVersion returns the version corresponding to the tag, or an error
func NewVersion(tag string) (*Version, error) {
	found := re.FindSubmatch([]byte(tag))
	if found == nil || len(found) != 5 {
		return nil, ErrNotSemver
	}
	major, err := strconv.Atoi(string(found[1]))
	if err != nil {
		return nil, ErrCantParseNumber
	}
	minor, err := strconv.Atoi(string(found[2]))
	if err != nil {
		return nil, ErrCantParseNumber
	}
	patch, err := strconv.Atoi(string(found[3]))
	if err != nil {
		return nil, ErrCantParseNumber
	}

	return &Version{major, minor, patch, string(found[4])}, nil
}

// String returns a string representation
func (v Version) String() string {
	if v.Rest == "" {
		return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
	}
	return fmt.Sprintf("v%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.Rest)
}

// Inc increases one field and returns a new Version
func (v Version) Inc(field int) Version {
	switch field {
	case SemverMajor:
		return Version{v.Major + 1, 0, 0, ""}
	case SemverMinor:
		return Version{v.Major, v.Minor + 1, 0, ""}
	case SemverPatch:
		return Version{v.Major, v.Minor, v.Patch + 1, ""}
	default:
		return Version{}
	}
}

var (
	// ErrNotSemver if a tag can't be parsed as a semantic versioning tag
	ErrNotSemver = fmt.Errorf("Tag is not SEMVER")
	// ErrCantParseNumber if a tag number can't be parsed
	ErrCantParseNumber = fmt.Errorf("Tag number can't be parsed")

	re = regexp.MustCompile(`^v([0-9]+)\.([0-9]+)\.([0-9]+)-?(.*)$`)
)

func main() {

	tagList := make(map[plumbing.Hash]string)

	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal(err)
	}

	/* Get all tags indexed by hash */

	//log.Println(">> find TAGS ")
	tags, err := repo.Tags()
	if err != nil {
		log.Fatal(err)
	}
	defer tags.Close()

	for ref, err := tags.Next(); err == nil; ref, err = tags.Next() {
		tagName := ref.Name().Short()
		tagList[ref.Hash()] = tagName
	}

	//log.Println(">> Inspect log")
	iter, err := repo.Log(&git.LogOptions{})
	defer iter.Close()

	for ref, err := iter.Next(); err == nil; ref, err = iter.Next() {
		tag, found := tagList[ref.Hash]
		if found {
			v, err := NewVersion(tag)

			if err == nil {
				log.Print("Version ", v)
				break
			}
		}
	}

}
