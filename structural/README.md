Structural design patterns are a category of the Gang of Four (GoF) design patterns that focus on how classes and objects can be composed or interact to form larger structures. They ease the design by identifying a simple way to realize relationships between entities. Here’s a more detailed look at each of the structural patterns:

### 1. [Adapter Pattern](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Adapter)
- **Purpose**: Allows objects with incompatible interfaces to work together.
- **Key Features**:
  - Acts as a bridge between two incompatible interfaces.
  - Can be implemented as a class or object adapter.
- **Common Uses**: Integrating with legacy systems, third-party libraries.

### 2. [Bridge Pattern](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Bridge)
- **Purpose**: Separates an object’s abstraction from its implementation, allowing them to vary independently.
- **Key Features**:
  - Decouples abstraction from implementation.
  - Generally used to avoid the explosion of classes typical in many class hierarchies.
- **Common Uses**: When both class and functionality changes are frequent, often seen in GUI toolkits.

### 3. [Composite Pattern](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Composite)
- **Purpose**: Composes objects into tree structures to represent part-whole hierarchies, allowing clients to treat individual objects and composites uniformly.
- **Key Features**:
  - Supports tree structures with leaves and composites.
  - Allows clients to work with composition of objects as if they were individual objects.
- **Common Uses**: Organization structures, graphic drawing applications.

### 4. [Decorator Pattern](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Decorator)
- **Purpose**: Adds responsibilities to objects dynamically, providing a flexible alternative to subclassing for extending functionality.
- **Key Features**:
  - Allows adding behavior to individual objects without affecting other objects from the same class.
  - Typically implemented by creating a set of decorator classes that are used to wrap concrete components.
- **Common Uses**: Adding features like scroll bars to windows, logging, or transaction management.

### 5. [Facade Pattern](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Facade)
- **Purpose**: Provides a simplified interface to a complex subsystem.
- **Key Features**:
  - Shields clients from subsystem components.
  - Can reduce the dependencies of outside code on the inner workings of a subsystem.
- **Common Uses**: Library isolation, enhancing readability and manageability of complex systems.

### 6. [Flyweight Pattern](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Flyweight)
- **Purpose**: Reduces memory usage by sharing as much data as possible with similar objects.
- **Key Features**:
  - Uses sharing to support a large number of fine-grained objects efficiently.
  - Includes intrinsic (shared) and extrinsic (unshared) state.
- **Common Uses**: Text editors, graphical applications for fonts or icon caching.

### 7. [Proxy Pattern](https://github.com/NikolaiKovalenko/edu-gof-patterns/tree/main/structural/Proxy)
- **Purpose**: Provides a surrogate or placeholder for another object to control access to it.
- **Key Features**:
  - Offers an identical interface to the real object.
  - Can be used for lazy initialization, access control, logging, etc.
- **Common Uses**: Remote proxies, virtual proxies, and protection proxies.

These structural patterns are essential for building extensible and maintainable software systems. They help manage the relationships and interactions between different parts of a system, allowing for flexible architectures that can easily adapt to change. Each pattern addresses specific problems and can be combined to solve more complex design challenges.