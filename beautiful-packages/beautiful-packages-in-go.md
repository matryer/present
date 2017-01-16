footer: Writing Beautiful Packages in Go - Mat Ryer
slidenumbers: true

# *Beautiful packages in Go*
# *Mat Ryer*

![](https://www.dropbox.com/s/dvxihxqkmbs5098/london2.jpg?dl=1)

---

# *Mat Ryer*

Coding Go since before v1

BLOG *matryer.com*

BOOK *Go Programming Blueprints: Second Edition*

WORK *graymeta.com* - 100% Go

Not Banksy

---

# *Available now*

![inline](https://www.dropbox.com/s/ekipavlfd9f45jd/bookcover-2nd.png?dl=1)

---

# *What is a package?*

---

# *What is a package?*

*Folder* full of `.go` (and hopefully `_test.go`) files

Can be *imported* into other projects

Has *exported* (public) and *unexported* (internal) stuff

Hopefully open-sourced

Not `package main` (that's a command)

---

# *Examples*

Standard library is a load of packages

`fmt`, `log`, `http`, `os`, `filepath`, `errors`, `testing`

Open source

`glog`, `testify/assert`, `pkg/errors`, `gorilla/mux`, `protobuf/proto`

---

# *What is a beautiful package?*

---

# *What is a beautiful package?*

Elegant & simple

Obvious (maybe even boring)

You already know how to use it

Look forward to using it

---

# *SIMPLE!*

---

# *How can we make our packages beautiful?*

---

# *User centred design*

> "the needs, wants, and limitations of end users of a product, service or process are given extensive attention at each stage of the design process"


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

# Some specific things...

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

# *Create single method interfaces*

If you're going to add an interface, try and keep them as tiny as possible

```go
// EaseScore returns a score for how easy an interface
// is to implement. Lower is better.
func EaseScore(i Interface) int {
	return len(i.ExportedMethods)
}
```

If an interface has *one* method, it's very easy to implement.

---

...and you can do things like this:

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

...and even:

```go
type StatusHandler int

func (s StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := int(s)
	w.WriteHeader(code)
	io.WriteString(w, http.StatusText(code))
}
```

---

# *SIMPLE!*

---

# *Context as the first argument*

`context.Context` is part of the standard library

- Pass as the first argument
- If your package performs cancellable tasks
- If your code calls code that takes a context

---

# *Quiz*

In which version of Go did `context` join the standard library?

---

# *Go 1.7*

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
```

and they know what's going on

---

# *You are probably the first user*

---

# *Test driven development*

Use the package before it even exists

Try different things

Aware of the API footprint from the very beginning

TIP: Put test code in different package

TIP: Use `godoc` early too

---

# *3 hardest things...*

The 3 hardest things in coding are:

1. Naming things

2. Invalidating the cache

3. Invalidating the cache

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
	fmt.Println(g.Greet("David Hernandez"))
}
// output: Hi David Hernandez
```

or

```go
func main() {
	g := Greeter{
		Format: "Hey there, %s",
	}
	fmt.Println(g.Greet("David Hernandez"))
}
// output: Hey there, David Hernandez
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

# Don't export interfaces for the sake of it

* I love interfaces - but I love smaller APIs even more

```go
type Greeter struct {}

func (g *Greeter) Greet(name string) string {
	return "Hello "+name
}
```

* The user can always make their own interface

```go
type Greeter interface {
	Greet(name string) string
}
```

---

# *Follow conventions*

Be similar to the standard library and other popular packages

Don't surprise people

Be obvious, not clever

---

# *Quiz*

What should you be instead of clever?

---

# *Obvious*

---

# *SIMPLE!*

---

# *Shameless*

![inline](https://www.dropbox.com/s/ekipavlfd9f45jd/bookcover-2nd.png?dl=1)

---

# *Computers can help too*

[github.com/alecthomas/gometalinter](github.com/alecthomas/gometalinter)

[goreportcard.com](github.com/alecthomas/gometalinter)

---

# How to spot a quality package

* Did they use `go fmt`?
* Are there `_.test.go` files?
* Does the code read well? Do the names sound weird?
* Have they just ported from a different language? (Are they using Java names?)
* How many dependencies does it have?

---

Go fourth and write beautiful packages

---

# *Cheers*

# *@matryer*
