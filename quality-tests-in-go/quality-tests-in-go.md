footer: Quality Tests in Go - Mat Ryer
slidenumbers: true
autoscale: true

# Quality Tests in Go

Mat Ryer
Gotham Go - New York City
November 18th 2016

![](https://www.dropbox.com/s/tpotnhdrczr7v94/new-york-manhattan.jpg?dl=1)

^ Red soks

---

# Mat Ryer

️⌨ Coding Go since before v1

🔗 *matryer.com*

📚 Go Programming Blueprints: Second Edition

👷 *graymeta.com* - 100% Go backend

🖌 Not Banksy

^ Blog at matryer.com

^ Written book

^ All new chapters
microservices with gRPC and protocol buffers
Making these questions up?

---

# Wayne Manor

![inline](https://www.dropbox.com/s/py0m913xf85nn91/brucemanor.jpeg?dl=1)

Wollaton Hall, Nottingham, England

^ Gotham Go - this is Batman's country escape

^ Wayne Manor in movies with Christian Bale

^ Actually in Nottingham - where I'm from

^ My link to Gotham and Batman

---

# Purpose of this talk

- Explain what testing is and why __the best projects in the world do it__
- Convince you to write tests __first__
- Look at some tips for writing quality tests

^ Who writes tests?
Who writes TDD?

---

# If you don't write test code, you're either very brave or

^ or something else

---

# What is a test?

```go
// in batify.go
func Batify(item string) string {
	return "Bat " + item
}

// in batify_test.go
func TestBatify(t *testing.T) {
	actual := Batify("Car")
	if actual != "Bat Car" {
		t.Errorf("expected Bat Car but got %s", actual)
	}
}
```

^ If we have some code that does something

^ (all the best code does something)

^ You have more code, that validates the other code behaves as expected

^ IF we do something to mess up Batify function, this test will fail altering us to that fact

---

# Why write tests?

- Get computers to check our work
- Run them automatically (every time we change something)
- End up with a suite of tests describing the functionality
- Bold and robust refactoring
- We know if we break something by mistake
- We know __what__ we've broken

^ We all do manual testing - takes time and effort

^ Robotic butler

---

# Why write tests?

- Demonstrate the intent
- Code becomes documentation

^ If you want to know how a function in the standard library works, you can check its tests

---

# Test driven development (TDD)

TDD is a development process where you write the test code first.

=

> How can you test something that doesn't exist?
-- Confused Person

^ In non-tech world, you can see why this is confusing

^ Asked to test something that doesn't exist

^ In real world: Can you test this bicycle for me?

^ But this is NOT the real world... software world

---

# Test driven development (TDD)

1. Write a small piece of test code
1. Run the test to see it fail
1. Make it pass

^ Write test code.
Error or fail.
A test that cannot fail, says nothing about your system (e.g. assert true true)
Do minimum amount of work

---

# Test driven development (TDD)

- Evidence based programming
- Prove bugs or missing features
- Protect from regression
- Write tests for each other
- Get it out of your brain
- Be your own user first

---

# Does Test driven development mean we don't have to do up-front design?

^ Design can evolve, but be intelligent

---

Does Test driven development mean we don't have to do up-front design?

# [fit] No

---

Does Test driven development mean we don't have to do up-front design?

# No

Don't skip the design step

TDD is an implementation time activity

---

# [fit] Quality tests

---

# Different levels of tests

```
            *       *             * Unit *
            *       * Integration * Unit *
            *       * ___________ * Unit *
            *  End  *             * Unit *
            *  to   * Integration * Unit *
            *  end  * ___________ * Unit *
            *       *             * Unit *
            *       * Integration * Unit *
            *       *             * Unit *
```

^ E2E: Touch all parts of the system
Integration tests usually check the bounaries between components
Unit tests (favourite) are at function/method level

---

# I 💙 unit tests

"Can I create an account and send money to my friend?"

vs

"When I sign in, am I given an auth token?"

^ If the auth token generator is borked, both tests will fail
but which will point me in the right direction?
Which helps me diagnose and fix the issue?

^ If top test fails, it could be because user accounts is broken
or transaction service etc.

---

# Unit tests :)

*Unit tests* are tiny, simple and laser focussed

- If something breaks, they point directly to it

# End to end tests :(

If __anything__ breaks, the end-to-end test will fail

---

# I 💙 table tests

```go
func TestBatify(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"Car", "Bat Car"},
		{"Bike", "Bat Bike"},
		{"Backpack", "Bat Backpack"},
	}
	for _, test := range tests {
		t.Run(test.out, func(t *testing.T) {
			actual := Batify(test.in)
			if actual != test.out {
				t.Errorf("expected %s, but got %s", test.out, actual)
			}
		})
	}
}
```

---

# I 💙 table tests

- Anonymous `struct` slice for each case
- One `for...range` loop iterating over each
- Run each case using `t.Run`
- Use a sensible value for the name, or add one (e.g. `map`)

---

# Setup and teardown

```go
func TestBatify(t *testing.T) {

	// TODO: setup
	// TODO: defer teardown

	for _, test := range tests {
		t.Run(test.out, func(t *testing.T) {

			// TODO: per-test setup
			// TODO: defer per test teardown

		})
	}
}
```

^ ...sometimes you want setup/teardown for non-table tests

---

# Shared setup and teardown

```go
func setup(t *testing.T) (*db, func()) {
	db, err := dial(testDatabase)
	if err != nil {
		t.Errorf("dial: %s", err)
		return nil, func(){}
	}
	return db, func(){
		if err := db.Close(); err != nil {
			t.Errorf("db.Close: %s", err)
		}
	}
}
```

* Return a teardown func so setup and teardown code is kept together

---

```go
func TestStuff(t *testing.T) {

	db, teardown := setup(t)
	defer teardown()

	// use the db

}
```

---

# Different package name

```go
// in batify.go
package batify

func Batify(item string) string...
```

```go
// meanwhile...
// in batify_test.go
package batify_test

import "batify"

func TestBatify(t *testing.T) {
	actual := batify.Batify("Car")
	//...
}
```

^ Usually in Go it's an error to have two packages in the same folder
Exception is test code.

^ `batify.go` has `batify` package
Meanwhile batify_test has different package name
We import batify package and use it externally 

---

# Different package name

- Can only test the public API
- IDEs will reveal the footprint of your package
- You will use your package name (avoid `tea.BrewTea` or `tea.TeaBrewer`)

> If you don't trust yourself to not use private methods in unit tests, then I cannot help you
-- Dave Cheney

^ Cannot fiddle with package internals, you use the package like a real user
IDEs: Autocomplete - hit dot and see list of exported things
Use package name in code

---

# Different package name

```go
package batify

func Item(item string) string {
	return "Bat " + item
}

// now we can call batify.Item
// instead of batify.Batify
```

^ In our Batify example, we might rename the Batify function to Item

---

# Interface tests[^1]

```go
type Store interface {
	Get(id string) (Storable, error)
	Put(v Storable) (string, error)
}

func TestStore(t *testing.T, s Store) {
	obj := someTestObject{}
	id, err := s.Put(obj)
	if err != nil {
		t.Errorf("Put: %s", err)
	}
	obj2, err := s.Get(id)
	if err != nil {
		t.Errorf("Put: %s", err)
	}
	if obj.ID() != obj2.ID() {
		t.Errorf("object IDs don't match")
	}
}
```

[^1]: For a real example, see [github.com/graymeta/stow](github.com/graymeta/stow) - see `/test` folder

^ Stow github.com/graymeta/stow

^ Stow is abstraction package on top of cloud blob storage providers like S3, Azure, Google Cloud Storage, etc.

^ Stow has a single test that all implementations must pass in order to be accepted into the library.

---

# Test code is code

- Treat it as a first class concern
- Look after it
- Clear away old tests
- Know your test codebase intimately

> Don't be afraid to write your own test harness helpers
-- Eleanor McHugh

^ "more tests the better" doesn't work

---

# [fit] Top tips from the community...

I asked people on __Gophers Slack__ for their top testing tips.

=

Btw, your invite is waiting at: [https://invite.slack.golangbridge.org]()

---

# [fit] TIP: Code coverage

> Shoot for 90% code coverage, 100% of the happy path
-- William Kennedy (Go in Action)

```
go test -cover
```

```
go test -covermode=count -coverprofile=count.out
go tool cover -html=count.out
```

^ Bill Kennedy author of Go in Action

^ Code coverage refers to how much of your code is touched when your tests run

^ Happy path is the expected flow, ignoring error checks (although sometimes they're important)

---

# btw, with test driven development, code coverage is automatically __very__ high

^ You don't have to keep going back trying to figure out how to test things

---

# TIP: Inject your dependencies

```go
func TestSignInNotFound(t *testing.T) {
	users := &mockedUserService{
		GetUserFunc: func(email string) (*User, error) {
			return nil, ErrNotFound
		},
	}
	srv := Server{
		UserService: users,
	}
	token, err := srv.SignIn("nope@me.com", password)
	if err != ErrNotFound {
		t.Errorf("ErrNotFound error found")
	}
}
```

from Ernesto Jimenez

---

# Mocking

```go
type UserService interface {
	GetUser func(email string) (*User, error)
}

// meanwhile...

type MockedUserService struct {
	GetUserFunc func(email string) (*User, error)
}

func (m *MockedUserService) GetUser(email string) (*User, error) {
	return m.GetUserFunc(email)
}
```

Automagically turn interfaces into mocked structs with *github.com/matryer/moq*

---

# TIP: Pass `testing.T` into helpers

```go
func callBatifyAPI(t *testing.T, item string) string {
	// TODO: make a request
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("the request failed" %s", err)
	}
}
```

from Coleman McFarland

---

# _Should I use a testing framework?_

Whatever makes you happy. Whatever makes you test.

- Always better to avoid dependencies if you can
- Don't give new team members something else to learn

^ I like assert.Equal pattern - easier to read and uses less code
But IS something else for newcomers to learn

^ You have to decide what's right for your project

---

# Finally...

- Women who Go - [womenwhogo.org](http://www.womenwhogo.org)
- Go Bridge - Building Bridges That Educate & Empower Underrepresented Communities [golangbridge.org](https://golangbridge.org)

**You** are welcome in the Go community; we want their loss to be our gain.

^ If you've ever felt excluded from a community, come and join us.

---

# [fit] *Thank you*

![inline](https://www.dropbox.com/s/ekipavlfd9f45jd/bookcover-2nd.png?dl=1)

__Preorder now__ Go Programming Blueprints: Second Edition

Tweet me: `@matryer`

Find me after if you have any questions.