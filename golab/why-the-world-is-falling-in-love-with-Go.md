# *__Why the world is falling in love with Go__*

Mat Ryer - @matryer
GoLab.io

![inline](https://www.dropbox.com/s/nwl8omdmhn3tt50/GoLab-logo.eps?dl=1)

---

# *__Mat Ryer__*

* __Gopher__ since *before v1 release*
* __Author__ *Go Programming Blueprints: Second Edition*
* __Blog__ about Go on *[matryer.com](http://matryer.com)*
* __Work__ GrayMeta - *100% Go backend*
* Not Banksy

^ NEXT: Great talks today

---

# This talk

* Go's *increasing* popularity
* Is Go ready for the *big time*?
* *Who* is using Go? How and why?
* How you can *introduce Go* at work
* *Challenge*

---

# Go is *growing*

![inline](https://www.dropbox.com/s/kp569ay6hxi6t3a/golang-trends-4.png?dl=1)

<sub>from Google Trends [https://www.google.co.uk/trends/explore?q=golang](https://www.google.co.uk/trends/explore?q=golang)</sub>

2017 will see more groups, conferences, meet-ups, projects.

---

# GitHub Octoverse 2016

![inline](https://www.dropbox.com/s/poflpbzht2vhv8l/octoverse-go.png?dl=1)

* Almost *double* the repositories[^1]

[^1]: https://octoverse.github.com

---

# Is Go ready for the *big time*?

Go *is* already running in *production* in far more places than you might realise...

---

# *__Who is using Go?__*

Adjust, Basecamp, BBC, Bit.ly, Bitbucket, Buzzfeed, Booking.com, Brightcove, CloudFlare, CoreOS, DailyMotion, DigitalOcean, Dell, Docker, DropBox, Google, GitHub, GrayMeta, eBay, Embedly, Facebook, GOV.UK, Intel, Medium, Monzo, Netflix, Outdoorsy, Pivotal, Slackline, SoundCloud, Segment, SendGrid, SongKick, SpaceX, Trenìt!, Twitch, Uber, Yahoo, *__and many more... I got bored typing them[^2]__*

[^2]: For a more complete list, see [https://github.com/golang/go/wiki/GoUsers](https://github.com/golang/go/wiki/GoUsers)

^ Lots of companies
Not insignificant use
e.g. Not just a dev tool 
Critical stuff is written in Go

---

# Go *success* stories

* *CloudFlare* - __Internet security and magic__ 🔥
* *Monzo* - __A bank[^5]__ ![inline](https://www.dropbox.com/s/usjg3jsyf2c9yb9/monzo-card.png?dl=1) 
* *Slackline.io* - __Shared channels between Slack teams__ 💬

[^5]: Search for "Matt Heath Golang" for talks about Go microservices at Monzo

^Different kinds of projects
Monzo: Using my card in Italy to pay for things
Slackline: Message brokering
SpaceX: Telemetry system
Google obviously
NEXT: Hands up if you've used the Internet recently.

---

> "If you've used the Internet recently, you've used CloudFlare and many Go programs handled your requests."
-- John Graham-Cumming, CloudFlare

---

It's definitely not a __fad__.

> "Abbiamo certamente superato la fase 'capriccio'"

---

# *__Why?__*

^ NEXT: Before we investigate why this is happening, let's look
at some terrible reasons...

---

# *"Terrible reasons"* to choose Go

* Trendy & *new* & different
* Doesn't have *semicolons*
* *Google* made it
* I 💛 *Go*
* Opensource

(I secretly think these are all great reasons)

^ Don't use these to try and convince your boss to let you do something in Go.
NEXT: I've spent a few years talking to people in to Go community...

---

# __The *top* reasons__

* *Simplicity & minimalism*
* *Productivity*
* *Performance*
* *Modern language features and excellent standard library*
* *Builds to a self-contained static binary*

^ In no particular order.
Simplicity: Obvious but Go codifies into the langauge
Productivity: Developers can quickly get stuff done 
Performance: Even when you don't optimise code, it runs quickly
Modern features: Recent language means it has more context
Static binary: Again simple, binary with no other dependencies is dev ops dream

---

# __Simplicity & minimalism__

![inline](https://www.dropbox.com/s/j41gqp8tpshhl00/minimalist.jpg?dl=1)

---

# __Simplicity & minimalism__

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

# __Diversity__

> "I think Go can be a great first language to learn due to its simplicity and minimalism. This greatly lowers the barrier for people generally excluded from server-side programming."
-- Carlisia Pinto, Founding member of GoBridge

![inline](https://www.dropbox.com/s/6oe14az5j5aizgu/carlisia.jpg?dl=1)

^ Community decided to maximize diversity (for selfish reasons)
Easy to get started helps there

---

# __Accessibility__

> "Go's simplicity means it's very easy to read or hear; that's especially important to me"
-- Parham Doustdar

![inline](https://www.dropbox.com/s/cy16anzo5ca8t49/parhamdoustdar.jpeg?dl=1)

* Blind Go programmer from Tehran, Iran, who uses a screen reader

^ We glance at code to understand it, if it has to be read out
you can see why not having semi-colons is pretty attractive.

---

# __Productivity__

![inline](https://www.dropbox.com/s/kwcvvtmrd6x7a2a/working.jpg?dl=1)

^ Could argue it's the most important?

---

# __Productivity__

* Getting stuff done
* Solve problems quickly
* Quickly be useful to a new team
* Join open-source projects with low friction
* Low investment: don't mind trying things or even throwing bits away

^ `gofmt` all Go code looks the same
Don't mind throwing things away because you haven't invested all this effort up-front.
Leads to feeling more free to be creative

---

# __Performance__

![inline](https://www.dropbox.com/s/vf0k2y9lef9bbh6/salmon.jpg?dl=1)

---

# __Performance__

> "We had a spike of about 2.4 million requests within ten minutes, and forgot to turn caching on! Luckily, since it only touched our Go code, it didn’t matter; we didn’t crash."
-- Dean Elbaz, SongKick

^ Without much effort, Go will perform great.
Doesn't mean you don't have to worry about it - but it's nice to have a headstart.
Compiles quick too (build and test on save)

---

# __Modern language features__
# __(and an excellent standard library)__

![inline](https://www.dropbox.com/s/9as6scdnjh5b0he/apples.jpg?dl=1)

^ Lets hear it for the std lib.
Go is a modern language, and was designed within the context of how we ended up writing and deploying software.
Changed a lot in last twenty years.

---

# __Concurrency__

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

# *`net/http`* __package__

> "The http package's comprehensiveness has made building and running a load of http servers super easy, which given we'd picked a service oriented architecture definitely made that faster."
-- Matt Heath, Monzo

![inline](https://www.dropbox.com/s/kb842aurrwxnq92/heath.jpg?dl=1)

^ Standard library is awesome. Built by world class engineers.
Extremely high standards.
World class HTTP capabilities: Everything you need to build web services

---

# __Builds to a self-contained static binary__

![inline](https://www.dropbox.com/s/27nz6md0yayky8a/cardboard-box.png?dl=1)

^ You end up with a single native binary
Contains EVERYTHING it needs to run
NEXT: Great for deployment

---

# __Builds to a self-contained static binary__

> "Deployments are extremely simple when you just have a self-contained binary without having to worry about special runtimes, package managers, external dependencies, proxies, or anything like that."
-- Ernesto Jiménez, [Slackline.io](https://Slackline.io)

![inline](https://www.dropbox.com/s/ebbpgpfxv54dklz/ernesto.png?dl=1)

^ `go build` produces an executable binary whose only dependency is the operating system
Other languages have specific dependencies, usually a specific version
It means the binaries can be chunky - but in modern situations, it's the right choice.

---

# How to *introduce* Go at work

![inline](https://www.dropbox.com/s/gcygq6obw9f4lft/skeletons.jpg?dl=1)

^ What if you want to introduce Go but are getting resistence?

* Is your team *resistant*?
* Too busy to "*try something new*"
* "It's too *risky*"

---

# [fit] If you want to introduce Go to your team
# [fit] *build something* small.

---

# *__Build something small__*

* *Solve* a little annoying problem for your team
* Do a short *presentation* after stand-up
* Show the code, *talk* about how it works
* *Invite others* to play with it, and even add a feature
* Put it on GitHub and ask for *reviews* too
<sub>Check out the `#reviews` channel on `gophers.slack.com`</sub>

^ Build something in your spare time
Pick something small, that's not on the roadmap

---

# Challenge

> __Introduce three new people to Go__

1. Existing programmer friend 
1. Someone you work with now (maybe your boss?)
1. Younger family member who has never coded before

---

# __Challenge ideas__

* Send them your favourite *Go video*
* Show them some *code you've written*
* Go through *[https://tour.golang.org/](https://tour.golang.org/)* with them

---

# *__Thank you__*

* __Twitter__ *`@matryer`*
* __gophers.slack.com__ *`matryer`*
* __Blog__ *`matryer.com`*

<sub>__Special thanks to__ Dave Cheney, Parham Doustdar, Dean Elbaz, John Graham-Cumming, Matt Heath, David Hernández, Ernesto Jiménez, Bill Kennedy, Brian Ketelsen, Carlisia Pinto, Alex Rodríguez</sub>
