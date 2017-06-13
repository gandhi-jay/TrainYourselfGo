# Interface in Go
### By Jay Gandhi
- <b>Bill Kennedy</b>
  - Polymorphism is the ability to write code that can take on different behavior through the implementation of types. Once a type implements an interface, an entire world of functionality can be opened up to values of that type.
  - Interfaces are types that just declare behavior.
    - This behavior is never implemented by the interface type directly, but instead by user-defined types via methods. When a user-defined type implements the set of methods declared by an interface type, values of the user-defined type can be assigned to values of the interface type.
    - This assignment stores the value of the user-defined type into the interface value.

- [<b>Tomas Senart</b>](https://youtu.be/xyDkyFjzFVc)
  - <b>Single Responsibility Principal</b>
    - The single responsibility principle is a computer programming principle that states that every module or class should have responsibility over a single part of the functionality provided by the software, and that responsibility should be entirely encapsulated by the class.
  - <b> SEPARATE ALL CONCERNS </b>
    - Embrace interface in golang.
    - If you want to reduce the possibilities of unexpected interaction, you need to reduce between all of these orthogonal concerns we need to reduce scope of our interface.
    - We need to make them small and precise.
    - Single method interface are most expressive and composable tool of abstraction.
    - If you don't do this then you're violating the second principle that is ISP (Interface Segregation Principle)
      - It states that clients should not be forced to implement interfaces they don't use. Instead of one fat interface many small interfaces are preferred based on groups of methods, each one serving one submodule.
    - If you don't follow ISP, your interface will be huge and people will have access to the functionality they don't need or you don't want them to use.
  - <b>Benefits of ISP.</b>
      - In go, they coexist with function as elegent and very powerful.

      ```
        // A client sends http.requests and returns http.response or error in case of failures.

        type Client interface {
          Do(*http.Request) (*http.Response, error)
        }

        // ClientFunc is a func that impements the client interface
        // Possible because in go any type can have method attached to it.

        type ClientFunc func(*http.Request) (*http.Response, error)

        func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
          return f(r)
        }
      ```
      - This is possible in go, any type can have method attached to it which enables function types like this to implement functional interface(single method interface).

  - <b>Decorator Pattern</b>
    - Decorator is pattern for adding functionality (i.e. logging, authentication) to another function by wrapping it.
    - Allows behavior to be added to an instance of type without affecting behavior other instances of the same type.

    ```
    // A decorator wraps a client with extra behavior.

    type Decorator func(Client) Client
    ```
    - In golang, it's easily expressed that you have function, which takes a client or any interface and extends it and return the same type.
    - This type within the function has been extended, it's behavior has been enriched.
    - This pattern helps to confirm to both Single responsibility pattern third Principle is Open/Close Principle.
    - Open/Close Principle
      - Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.
    - In clear words, Entity can allow its behavior to be extended without modifying it's source code.
    - In golang, it's really simple. In Java/C++ quite a lot code to do.
    - Let's add log into a client without modifying in source code.
    ```
    // Logging return a Decorator that logs a Client's requests

    func Logging (l *log.Logger) Decorator {
      return func (c Client) Client {
        return ClientFunc(func(r *http.Request) (*http.Response, error){
            l.Printf("%s: %s %s", r.UserAgent, r.Method, r.URL)
            return c.Do(r)
          })
      }
    }
    ```
