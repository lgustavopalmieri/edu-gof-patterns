### Chain of Responsibility Pattern

The Chain of Responsibility pattern allows a request to be passed along a chain of handlers. Each handler decides either to process the request or pass it to the next handler in the chain. This approach promotes loose coupling and flexibility in processing requests.

#### Key Features:
- **Request Handling**: Each handler in the chain can either process the request or pass it to the next handler.
- **Decoupling**: The sender and receiver of a request are decoupled.
- **Dynamic Chains**: Chains of handlers can be modified dynamically.

Chain of Responsibility pattern allows for flexible request processing by decoupling the sender and the receiver, enabling modification of the processing chain without altering the client code.