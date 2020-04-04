# Design Patterns

Notes and implementations from learning on Design Patterns.

Based on reading of Design Patterns: Elements of Reusable Object-Oriented Software.

## Contents

### [Behavioural Patterns](/behavioural)
These patterns characterize the ways that classes or objects interact and distribute responsibility.

Class patterns in this category defer parts of object creation to subclasses.
Object patterns defer object creation to a different object.

### [Creational Patterns](/creational)
These patterns concern the process of object creation.

Class patterns in this category use inheritance to describe algorithms and flow of control.
Object patterns describe how a group of objects cooperate to perform a task too large for a single object.


### [Structural Patterns](/structural)
These patterns deal with the composition of classes and objects.

Class patterns in this category use inheritance to compose classes.
Object patterns describe ways to assemble objects.

## General Notes
Some general notes from the book about design patterns.

### How Design Patterns Solve Design Problems

Design patterns help solve problems in many ways. They are listed below.

#### Finding Appropriate Objects
<details>
<summary>Expand Content</summary>

An **object** packages **data** and **operations**. These operations are only performed when an object receives a **request** (or **message**) from a **client**.

These requests are the *only* way to get an object to execute an operation and these operations are the only way to change an object's internal data (**encapsulation**).

The majority of the effort in object-oriented design is decomposing a system into objects. This is difficult due to:

1. Encapsulation
1. Granularity
1. Dependency
1. Flexibility
1. Performance
1. Evolution
1. Reusability

And so on.

All of these influence the decomposition, often in ways that conflict.

Design patterns can often surface and identify less obvious abstractions and the objects that can capture them. These objects and abstractions can often be over looked at they don't have real work counter parts.

Ex: objects that represent an algorithm or process don't typically occur in the real world yet are crucial to flexible design. Often these objects are not discovered until later in the course of the design when making a design more flexible and reusable.
</details>

#### Object Granularity

<details>
<summary>Expand Content</summary>

Objects can vary greatly in size, complexity, and number. How do we decide what should be an object? Patterns help us determine this in many different ways. From describing how to represent complete subsystems as objects, how to support massive numbers of objects at the finest granularities, to specific ways of decomposing an object into multiple smaller ones.
</details>

#### Specifying Object Interfaces

<details>
<summary>Expand Content</summary>

Each operation declared by an object specifies the operations name, the objects it takes as parameters, and the return value - IE it's **signature**.

The set of all operations declared by an object is known as it's **interface**. This interface describes the complete set of requests that can be sent to the object.

We use the term **type** to describe a particular interface. An object can have many types, it is said to have the type `Window` if it accepts all requests for the operations given by the `Window` interface.

Interfaces can contain other interfaces as subsets, a type is a **subtype** when it's interface contains the interface of it's **supertype**.

Without interfaces there is no way to know what the object can do or how to ask it to do something. An objects interface says nothing about its implementation.

When a request is sent to an object the particular operation that is performed depends on *both* the request *and* the receiving object. Different objects may have different implementations of the operations that fulfill the request.

The run-time association of a request to an object and one of it's operations is known as as **dynamic binding**. The consequence of dynamic binding is that we are not committed to a particular implementation until run-time. With this in mind we can write programs that expect an object with a particular interface, knowing that any object with the correct interface can accept the request.

Dynamic binding allows us to substitute objects that have identical interfaces at run-time, known as **polymorphism**. This allows us to make few assumptions about other objects beyond supporting a particular interface. It allows simplifying the definition of clients, decoupling objects from each other, and lets relationships vary at run-time.

Design patterns can help define interfaces by identifying their key elements and the kinds of data that get sent across an interface. They can also highlight what not to put into the interface. Further they ca specify relationships between interfaces, and often require classes to have similar interfaces or place constraints on the interfaces of some classes.
</details>

#### Specifying Object Implementation

<details>
<summary>Expand Content</summary>

An objects implementation is defined by its **class**.

Objects are created by **instantiating** a class and it is said to be an **instance** of the class. This process of instantiating the class allocates storage for the object's internal data - **instance data** - and associates the operations with this data.

New classes can be defined in terms of old classes via **class inheritance**. When a **subclass** inherits from a **parent class** it includes the definitions of all the data and operations that the parent class defines. Objects that are instances of the subclass will contain the data and operations from the parent class and the subclass.


Subclasses often **override** operations defined by its parent class in order to refine and redefine behaviours. Overriding allows subclasses a chance to handle requests instead of the parent class.

Inheritance allows defining classes by simply extending other classes, making it easy to define families of objects having related functionality.

A **mixin class** is a class that is intended to provide an optional interface or optional functionality to other classes. It is not intended to be instantiated and require multiple inheritance.

##### Class versus Interface Inheritance

It is important to understand the difference between an object's *type* and its *class*. The class defines how the object is implemented, its state and the implementation of its operations. The type only refers to the object's interface - the set of requests to which it can respond.

There is also a close relationship between the two, because a class defines the operations an object can perform, it also defines the type. Saying an object is an instance of a class, we imply the object supports the interface defined by the class.

Another importance is the distinction between class inheritance and interface inheritance (aka sub typing). Class inheritance defined an objects inheritance in terms of of another implementation, IE: it's a way to share code and representation. Inheritance inheritance describes when an object can be used in place of another.

##### Programming to an Interface, not an Implementation

Class inheritance allows us to extend an applications functionality through reuse of the functionality in the parent class. It provides a way to get new implementations almost for free, inheriting most of what you need from existing classes.

More importantly though, inheritance allows us to define families of classes with *identical* interfaces. This is key as polymorphism depends on it.

When inheritance is used carefully, all classes derived from an abstract class will share it's interface. This implies that a subclass only adds or overrides operations and does not hide operations of the parent class. All subclasses can then respond to the requests in the interface of this abstract class, making them all subtypes of the abstract class.

There are two benefits to manipulating an object solely through interfaces defined by abstract classes:

1. Clients remain unaware of the specific types of object they use, so long as the object adhere to the expected interface.
2. Clients remain unaware of the classes that implement these objects. They only know the abstract classes.

This so greatly reduces implementation dependencies between subsystems it lead to the following principal in reusable object-oriented design:

> Program to an interface, not an implementation.

Don't declare variables to be instances of particular concrete classes. Instead, commit only to an interface defined by an abstract class.
</details>

#### Putting Reuse Mechanisms to work.

<details>
<summary>Expand Content</summary>

Understanding concepts like objects, interfaces, classes, etc. is the easy part. The challenge lies in applying them to build flexible, reusable software. Design Patterns can illustrate how to do this.

##### Inheritance versus composition

Inheritance and **object composition** are the two most common techniques for reusing functionality in object-oriented software.

To reiterate, class inheritance allows for defining the implementation of one class in terms of another. Reuse by subclassing is often referred to as **white-box reuse**. This is due to the visibility the subclass has of the internals of the parent class.

Object composition is an alternative to this. In object composition new functionality is obtained by assembling (or *composing*) objects to achieve more complex functionality. Object composition requires that the objects have well-defined interfaces. This style of reuse is referred to as **black-box reuse** as the objects do not have visibility into the other objects internals.

Each method has it's advantages and disadvantages.

Class inheritance is defined statically and at compile time. It is straightforward to use as it is supported directly by the programming language. It also allows for easy modifications to the implementation being used.

However, changing the implementation inherited from the parent class at run-time is not possible. Further, parent classes often define at least part of their subclasses physical representation.

Due to inheritance exposing the internals of the parent class, it is often said that it "breaks encapsulation". The implementation of a subclass becomes so bound to the parent class, that any changes to the parent often requires a change to the subclass as well.

Implementation dependencies can also cause problems when trying to reuse subclasses, should any aspect of the inherited implementation not be appropriate for new problem domains then the parent class must be rewritten or replaced by something more appropriate.

Object inheritance, on the other hand, is defined at run-time through objects acquiring references to other objects. It also requires objects to respect other objects interfaces, which in turn requires carefully designed interfaces that do not prevent the objects use with many others. Since access is done only through interfaces, encapsulation is not broken. Any object can be replaced by another at run-time, so long as it has the same type.

Moreover, because an objects implementation will be written in terms of object interfaces, there are substantially fewer implementation dependencies.

Favouring composition over class inheritance helps keeps each class focussed. Classes and their hierarchies will remain small and will be less likely to grow into unmanageable monsters.

With object composition, you will end up with a greater number of objects (if fewer classes), and the systems behaviour will depend on their interrelationships instead of being defined in a single class.

This leads to the second principle of object-oriented design:

> Favour object composition over class inheritance.

Ideally, creating new components should not be necessary for achieving reuse. We should be able to get all functionality we need through assembling existing components through object composition. In practice this isn't always the case as the set of available components is never quite rich enough.

Reuse by inheritance then makes it easier to make new components that can be composed with old ones.

Because of the above inheritance and object composition work together.

Typically, designs can be made simpler and more reusable by depending more on object composition.

##### Delegation

In **Delegation**, two objects are involved in handling a request: a receiving object delegates operations to its **delegate**. This is analogous to a subclass deferring requests to its parent. In inheritance, an inherited operation can always use the member variable`this` to refer to the receiving object. To achieve the same in delegation, the receiver passes itself to the delegate.

The chief advantage to delegation is that is makes it easy to compose behaviours at run-time and change the way they are composed. It does have a disadvantage, and that is dynamic, highly parameterized software is hard to understand than more static software. There are also run-time inefficiencies, but the human inefficiencies are more important in the long run. As such, delegation is a good design choice only when it simplifies more than it complicates.

##### Inheritance versus parameterized types

Another mechanism for reusing functionality is through **parameterized types**, aka **generics** and **templates**. This allows for specifying a type without specifying all the other types it uses. The unspecified types are provided as **parameters** at the point of use.

An example of this is parametrizing the operation used for comparison by a sorting algorithm.
</details>

#### Relating Run-Time and Compile-Time Structures

<details>
<summary>Expand Content</summary>

The run-time structure of a program is often drastically different from the code structure. Code structure is frozen at compile-time; it consists of classes in fixed inheritance relationships. The run-time structure consists of rapidly changing networks of communicating objects. In fact, the two structures - run-time and compile-time - are largely independent.

With such disparity between the structures, it is clear that code won't reveal everything about how a system works. The run-time structure of the system must be imposed by the designer more than the language. The relationships between objects and their types must be designed with great care, because they will determine how good or bad the run-time structure is.
</details>

#### Designing for Change

<details>
<summary>Expand Content</summary>

The key to maximizing reuse lies in anticipating new requirements and changes to existing requirements, and designing your system so they can evolve appropriately.

To design systems robust to such changes, we must consider how the system may need to change over its lifetime. A design that does not take change into account risks major redesign in the future.

Some common causes of redesign:

1. Creating an object by specifying a class directly - commits us to a specific implementation, not interface.
1. Dependence on specific operations - commits us to a single way of satisfying a request.
1. Dependence on hardware and software platform - commits to a single software/hardware platform.
1. Dependence on object representations or implementations - clients that are aware of how an object is stored, represented, implemented may need changing when the object changes.
1. Algorithmic dependencies - Objects that depend on an algorithm will have to change when the algorithm changes.
1. Tight coupling - tightly coupled classes are hard to reuse in isolation. Leads to monolithic systems, where you can't change or remove a class without understanding many other classes.
1. Extending functionality by subclassing - requires understanding of parent classes and can lead to explosion of classes to make small changes.
1. Inability to alter classes conveniently - requires work arounds.
</details>

### How to Select a Design Pattern

There are many things to consider when selecting a design pattern.

1. Consider how the design pattern solve design problems.
1. Consider the intent of the design pattern.
1. Consider how design patterns interrelate. Maybe you need a group of patterns.
1. Consider patterns of like purpose.
1. Consider the causes of redesign to see if the design problem suffers from one or more.
1. Consider what should be variable in your design.

### How to Use a Design Pattern

Once a pattern has been selected, how do you use it? 

1. Read the pattern through for an overview.
    - Pay close attention to the applicability and consequences of the pattern.
1. Go back and study the structure, participants, and Collaborations of the pattern.
    - Ensure understanding of the classes, objects, and how they relate.
1. Look at the code section to see a concrete example of code in the pattern.
    - Studying the code helps to learn the pattern.
1. Choose names for pattern participants that are meaningful in the application context.
    - Often the names in the pattern are too abstract to appear directly in the application.
    - We should still include the participant name into the name that appears in the application for clarity.
1. Define the classes.
    - Declare the interfaces, establish inheritance relationships, and define the instance variables that represent data and object references.
    - Identify existing classes in the application that the new pattern will effect, and modify them accordingly.
1. Define application specific names for operations in the pattern.
    - Use the responsibilities and collaborations associated with each operation as a guide.
1. Implement the operations to carry out the responsibilities and collaborations in the pattern.
    - The code implementation and code sample sections provide guidance.

These are just guidelines to help get started.

### How Not to Use Design Patterns

It should also be mentioned how not to use them.

1. Do not use them indiscriminately.
    - Often the added flexibility and variability comes at the cost of indirection.
    - Pattern should only be used when the flexibility offered is actually needed.
