# [fit] *Writing Beautiful Packages in Go*

Mat Ryer
@matryer

Golang UK Conference
London, August 2017

![](https://www.dropbox.com/s/rju5s6e3rffqq8n/londonoptimised.jpg?dl=1)

---

# *Who is Mat Ryer?*

![inline](https://www.dropbox.com/s/5a6z48ul8g13em6/ThisIsMat.JPG?dl=1&egg=true)

---

# *What is a package / library?*

*Folder* full of `.go` (and hopefully `_test.go`) files

Can be *imported* into other projects

Has *exported* (public) and *unexported* (internal) stuff

Hopefully open-sourced

Not `package main` (that's a command)

---

# *What is a beautiful package?*

Elegant & simple

Obvious (maybe even boring)

You already know how to use it

Look forward to using it

---

# *Motivation*

Do you have the right motivation for writing a package?

Don't write packages for the sake of it

Are you excited to write it or does it sound boring?

If you aren't solving a problem you have, don't bother

---

# *Keep packages internal to start with*

Don't automatically open-source your packages

Live with them first internally

But design them as if you will one day

---

# *User centred design*

> "the needs, wants, and limitations of end users of a product, service or process are given extensive attention at each stage of the design process"
-- Wikipedia


Personas, Scenarios, Use cases

---

# *User centred design*

Packages have an interface (API)

They're used by humans

-

Why don't we apply the same kind of thinking?

---

# *Ask yourself...*

Who is the user of the package?

What are they trying to do?

Why are they doing it?

Why are they using your package?

Are they building a PoC? Or is this part of a bigger system?

---

# [fit] "I like specific things..."

![](https://www.dropbox.com/s/b9w8agkm9az4acx/SpitalfieldsMarket.jpg?dl=1)

---

# [fit] *Smaller footprints are better*

![inline](https://www.dropbox.com/s/tq0y5b2o2j2lkn9/package-footprint.png?dl=1)

---

# *Narrow types are simpler*

```go
func WriteJSON(v interface{}, f *os.File) error
```

Can only write to file

-

```go
func WriteJSON(v interface{}, w io.Writer) error
```

Now I can write to...

---

# `io.Writer`

`os.File`
`bytes.Buffer`
`http.ResponseWriter`
`encoding/zip`

-

and whatever types our users have created

---

# *Seek out single method interfaces*

If an interface has *one* method, it's very easy to implement

It will be used more

---

...and your users can do cute things like this:

```go
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return h(w, r)
}
```

---

...and even

```go
type StatusHandler int

func (s StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := int(s)
	w.WriteHeader(code)
	io.WriteString(w, http.StatusText(code))
}

var (
	NotFoundHandler       = StatusHandler(404)
	NotImplementedHandler = StatusHandler(501)
	NotLegalHandler       = StatusHandler(451)
)
```

---

# *Structure your code properly*

People will look at your source if they're thinking of using your package

* Give them a warm welcome, and make them feel at home
* Glance at directory listing, and go from there

> "The first place I look when considering using a package are the test files. If there aren't any, that tells me something. Otherwise, I can quickly see how to use it without any documentation."
-- David Hernandez

---

# *Subfolders*

---

# *Listen to @rakyll*[^1]

* Use multiple `.go` files
* Keep types close to where they're used
* Organise by responsibility (`User` type goes into `users.go`)
* Optimise for `godoc` and provide examples

[^1]: via @rakyll at [https://rakyll.org/style-packages/](https://rakyll.org/style-packages/)

---

# *Leave concurrency to the user*

Avoid this:

```go
package thing

func DoAmazingThing() {
	go startDoingSomethingAmazing()
}
```

User won't necessarily know it's going to do this.

---

If they want to, the user can do this:

```go
go thing.DoAmazingThing()
// do other stuff at the same time...
```

and they know what's going on

And they can use it in a blocking way too:

```go
thing.DoAmazingThing()
// after amazing thing
```

---

# *You are probably (hopefully) the first user of your package*

Force yourself to see things from the user's point of view...

---

# *Test driven development*

Use the package before it even exists

Try different things

Aware of the API footprint from the very beginning

---

# *Follow conventions*

Be similar to the standard library and other popular packages

Don't surprise people

Be obvious, not clever

---

# *3 hardest things...*

The 3 hardest things in coding are:

1. Naming things

2. Invalidating the cache

3. Invalidating the cache

PROPER JOKE

---

# *3 hardest things...*

The 3 hardest things in coding are:

1. Naming things

2. Off by one errors

---

# *Naming things*

Remember the package name is part of the API

```go
package tea

func BrewTea(steeptime time.Duration) error {
	//...
}
```

Not nice from the outside:

```go
err := tea.BrewTea(1 * time.Minute)
```

---

```go
package tea

func Brew(steeptime time.Duration) error {
	//...
}
```

Means we can call:

```go
err := tea.Brew(1 * time.Minute)
```

---

# *Expose yourself to the API footprint from the beginning*

* Put test code in different package

```go
package proj_test

func Test(t *testing.T) {
	g := proj.FormatGreeter("Hi %s")
	g.Greet("Piotr")
}
```

* Definitely do this for examples
* Use `godoc` early: `godoc -http=:8080`

---

# *Don't log stuff out*

...or at least let users turn it off

---


# *Make zero values useful*

```go
type Greeter struct {
	Format string
}

func (g Greeter) Greet(name string) string {
	if g.Format == "" {
		g.Format = "Hi %s"
	}
	return fmt.Sprintf(g.Format, name)
}
```

---

```go
func main() {
	var g Greeter
	fmt.Println(g.Greet("Dean"))
}
// output: Hi Dean
```

or

```go
func main() {
	g := Greeter{
		Format: "Hey there, %s",
	}
	fmt.Println(g.Greet("Dean"))
}
// output: Hey there, Dean
```

---

# *Avoid constructors if you can*

```go
b := tea.NewBrewer(2 * time.Minute)
```

vs

```go
b := tea.Brewer{
	SteepTime: 2 * time.Minute,
}
```

Which one is clearer?

---

# *Don't automatically add interfaces*

I love interfaces - but I love smaller APIs even more

```go
type Greeter interface {
	Greet(name string) string
}

type FormatGreeter struct {
	Format string
}

func (g FormatGreeter) Greet(name string) string {
	return fmt.Sprintf(g.Format, name)
}
```

No need for both `Greeter` and `FormatGreeter`

---

# *Another way to shrink the API footprint...*

```go
type Greeter interface {
	Greet(name string) string
}

func FormatGreeter(format string) Greeter {
	return formatGreeter{format: format}
}

type formatGreeter struct {
	format string
}

func (g formatGreeter) Greet(name string) string {
	return fmt.Sprintf(g.format, name)
}
```

---

# *Use Go-like names*

* `time.ParseTimeString` -> `time.Parse`
* `MarshalWithIndentation` -> `MarshalIndent`
* `TearDown` -> `Close`

---

# *Ask the user for `http.Client`*

If your package makes HTTP requests, ask the user to provide a `http.Client`

* Let user specify timeouts, redirect policy, proxies, etc.
* You can still default to `http.DefaultClient`
* Google App Engine has `urlfetch.NewClient` - so your package wouldn't work there

---

# *Context*

If you use context, it should be the first argument:

```go
func StartSomethingAmazing(ctx context.Context, region string, intensity int) {
	// code
}
```

---

# *Don't mess with the global state*

* Don't add flags
* Avoid `init`
* Don't mess with things in the standard library (e.g. changing `http.DefaultClient` to set a timeout)
* Importing your package shouldn't have side-effects

---

# *Subpackages are just packages*

```go
import "github.com/matryer/vice/test"

func TestTransport(t *testing.T) {
	test.Transport(t, newTestTransport)
}
```

changed to

```go
import "github.com/matryer/vice/vicetest"

func TestTransport(t *testing.T) {
	vicetest.Transport(t, newTestTransport)
}
```

---

# *Come up with a good name*

* Short
* Clear
* To the point
* Be creative, and have fun
* Not every Go project needs to mention Go

---

# *Give your project a logo or mascot*

Projects that have a logo, are 85% more likely to get used[^2]

Great way for artists to contribute to open-source projects

Not kidding...

[^2]: Totally made up statistics because we're fine with fake news now

---

# *Let's play... Name that project*

Shout out the project name that these logos belong to...

---

![inline](https://www.docker.com/sites/default/files/Whale%20Logo332_5.png)

---

![inline](https://www.docker.com/sites/default/files/Whale%20Logo332_5.png)

# [fit] *Docker*

---

![inline](http://gobuffalo.io/assets/images/logo_med.png)

---

![inline](http://gobuffalo.io/assets/images/logo_med.png)

# [fit] *GoBuffalo*

---

![inline](https://www.dropbox.com/s/p6b2f9gdawt8u3n/gorilla.png?dl=1)

---

![inline](https://www.dropbox.com/s/p6b2f9gdawt8u3n/gorilla.png?dl=1)

# [fit] *Gorilla Toolkit*

---

![inline](https://www.dropbox.com/s/e272821pcqqlbg7/travis.jpg?dl=1)

---

![inline](https://www.dropbox.com/s/e272821pcqqlbg7/travis.jpg?dl=1)

# [fit] *Travis CI*

---

![inline](https://www.dropbox.com/s/6kw62t5kszow8w7/machina.jpg?dl=1)

---

![inline](https://www.dropbox.com/s/6kw62t5kszow8w7/machina.jpg?dl=1)

# [fit] *Machine Box*

---

![inline](https://www.dropbox.com/s/d6kmix2wunxw1al/prometheus.jpg?dl=1)

---

![inline](https://www.dropbox.com/s/d6kmix2wunxw1al/prometheus.jpg?dl=1)

# [fit] *Prometheus*

---

![inline](https://www.dropbox.com/s/54ny52zs5ybrawn/nsq.png?dl=1)

---

![inline](https://www.dropbox.com/s/54ny52zs5ybrawn/nsq.png?dl=1)

# [fit] *NSQ*

---

![inline](https://www.dropbox.com/s/8ndkab5jf1prp2x/hugo-logo.png?dl=1)

---

![inline](https://www.dropbox.com/s/8ndkab5jf1prp2x/hugo-logo.png?dl=1)

# [fit] *Hugo*

---

![inline](https://www.dropbox.com/s/yyq0t4kfkh01sad/github.png?dl=1)

---

![inline](https://www.dropbox.com/s/yyq0t4kfkh01sad/github.png?dl=1)

# [fit] *GitHub*

---

# *Computers can help too*

* [github.com/alecthomas/gometalinter](https://github.com/alecthomas/gometalinter) - runs many linters

* [goreportcard.com](https://goreportcard.com/) 

---

![inline](https://www.dropbox.com/s/iu053rk9y2pd9qn/goreportcard.png?dl=1)

---

# How to spot a quality package

* Did they use `go fmt`?
* Are there `_.test.go` files?
* Does the code read well? Do the names sound weird? Does it feel like Go code?
* Have they just ported from a different language? (Are they using Java names?)
* How many dependencies does it have? (less is better)

---

# [fit] bye
