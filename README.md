# app directories

Small helper library to help you locate useful directories based on the following resources:

- https://stackoverflow.com/a/41598938
- https://stackoverflow.com/a/39868933
- https://apple.stackexchange.com/questions/28928/what-is-the-macos-equivalent-to-windows-appdata-folder

## Add it to your project

```bash
go get github.com/zarthus/appdir
```

## Description

API:

```js
GetApplicationDirectory(); // not for linux (presently)
GetUserConfigDirectory();
```
