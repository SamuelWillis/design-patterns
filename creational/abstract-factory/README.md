# Abstract Factory

## Intent

Provide an interface for creating families of related or dependent object without specifying their concrete classes.

## AKA

Kit

## Applicability

Use this pattern when:

1. A system should be independent of how its products are created, composed, and related.
1. A system should be configured with one of multiple families of products.
1. A family of related product objects is designed to be used together, and you need to enforce this constraint.
1. You want to provide a class library of products, and you want to reveal just their interfaces, not their implmentations.

## Structure

<!! TODO DIAGRAM !!>

## Participants

### AbstractFactory

Declares an interface for operations that create abstract product objects.

### ConcreteFactory

Implements the operations to create concrete product objects.

### AbstractProduct

Declares an interface for a type of product object

### ConcreteProduct

Defines a product object to be created by the corresponding concrete factory.

Implements the AbstractProduct interface.

### Client

Uses only the interfaces declared by AbstractFactory and AbstractProduct classes.

## Collaborations

Normally a single instance of a ConcreteFactory class is created at run-time.
This concrete factory creates the product objects having a particular implementation.
To create different product objects, clients should use a different concrete factory.

AbstractFactory defers the creation of product objects to its ConcreteFactory subclass.

## Consequences

The following are the benefits and liabilities of the AbstractFactory pattern:

### Isolation of Concrete Classes

This pattern helps you control the classes of objects that an application creates.
Because a factory encapsulates the responsibility and process of creating product objects, it isolates clients from implementation classes.

This means client manipulate instances through their abstract interfaces.

Product class names are isolated in the implementation of the concrete factory and they do not appear in client code.

### Exchanging Product Families is easy


The class of the ConcreteFactory appears only once in the application -- when it is instantiated.
This makes it easy to change the concrete factory an application uses.

Different product configurations can be used by simply changing the concrete factory.
Because the abstract factory creates a complete family of products, the whole product family is changed at once.

### Promotes Consistency Among Products

When product objects in a family are designed to work together, an application should only use objects from one family at a time.
AbstractFactory makes this trivial to enforce.

### Supporting New Kinds of Products is Difficult

Extending abstract factories to produce new kinds of products is not easy.
This is due to the AbstractFactory interface fixing which types of objects can be created.
To support new kinds of products the interface needs to be extended, this involves changing the AbstractFactory class and all of its subclasses.

## Implementation

Some helpful techniques for implementing this pattern:

### Factories as Singletons

An application typically only needs one instance of a ConcreteFactory per product family. So it's usually best implemented as a singleton.

### Creating the Products

The AbstractFactory only declares an interface for the products, it is up to the ConcreteFactory to actually create them.G

The most common way to do this is to define a factory method for each kind of product.
Then a ConcreteFactory will specify its products by overriding the factory method for each.

While this is simple, it requires a new concrete factory subclass for each product family.

If many product families are possible, the concrete family could be implemented using the Prototype pattern.
The factory is initialized with a prototypical instance of each product in the family, and it creates a new product by cloning the prototype.
This eliminates the need for a new concrete factory for each product family.

### Define Extensible Classes

The AbstractFactory usually defines a different operation for each kind of product it can produce.
These kinds are encoded into the operation signatures.
Adding a new product requires changing the AbstractFactory interface and all classes that depend on it.

A more flexible, but less safe, design is to add a parameter to operations that create objects.
In this approach, the AbstractFactory only needs a single "make" operation with a parameter indicating the kind of object to create.
This technique is used in the Prototype based approach to the AbstractFactory pattern.

There is an inherit problem with the above approach however.
All the products returned to the client with the same abstract interface, as given by the return type.
The client will not be able to differentiate or make safe assumptions about the class of a product.
This means subclass operations may not be available.

## Related Patterns

- Factory Method (AbstractFactory)
- Prototype (AbstractFactory)
- Singleton (ConcreteFactory)
