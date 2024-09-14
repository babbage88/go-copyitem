$ExistingInstance = (Get-Command GoCopy-Item.exe).Source
Remove-Item -Force $ExistingInstance -Verbose
go build -o $env:TEMP\GoCopy-Item.exe
Move-Item $env:TEMP\GoCopy-Item.exe $ExistingInstance -Force -Verbose