# Struct in GO

### Definition

- A struct is a sequence of named elements, called fields, each
of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (AnonymousField). Within a struct, non-blank field names must be unique.
- A field declared with a type but no explicit field name is an anonymous field, also called an embedded field or an embedding of the type in the struct. An embedded type must be specified as a type name T or as a pointer to a non-interface type name *T, and T itself may not be a pointer type. The unqualified type name acts as the field name.

### Go is object oriented




- Encapsulation
  - state ("fields")
  - behavior ("methods")
  - exported / un-exported

- Reusability
  - inheritance ("embedded types")

- Polymorphism
  - interfaces

- Overriding
  - "promotion"

### Traditional OOP

- Classes
  - data structure describing a type of object
  - you can then create "instances"/"objects" from the
  - class/blue-print
  - classes hold both:
  - state / data / fields
  - behavior / methods
  - Public / private

### Inheritance in GO

In Go:
  - you don't create classes, you create a type
  - you don't instantiate, you create a value of a type
