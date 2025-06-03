# CHANGELOG.md Update Process

- Please add a new version and summarize update information based on the CHANGELOG.md information
- Current version: This is the version at the top of CHANGELOG.md.
- Added changes: Compare the differences between the main branch and the current version tag.
- How to check changed files: Create a diff file and check the contents.

```bash
git diff v`{current-version}`..main > diff.txt
```
