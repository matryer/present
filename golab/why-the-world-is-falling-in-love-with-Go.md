# Why the world is falling in love with Go

Mat Ryer - @matryer
GoLab.io

![inline](https://www.dropbox.com/s/nwl8omdmhn3tt50/GoLab-logo.eps?dl=1)

---

# Mat Ryer

* __Gopher__ since before v1 release
* __Author__ Go Programming Blueprints: Second Edition
* __Blog__ about Go on [matryer.com](http://matryer.com)
* __Work__ GrayMeta - 100% Go backend
* Not Banksy

---

# This talk

* Go's increasing popularity
* Is Go ready for the big time?
* Who's using Go? How and why?
* How you can introduce Go at work
* Challenge

^ We in Go conference have our view, but people in the world who haven't been exposed to Go yet, and it's our job to take it to them.

---

# Go is growing

![inline](https://www.dropbox.com/s/kp569ay6hxi6t3a/golang-trends-4.png?dl=1)

<sub>from Google Trends [https://www.google.co.uk/trends/explore?q=golang](https://www.google.co.uk/trends/explore?q=golang)</sub>

2017 will see more groups, conferences, meet-ups, projects.

---

# GitHub Octoverse 2016

![inline](https://www.dropbox.com/s/poflpbzht2vhv8l/octoverse-go.png?dl=1)

* Almost double the repositories[^1]

[^1]: https://octoverse.github.com

---

# Is Go ready for the big time?

Go is already running in production in far more places than you might realise.

It's definitely past the __fad__ stage.

---

# Who's using Go?

Adjust, Basecamp, BBC, Bit.ly, Bitbucket, Buzzfeed, Booking.com, Brightcove, CloudFlare, CoreOS, DailyMotion, DigitalOcean, Dell, Docker, DropBox, Google, GitHub, GrayMeta, eBay, Embedly, Facebook, GOV.UK, Intel, Medium, Monzo, Netflix, NotOnTheHighStreet, Outdoorsy, Pivotal, SoundCloud, Segment, SendGrid, SongKick, SpaceX, Trenìt!, Twitch, Uber, Yahoo, __and many more... I got bored typing them[^2]__

[^2]: For a more complete list, see [https://github.com/golang/go/wiki/GoUsers](https://github.com/golang/go/wiki/GoUsers)

---

# ![inline](https://www.dropbox.com/s/usjg3jsyf2c9yb9/monzo-card.png?dl=1) Monzo

* A bank - written in Go
* Microservices architecture
* Imagine if your current account had an API
* Search for "Matt Heath golang" to learn more about their tech

^ No question about Go running in production.
I'm using my card in Italy - doesn't get more 'production ready' than this!

---

# Why?

^ Before we investigate why this is happening, let's look
at some terrible reasons...

---

# Terrible reasons to choose Go

* Trendy & different
* Doesn't have semicolons
* Google made it
* I <3 Go
* Opensource

(I secretly think these are all great reasons)

^ Don't use these to try and convince your boss to let you do something in Go.
I've spent a few years talking to people in to Go community...

---

# The top reasons

* Simplicity & minimalism
* Productivity (especially `gofmt`)
* Performance
* Modern language features and excellent standard library
* Builds to a self-contained static binary

^ Simplicity: Obvious but Go codifies into the langauge
Productivity: Developers can quickly get stuff done
Performance: Even when you don't optimise code, it runs quickly
Built-in concurrency: Get the most out of a processor by using all the cores
Static binary: Again simple, binary with no other dependencies is dev ops dream

---

# Simplicity & minimalism

![inline](https://www.dropbox.com/s/j41gqp8tpshhl00/minimalist.jpg?dl=1)

---

# Simplicity & minimalism

* Small language spec (only 25 keywords)[^3]
* Familiar syntax
* Only a few core features
* Easy to pick up
* Addictive philosophy

^ Tiny language.
Looks like C.
Core features: It doesn't try and solve every problem, give you primitives that you can use
like structs, interfaces, functions, vars, goroutines, channels, maps, slices, not much else
Philosophy spreads: Seek simplicity everywhere: Process, tools
...but so what? Some obvious benefits, but some surprising ones...

[^3]: See the complete list (as well as some Easter eggs) at [route.run/gokeywords](https://route.run/gokeywords)

---

# Diversity

> "I think Go can be a great first language to learn due to its simplicity and minimalism. This greatly lowers the barrier for people generally excluded from server-side programming."
-- Carlisia Pinto, Founding member of GoBridge

![inline](https://www.dropbox.com/s/6oe14az5j5aizgu/carlisia.jpg?dl=1)

^ Community decided to maximize diversity (for selfish reasons)
Easy to get started helps there

---

# Accessibility

> "Go's simplicity means it's very easy to read or hear; that's especially important to me"
-- Parham Doustdar

![inline](https://www.dropbox.com/s/cy16anzo5ca8t49/parhamdoustdar.jpeg?dl=1)

* Blind Go programmer from Tehran, Iran, who uses a screen reader

^ We glance at code to understand it, if it has to be read out
you can see why not having semi-colons is pretty attractive.

---

# Productivity

![inline](https://www.dropbox.com/s/kwcvvtmrd6x7a2a/working.jpg?dl=1)

---

# Productivity

* Getting stuff done
* Solve problems quickly
* Quickly be useful to a new team
* Join open-source projects with low friction
* Low investment: don't mind trying things or even throwing bits away

^ `gofmt` all Go code looks the same
Don't mind throwing things away because you haven't invested all this effort up-front.

---

# Productivity

> "When building a startup, you have to start by solving a single problem as simply as possible. Go enabled us to build powerful, defect-free capabilities for customers quickly (and therefor cheaply) delivering value to them and confidence to our investors."
-- Aaron Edell, GrayMeta

---

# Performance

![inline](https://www.dropbox.com/s/vf0k2y9lef9bbh6/salmon.jpg?dl=1)

---

# Performance

> "We had a spike of about 2.4 million requests within ten minutes, and forgot to turn caching on! Luckily, since it only touched our Go code, it didn’t matter; we didn’t crash."
-- Dean Elbaz, SongKick

^ Without much effort, Go will perform great.
Doesn't mean you don't have to worry about it - but it's nice to have a headstart.

---

# Modern language features
# (and an excellent standard library)

![inline](https://www.dropbox.com/s/9as6scdnjh5b0he/apples.jpg?dl=1)

^ Go is a modern language, and was designed within the context of how we ended up writing and deploying software.
Changed a lot in last ten years.

---

# Concurrency

* Make the most of the CPU
* Easy to run code in the background:

```go
	doWork() 	// block
	go doWork() // run in background
```

* Channels for communicating and syncing

```go
	ch := make(chan string)
	ch <- "Hi there"
```

^ sync package too.
Design: Even on a single core processor (where you would get no benefit) the code design might still improve productivity.

---

# `net/http` package

> "The http package's comprehensiveness has made building and running a load of http servers super easy, which given we'd picked a service oriented architecture definitely made that faster."
-- Matt Heath, Monzo

![inline](https://www.dropbox.com/s/kb842aurrwxnq92/heath.jpg?dl=1)

^ Standard library is awesome. Built by world class engineers.
Extremely high standards.
World class HTTP capabilities: Everything you need to build web services

---

# Builds to a self-contained static binary

![inline](https://www.dropbox.com/s/27nz6md0yayky8a/cardboard-box.png?dl=1)

---

# Builds to a self-contained static binary

> "Deployments are extremely simple when you just have a self-contained binary without having to worry about special runtimes, package managers, external dependencies, proxies, or anything like that."
-- Ernesto Jiménez, [Slackline.io](https://Slackline.io)

![inline](https://www.dropbox.com/s/ebbpgpfxv54dklz/ernesto.png?dl=1)

^ `go build` produces an executable binary whose only dependency is the operating system
Other languages have specific dependencies, usually a specific version
It means the binaries can be chunky - but in modern situations, it's the right choice.

---

# How to introduce Go at work

![inline](https://www.dropbox.com/s/gcygq6obw9f4lft/skeletons.jpg?dl=1)

^ What if you want to introduce Go but are getting resistence?

* Is your team resistant?
* Too busy to "try something new"
* "It's too risky"

---

# Build something

* Solve a little annoying problem for your team
* Do a short presentation after stand-up
* Show the code, talk about how it works
* Invite others to play with it, and even add a feature

^ Build something in your spare time
Pick something small, that's not on the roadmap

---

# Challenge

__Introduce three new people to Go__

1. Existing programmer friend 
1. Someone you work with now (maybe your boss?)
1. Younger family member who has never coded before

Share your progress on Twitter with #golangshare

<sub>TIP: Take a picture of this screen</sub>

^ Send them a video.
Show them some Go code.
Link them to https://tour.golang.org/
You can do this anytime - if you're watching this video in the future

---

# Challenge ideas

* Send them your favourite Go video
* Show them some code you've written and explain it
* Sit down and go through [https://tour.golang.org/](https://tour.golang.org/) with them

---

# Thank you

* Twitter: @matryer
* Blog: matryer.com

<sub>Special thanks to Dave Cheney, Parham Doustdar, Dean Elbaz, Matt Heath, David Hernández, Ernesto Jiménez, Bill Kennedy, Brian Ketelsen, Carlisia Pinto, Alex Rodríguez</sub>
