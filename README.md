The Gang of Four (GoF) design patterns refer to 23 foundational design patterns introduced in the book "Design Patterns: Elements of Reusable Object-Oriented Software" by Erich Gamma, Richard Helm, Ralph Johnson, and John Vlissides. Here's an overview of these patterns, categorized based on their purposes:

### [Creational Patterns](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/creational)
These patterns focus on the process of object creation, aiming to make it more flexible and reusable.

1. [**Singleton**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/creational/Singleton): Ensures a class has only one instance and provides a global point of access to it.
2. [**Factory Method**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/creational/Factory_Method): Defines an interface for creating objects but allows subclasses to alter the type of objects that will be created.
3. [**Abstract Factory**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/creational/Abstract_Factory): Provides an interface for creating families of related or dependent objects without specifying their concrete classes.
4. [**Builder**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/creational/Builder): Separates the construction of a complex object from its representation, allowing the same construction process to create different representations.
5. [**Prototype**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/creational/Prototype): Creates new objects by copying an existing object, known as the prototype.

### [Structural Patterns](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural)
These patterns deal with object composition, organizing classes and objects to form larger structures.

1. [**Adapter**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Adapter): Allows incompatible interfaces to work together by acting as a bridge.
2. [**Decorator**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Decorator): Adds responsibilities to objects dynamically, providing a flexible alternative to subclassing for extending functionality.
3. [**Facade**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Facade): Provides a simplified interface to a complex subsystem.
4. [**Composite**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Composite): Composes objects into tree structures to represent part-whole hierarchies, allowing clients to treat individual objects and compositions uniformly.
5. [**Bridge**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Bridge): Separates an object’s abstraction from its implementation so that they can vary independently.
6. [**Flyweight**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Flyweight): Reduces the cost of creating and managing a large number of similar objects by sharing as much data as possible.
7. [**Proxy**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Proxy): Provides a surrogate or placeholder for another object to control access to it.

### [Behavioral Patterns](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/behavioral)
These patterns focus on effective communication and responsibility between objects, facilitating interaction and responsibility sharing.

1. [**Observer**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/behavioral/Observer): Defines a one-to-many dependency between objects, so that when one object changes state, all its dependents are notified and updated automatically.
2. [**Strategy**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/behavioral/Strategy): Defines a family of algorithms, encapsulates each one, and makes them interchangeable, allowing the algorithm to vary independently from the clients that use it.
3. [**Command**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/behavioral/Command): Encapsulates a request as an object, thereby allowing for parameterization of clients with different requests, queuing of requests, and logging of the requests.
4. [**Chain of Responsibility**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/behavioral/Chain_Of_Responsibility): Passes requests along a chain of handlers, allowing the request to be processed by one of the handlers without the sender knowing which one.
5. [**Mediator**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/behavioral/Mediator): Defines an object that encapsulates how a set of objects interact, promoting loose coupling by keeping objects from referring to each other explicitly.
6. [**Memento**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/behavioral/Memento): Captures and externalizes an object’s internal state so that it can be restored later, without violating encapsulation.
7. [**Visitor**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/behavioral/Visitor): Represents an operation to be performed on the elements of an object structure, allowing new operations to be defined without changing the classes of the elements on which it operates.
8. [**State**](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/behavioral/State): Allows an object to alter its behavior when its internal state changes, appearing to change its class.
9. **Template Method**: Defines the skeleton of an algorithm in a method, deferring some steps to subclasses, allowing them to redefine certain steps without changing the algorithm's structure.
10. **Iterator**: Provides a way to access elements of an aggregate object sequentially without exposing its underlying representation.
11. **Interpreter**: Defines a representation of a language’s grammar, along with an interpreter that uses the representation to interpret sentences in the language.

These patterns provide a valuable resource for software developers, offering proven solutions to common design problems in software development, and promoting code reuse and flexibility.