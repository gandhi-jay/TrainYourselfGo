# Interface in Go
### By Jay Gandhi
- <b>Bill Kennedy</b>
  - Polymorphism is the ability to write code that can take on different behavior through the implementation of types. Once a type implements an interface, an entire world of functionality can be opened up to values of that type.
  - Interfaces are types that just declare behavior.
    - This behavior is never implemented by the interface type directly, but instead by user-defined types via methods. When a user-defined type implements the set of methods declared by an interface type, values of the user-defined type can be assigned to values of the interface type.
    - This assignment stores the value of the user-defined type into the interface value.
- [<b>Tomas Senart</b>](https://youtu.be/xyDkyFjzFVc)
  - Single Responsibility Principal
    - The single responsibility principle is a computer programming principle that states that every module or class should have responsibility over a single part of the functionality provided by the software, and that responsibility should be entirely encapsulated by the class.
  - <b> SEPARATE ALL CONCERNS </b>
    - Embrace interface in golang.
    - If you want to reduce the possibilities of unexpected interaction, you need to reduce between all of these orthogonal concerns we need to reduce scope of our interface.
    - We need to make them small and precise.
    - Single method interface are most expressive and composable tool of abstraction.
    - If you don't do this then you're violating the second principle that is ISP (Interface Segregation Principle)
      - It states that clients should not be forced to implement interfaces they don't use. Instead of one fat interface many small interfaces are preferred based on groups of methods, each one serving one submodule.
    - If you don't follow ISP, your interface will be huge and people will have access to the functionality they don't need or you don't want them to use.
    - Benefits of ISP.
      - In go, they coexist with function as elegent and very powerful.
      ```
        // A client sends http.requests and returns http.response or error in case of failures
        type Client interface {
          Do(*http.Request) (*http.Response, error)
        }
        // ClientFunc is a func that impements the client interface
        type ClientFunc func(*http.Request) (*http.Response, error)
      ```
