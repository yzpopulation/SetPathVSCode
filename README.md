# SetPathVSCode
Add right menu for https://portapps.io/app/vscode-portable/

go install github.com/akavel/rsrc 

rsrc -manifest app.manifest -ico icon.ico -o app.syso

icon from https://icons8.com/icons/set/vscode

<!-- set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64 -->
go build -ldflags="-s -w" -trimpath