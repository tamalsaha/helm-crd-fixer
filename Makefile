# This version-strategy uses git tags to set the version string
git_branch       := $(shell git rev-parse --abbrev-ref HEAD)
git_tag          := $(shell git describe --exact-match --abbrev=0 2>/dev/null || echo "")
commit_hash      := $(shell git rev-parse --verify HEAD)
commit_timestamp := $(shell date --date="@$$(git show -s --format=%ct)" --utc +%FT%T)

publish:
ifeq (,$(git_tag))
	ko publish . -B --platform=linux/amd64,linux/arm64
else
	ko publish . -B --platform=linux/amd64,linux/arm64 -t latest -t $(git_tag)
endif
