# Dynamic Wallpapers

![example workflow](https://github.com/dawsonalex/dynamic-wallpaper/workflows/Build/badge.svg)

Time based dynamic wallpapers.

## GitHub Workflow

The GitHub workflow defined in `base.yml` attempts to do some common things in a simple way. Currently, it does the
following steps under a single job called `Build`:

    - Set up Go environment.
    - Run Go Tests (ignoring if there are none).
    - Runs `make release` to create all binaries.
    - Bump the version based on the commit message.
        - Use `#major`, `#minor`, or `#patch` in your commit message to bump the version and create a new release.
        - Leaving out the above tags will not create a new tag or release version.
    - Generate release logs from the commits between this tag and the last.
    - Create a GitHub release and upload the content of `bin`.