# Interface in Go
### By Jay Gandhi
- <b>Bill Kennedy</b>
  - Polymorphism is the ability to write code that can take on different behavior through the implementation of types. Once a type implements an interface, an entire world of functionality can be opened up to values of that type.
  - Interfaces are types that just declare behavior.
    - This behavior is never implemented by the interface type directly, but instead by user-defined types via methods. When a user-defined type implements the set of methods declared by an interface type, values of the user-defined type can be assigned to values of the interface type.
    - This assignment stores the value of the user-defined type into the interface value.
  - Interfaces are types that just declare behavior. This behavior is never implemented by the interface type directly, but instead by user-defined types via methods.
  - When a user-defined type implements the set of methods declared by an interface type, values of the user-defined type can be assigned to values of the interface type. This assignment stores the value of the user-defined type into the interface value.
  - If a method call is made against an interface value, the equivalent method for the
 stored user-defined value is executed.
  - Since any user-defined type can implement any interface, method calls against an interface value are polymorphic in nature.
  - The user-defined type in this relationship is often called a concrete type, since interface values have no concrete behavior without the implementation of the stored user-defined value.

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
- <b>The Go Programming language by Donovan and Kernighan.</b>
  - Interface types express generalizations or abstractions about the behaviors of other types.
    - By generalizing, interfaces let us write functions that are more flexible and adaptable because they are not tied to the details of one particular implementation.
  - Many object-oriented languages have some notion of interfaces, but what makes Go's interfaces so distinctive is that they are SATISFIED IMPLICITLY.
  - In other words, there's no need to declare all the interfaces that a given CONCRETE TYPE satisfies; simply possessing the necessary methods is enough.
  - This design lets you create new interfaces that are satisfied by existing CONCRETE TYPES without changing the existing types, which is particularly useful for types defined in packages that you don't control.
  - All the types we've looked at so far have been CONCRETE TYPES. A CONCRETE TYPE specifies the exact representation of its values and exposes the intrinsic operations of that representation, such as arithmetic for numbers, or indexing, append, and range for slices.
  - A CONCRETE TYPE may also provide additional behaviors through its methods. When you have a value of a CONCRETE TYPE, you know exactly what is IS and what you can DO with it.
  - There is another kind of type in Go called an INTERFACE TYPE. An interface is an ABSTRACT TYPE. It doesn't expose the representation or internal structure of its values, or the set of basic operations they support;
  - It reveals only some of their methods. When you have a value of an interface type, you know nothing about what it IS; you know only what it can DO, or more precisely, what BEHAVIORS ARE PROVIDED BY ITS METHODS.

  ```
  type ReadWriter interface {
    Reader
    Writer
  }

  // This is called EMBEDDING an interface.
  ```
  - A type SATISFIES an interface if it possesses all the methods the interface requires.
  - Conceptually, a value of an interface type, or **INTERFACE VALUE**, has two components, a **CONCRETE TYPE** and a **VALUE OF THAT TYPE**. These are called the interface's **DYNAMIC TYPE** and **DYNAMIC VALUE**.
  - For a statically typed language like **Go**, types are a compile-time concept, so a type is not a value. In our conceptual model, a set of values called **TYPE DESCRIPTORS** provide information about each type, such as its name and methods. In an interface value, the type component is represented by the appropriate type descriptor.

  ```
  var w io.Writer
  w = os.Stdout
  w = new(bytes.Buffer)
  w = nil

  var w io.Writer
  w
  type: nil
  value: nil

  w = os.Stdout
  w
  type: *os.File
  // value: the address where a value of type os.File is stored

  w = new(bytes.Buffer)
  w // type: *bytes.Buffer
  // value: the address where a value of type bytes.Buffer is stored

  w = nil
  w
  type: nil
  value: nil
  ```
