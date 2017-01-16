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

# Contents

* Why I love Go
* Success stories from around the community
* How to get Go into your company

---

# Terrible reasons to choose Go

* Trendy & different
* Doesn't have semicolons
* Google made it
* We love it
* Opensource

(I secretly think these are all great reasons)

---

# Go is growing

![inline](https://www.dropbox.com/s/kp569ay6hxi6t3a/golang-trends-4.png?dl=1)

<sub>from Google Trends [https://www.google.co.uk/trends/explore?q=golang](https://www.google.co.uk/trends/explore?q=golang)</sub>

2017 will see more groups, conferences, meet-ups, projects.

---

# Who's using Go?

Adjust, Basecamp, BBC, Bit.ly, Bitbucket, Buzzfeed, Booking.com, Brightcove, CloudFlare, CoreOS, DailyMotion, DigitalOcean, Dell, Docker, DropBox, Google, GitHub, GrayMeta, eBay, Embedly, Facebook, GOV.UK, Intel, Medium, Monzo, Netflix, NotOnTheHighStreet, Outdoorsy, Pivotal, SoundCloud, Segment, SendGrid, SongKick, SpaceX, Trenìt!, Twitch, Uber, Yahoo, __and many more... I got bored typing them[^1]__

[^1]: For a more complete list, see [https://github.com/golang/go/wiki/GoUsers](https://github.com/golang/go/wiki/GoUsers)

---

# Why?

---

# Why choose Go?

* Tiny language
* A dev team can learn the basics extremely quickly
* Designed for modern programming
* Concurrency is finally easy
* Builds to a single static binary

---

# All Go code looks the same?

`gofmt` tool provides standard formatting

* Code feels familiar
* No steep learning curve when joining new projects

---

# Is it because of how easy concurrency is?

```go
for i := 0; i < 1000000; i++ {
	go doWork(i)
}
```

* Goroutines are cheap vs expensive threads
* Built-in channel types let Goroutines communicate
* New way of thinking about software

---

# Simplicity

---

> "I think Go can be a great first language to learn due to its simplicity and minimalism. This should greatly lower the barrier for people generally excluded from server side programming."
-- Carlisia Pinto, Founding member of GoBridge

---

![](https://www.dropbox.com/s/j41gqp8tpshhl00/minimalist.jpg?dl=1)

---

# Meet Parham Doustdar

* Blind Go programmer from Tehran, Iran
* Uses a screen reader to read code
* "Go's simplicity means it's easy to read; that's especially important to me"

---

# ![inline](https://www.dropbox.com/s/usjg3jsyf2c9yb9/monzo-card.png?dl=1) Monzo

A bank - written in Go. 

* "`net/http` package's comprehensiveness makes running a load of servers super easy which given we'd picked a service oriented arch definitely made that faster"[^2]
* "Static linking so we can deploy our services in scratch containers without even a shell"[^2]

[^2]: Matt Heath, January 2017 - monzo.com

---

# Selling Go to your boss

---

# Community

---

# Three challenges

---

# Challenge #1

## Introduce three new people to Go

1. Existing programmer friend 
1. Someone you work with now (maybe your boss?)
1. Younger family member who has never coded before

Show them something, tell them about it, send them a video, maybe go through [https://tour.golang.org](https://tour.golang.org) with them?

---

![inline](https://www.dropbox.com/s/aesvuocnqcsfkg6/ethan-coding2.jpg?dl=1)

Ethan Edwards wrote a program to figure out how old he was

---

# Challenge #2

If you don't use Go at work, solve a real problem using Go in your spare time

Then show your team

---

# Challenge #3

Write a blog post showing something simple in Go and share it with your network

Tweet me @matryer for a retweet

---

# Thank you

Mat Ryer
@matryer