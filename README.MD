# go-copyitem
This is just a toy project, been learning Golang for about 9 months or so and like to recreate/improve existing basic commandline utilities from bash or Powershell to see how far I can get. It's fun to me and its nice to get a better understanding of something that you normally take for granted in your OS. 

## Goals: Copy-Item shortcomings
How this in particular started was that I find it frustrating that the Copy-Item Commandlet doesn't have an easy way to track display progress when Copying a file. I know robocopy is still pretty much the standard tool most admins reach for anything large or important that needs to be copied (for good reason), but I use Copy-Item a lot in scripts when I'm just copying a single file or directory or at home to copy a movie or something to my NAS, mostly because I'm too lazy to lookup/remember robocopy params. 

I always find myself wishing/googling an easy way to have Copy-Item show the progress percentage of the Copy job and coming to the conclusion it's not worth the effort. So my main goal here is to make a simple cli utility that I can use like Copy-Item that will show a progress bar by default. Maybe more things like Completion ETA or other enhancements.

## OS
I use 100% Linux on all my homelab servers, but still run Windows 11 on my laptop that I use for dev. I'd like to migrate to Linux for desktop/laptop at somepoint, but haven't done so yet. So I am initially developing/running this on Windows. I haven't done enough golang dev for Windows to know if that will create any compatibility issues if trying on Linux. I would imaging there will be OS/Kernel specific things I'll have to call, we'll see.   
