# CHANGELOG.md Update Process

- Please add a new version and summarize update information based on the CHANGELOG.md information
- Current version: This is the version at the top of CHANGELOG.md.
- Added changes: Compare the differences between the main branch and the current version tag.
- How to check changed files: Create a diff file and check the contents.

```bash
git checkout main
git pull
git diff v`{current-version}`..main > diff.txt
```

## Versioning Rules

While we follow semantic versioning principles, we maintain the major version at 0.x.y during the development phase. The versioning rules are adjusted as follows:

1. **Major Version (0)**: Kept at 0 during the development phase.
   - Breaking changes are introduced (which would normally increment the major version)
2. **Minor Version (x)**: Incremented when:
   - New features are added
   - API changes that affect backward compatibility
3. **Patch Version (y)**: Incremented when:
   - Bug fixes are made
   - Non-breaking enhancements are added
   - Documentation updates or internal changes with no user-facing impact

### Examples:

- 0.15.2 → 0.15.3: Non-breaking changes, bug fixes, or minor enhancements
- 0.15.2 → 0.16.0: New features or breaking changes that would normally increment the major version
