# go-copyitem
This is just a toy project, been learning Golang for about 9 months or so and like to recreate/improve existing basic commandline utilities from bash or Powershell to see how far I can get. It's fun to me and its nice to get a better understanding of something that you normally take for granted in your OS. 

When learning a new language I like starting new projects and see how far I can get before I have to google something beyond simple syntax or Copy-Paste from other projects. I like trying to get all the way to the first build/run to gauge my progress in a language.

I'm not sure how long I'll go before I ghost this repo, but I do like to get at least a very janky """MVP""" of these toy projects.

I ended up finishing v1.0.0 completed over a weekend. I don't want to spend too much time, but there are some more optimizations that I'd like to make.

## Original Goals: Copy-Item shortcomings
How this repo in particular started was from finding it frustrating that the Copy-Item Cmdlet doesn't have an easy way to display progress when Copying a file. I know robocopy is still pretty much the standard tool most admins reach for anything large or important that needs to be copied (for good reason), but I use Copy-Item a lot in scripts when I'm just copying a single file or directory or at home to copy a movie or something to my NAS, mostly because I'm too lazy to lookup/remember robocopy params. 

I always find myself wishing/googling an easy way to have Copy-Item show the progress percentage of the Copy job and coming to the conclusion it's not worth the effort. So my main goal here is to make a simple cli utility that I can use like Copy-Item that will show a progress bar by default. Maybe more things like Completion ETA or other enhancements.

Other goals would be performance related, I think this could be a good place to improve my skills with channels and concurrency in Go. 

Way down on the list is interop with other PowerShell Cmdlets. I'm not sure how feasable it would be to be able to pipe back and forth between a go binary and powershell comdlets or return objects with properties from go to PS, but it might be fun to look into.

## Possible Future Enhancements
- Average speed calculation
- ETA for completion
- Better logging and error handling
- General performance and concurrency
    - Currently the CopyFile() function is blocking, may be nice to copy the file in chunks with multiple go routines.
    - Test against other major copy cli utilities like cp, robocopy, rsync to see if it's equivalent
- Copy using SSH like rsync
- Support for Directories
- Autocomplete for parameters
- Create Installer

## Build and Install 
Clone down the repo
```
PS C:\Users\babbage88\AppData\Local\Temp> git clone https://github.com/babbage88/go-copyitem.git
Cloning into 'go-copyitem'...
remote: Enumerating objects: 78, done.
remote: Counting objects: 100% (78/78), done.
remote: Compressing objects: 100% (55/55), done.
remote: Total 78 (delta 42), reused 57 (delta 23), pack-reused 0 (from 0)
Receiving objects: 100% (78/78), 17.24 KiB | 430.00 KiB/s, done.
Resolving deltas: 100% (42/42), done.
```

Build the binary, using -o to specify the name of the executable
```
PS C:\Users\babbage88\AppData\Local\Temp> go build -o GoCopy-Item.exe
```
Copy the binary to C:\Program Files\Go\bin\ or any other direcotry in your path
```
PS C:\Users\babbage88\AppData\Local\Temp> cp .\GoCopy-Item.exe "C:\Program Files\Go\bin\GoCopy-Item.exe"
```

## Usage
Call the GoCopy-Item.exe and specifiy the source file with --source and destination with --destination

```
PS C:\Users\babbage88\AppData\Local\Temp> & GoCopy-Item.exe --source Z:\test.mkv --destination C:\Users\babbage88\AppData\Local\Temp\test.mkv
Source File Z:\test.mkv size is 1559.62 MB
Destination file C:\Users\babbage88\AppData\Local\Temp\test.mkv size is 0.00 MB

Starting Copy Job
Source File: Z:\test.mkv
Destination File: C:\Users\babbage88\AppData\Local\Temp\test.mkv
Starting File Copy Job src: Z:\test.mkv
dst: C:\Users\babbage88\AppData\Local\Temp\test.mkv
size_kb: 1597055.84
```

## OS
I use 100% Linux on all my homelab servers, but still run Windows 11 on my laptop that I use for dev. I'd like to migrate to Linux for desktop/laptop at somepoint, but haven't done so yet. 

So, although I'll verify it works on linux, it's currently being designed/used day to day on Windows.