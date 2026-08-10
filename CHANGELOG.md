# Changelog

## v0.3.0 - 2026-08-10

### Added

* Extending templates now merges objects with the same `name` property. (#49)
* The path to generated manifests can now be set from the command line using `--generated-dir` (#47)

### Fixed

* Syntax errors with the generated manifest now report which resource generated the issue. (#52)
* The perEnvironment template function now returns the default value if none is declared at all. (#48)

--------

## v0.2.0 - 2026-08-03

### Changed

* The perEnvironment template function now accepts a default value as the second parameter.

--------

## v0.1.0 - 2026-07-24

* Initial Release
