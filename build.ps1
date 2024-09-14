$ExistingInstance = (Get-Command GoCopy-Item.exe).Source
go build -o $env:TEMP\GoCopy-Item.exe
Move-Item $env:TEMP\GoCopy-Item.exe $ExistingInstance -Force -Verbose