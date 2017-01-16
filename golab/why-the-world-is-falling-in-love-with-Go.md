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

# Who's using Go?

Adjust, Basecamp, BBC, Bit.ly, Bitbucket, Buzzfeed, Booking.com, Brightcove, CloudFlare, CoreOS, DailyMotion, DigitalOcean, Dell, Docker, DropBox, Google, GitHub, GrayMeta, eBay, Embedly, Facebook, GOV.UK, Intel, Medium, Monzo, Netflix, NotOnTheHighStreet, Outdoorsy, Pivotal, SoundCloud, Segment, SendGrid, SongKick, SpaceX, Trenìt!, Twitch, Uber, Yahoo, __and many more... I got bored typing them[^2]__

[^2]: For a more complete list, see [https://github.com/golang/go/wiki/GoUsers](https://github.com/golang/go/wiki/GoUsers)

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

^ Don't use these to try and convince your boss to let you do something in Go

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

![inline](https://www.dropbox.com/s/cy16anzo5ca8t49/parhamdoustdar.jpeg?dl=1)

* Meet Parham Doustdar a Blind Go programmer from Tehran, Iran
* Uses a screen reader to read code
* "Go's simplicity means it's very easy to read or hear; that's especially important to me"

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
Don't mind throwing things away because you haven't invested all this
effort up-front.

---

# Performance

![inline](https://www.dropbox.com/s/vf0k2y9lef9bbh6/salmon.jpg?dl=1)

---

# Performance



---

# Modern language features
# (and an excellent standard library)

![inline](https://www.dropbox.com/s/9as6scdnjh5b0he/apples.jpg?dl=1)

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
Design: Even on a single core processor (where you would get no benefit) the
code design might still improve productivity.

---

# `net/http` package

> "The http packages comprehensiveness has made building and running a load of http servers super easy, which given we'd picked a service oriented architecture definitely made that faster."
-- Matt Heath, Monzo

![inline](https://www.dropbox.com/s/kb842aurrwxnq92/heath.jpg?dl=1)

^ World class HTTP capabilities
Everything you need to build web services

---

# Builds to a self-contained static binary

![inline](https://www.dropbox.com/s/27nz6md0yayky8a/cardboard-box.png?dl=1)
