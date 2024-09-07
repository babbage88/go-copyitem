# go-copyitem
This is just a toy project, been learning Golang for about 9 months or so and like to recreate/improve existing basic commandline utilities from bash or Powershell to see how far I can get. It's fun to me and its nice to get a better understanding of something that you normally take for granted in your OS. 

When learning a new language I like starting new projects and see how far I can get before I have to google something beyond simple syntax or Copy-Paste from other projects. I like trying to get all the way to the first build/run to gauge my progress in a language.

I'm not sure how long I'll go before I ghost this repo, but I do like to get at least a very janky """MVP""" of these toy projects. I'd love o be able to hammer out all the goals/requirements in a single Saturday, but I'm not quite there with Golang yet. 

## Goals: Copy-Item shortcomings
How this in particular started was that I find it frustrating that the Copy-Item Commandlet doesn't have an easy way to track display progress when Copying a file. I know robocopy is still pretty much the standard tool most admins reach for anything large or important that needs to be copied (for good reason), but I use Copy-Item a lot in scripts when I'm just copying a single file or directory or at home to copy a movie or something to my NAS, mostly because I'm too lazy to lookup/remember robocopy params. 

I always find myself wishing/googling an easy way to have Copy-Item show the progress percentage of the Copy job and coming to the conclusion it's not worth the effort. So my main goal here is to make a simple cli utility that I can use like Copy-Item that will show a progress bar by default. Maybe more things like Completion ETA or other enhancements.

Other goals would be performance related, I think this could be a good place to lear channels and general concurrency. 

Way down on the list is interop with other PowerShell Cmdlets. I'm not sure how feasable it would be to be able to pipe back and forth between a go binary and powershell comdlets or return objects with properties from go to PS, but it might be fun to look into. 

## OS
I use 100% Linux on all my homelab servers, but still run Windows 11 on my laptop that I use for dev. I'd like to migrate to Linux for desktop/laptop at somepoint, but haven't done so yet. So I am initially developing/running this on Windows. I haven't done enough golang dev for Windows to know if that will create any compatibility issues if trying on Linux. I would imaging there will be OS/Kernel specific things I'll have to call, we'll see.   
