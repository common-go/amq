# activemq
Apache ActiveMQ is an open-source message broker written in Java, which supports a variety of messaging protocols. It is designed to facilitate communication between different applications and systems, allowing them to exchange information asynchronously and reliably. ActiveMQ is known for its robustness, flexibility, and performance, making it suitable for a wide range of use cases, from simple message queuing to complex enterprise-level messaging solutions.

### Libraries for Active MQ
- GO: [activemq](https://github.com/core-go/activemq), to wrap and simplify [go-stomp](https://github.com/go-stomp/stomp). Example is at [go-active-mq-sample](https://github.com/project-samples/go-active-mq-sample)
- nodejs: [activemq](https://www.npmjs.com/package/activemq), to wrap and simplify [activemq](https://www.npmjs.com/package/amqplib). Example is at [activemq-sample](https://github.com/typescript-tutorial/activemq-sample)

#### A common flow to consume a message from a message queue
![A common flow to consume a message from a message queue](https://cdn-images-1.medium.com/max/800/1*Y4QUN6QnfmJgaKigcNHbQA.png)
- The libraries to implement this flow are:
  - [mq](https://github.com/core-go/mq) for GOLANG. Example is at [go-active-mq-sample](https://github.com/project-samples/go-active-mq-sample)
  - [mq-one](https://www.npmjs.com/package/mq-one) for nodejs. Example is at [activemq-sample](https://github.com/typescript-tutorial/activemq-sample)

### Key Features of ActiveMQ
#### Multi-Protocol Support
- Supports several messaging protocols including AMQP, STOMP, MQTT, OpenWire, and WebSockets.
#### Persistence
- Ensures message durability by supporting message persistence to databases, file systems, and other storage mechanisms.
#### High Availability and Clustering
- Supports clustering and network of brokers for load balancing and fault tolerance, ensuring high availability and reliability.
#### Flexible Deployment Options
- Can be deployed as a standalone server, embedded within applications, or in a cloud environment.
#### JMS Compliant
- Fully compliant with the Java Message Service (JMS) API, making it easy to integrate with Java applications.
#### Advanced Messaging Features
- Supports a wide range of messaging patterns including point-to-point (queues), publish-subscribe (topics), request-reply, and more.
#### Security
- Provides robust security features including SSL/TLS encryption, authentication, and authorization.
#### Management and Monitoring
- Includes a web-based administration console and supports JMX for management and monitoring.
#### Pluggable Architecture
- Allows customization and extension through plugins and integration with other systems.

### Use Cases of ActiveMQ
#### Microservices Architecture
- Enables asynchronous communication between microservices, improving decoupling and scalability.
  ![Microservice Architecture](https://cdn-images-1.medium.com/max/800/1*vKeePO_UC73i7tfymSmYNA.png)
#### Event-Driven Architectures
- Implements event-driven patterns, allowing applications to react to events and changes asynchronously.
  ![A typical micro service](https://cdn-images-1.medium.com/max/800/1*d9kyekAbQYBxH-C6w38XZQ.png)
#### Real-Time Data Processing
- Supports real-time data streaming and processing, integrating with big data and analytics platforms.
#### Enterprise Application Integration
- Facilitates communication and data exchange between disparate enterprise applications and systems.
#### IoT Systems
- Handles messaging between IoT devices and backend systems, supporting protocols like MQTT for IoT communication.

### Example Scenario: Order Processing System
- In an order processing system, ActiveMQ can be used to integrate various components such as order entry, inventory management, payment processing, and shipping services.
#### Order Placement
- The order entry system sends an order message to a queue.
#### Inventory Check
- An inventory service consumes the order message, updates stock levels, and sends a message to the payment queue.
#### Payment Processing
- A payment service processes the payment and sends a confirmation message to a shipping queue.
#### Shipping
- A shipping service schedules shipment based on the confirmation message and updates the order status.
#### Throughout this process, ActiveMQ ensures reliable and asynchronous communication between the different services, improving scalability and fault tolerance.

### How ActiveMQ Works
ActiveMQ operates using the following core concepts:

#### Broker
- The central server component that routes messages between producers and consumers. A broker manages one or more queues and topics.
#### Queues
- Used in point-to-point messaging. Messages sent to a queue are delivered to one consumer.
#### Topics
- Used in publish-subscribe messaging. Messages sent to a topic are delivered to all subscribed consumers.
#### Messages
- The data or payload that is exchanged between producers and consumers. Messages can include headers, properties, and the body.
#### Producers
- Components or applications that send messages to queues or topics.
#### Consumers
- Components or applications that receive and process messages from queues or topics.
#### Persistent and Non-Persistent Messaging
- Persistent messages are stored to ensure they are not lost in case of broker failure.
- Non-persistent messages are faster but can be lost if the broker crashes.

### Advantages of ActiveMQ
#### Flexibility and Extensibility:
- Supports multiple protocols and can be extended to meet specific requirements through plugins.
#### Open Source
- Being open-source, it offers cost advantages and the ability to customize and contribute to its development.
#### Integration
- Seamlessly integrates with various enterprise systems, including Java applications, web applications, and other messaging systems.
#### Scalability
- Supports clustering and distributed configurations, allowing it to scale horizontally to handle increasing load.
#### Community and Support
- Backed by a large community and extensive documentation, providing ample resources for learning and troubleshooting.
#### Comprehensive Feature Set
- Offers a wide range of features and capabilities, from basic queuing to advanced messaging patterns and high availability.

### Disadvantages of ActiveMQ
#### Complexity
- The extensive feature set can introduce complexity, especially for small-scale or simple use cases.
#### Resource Intensive
- May require significant resources (CPU, memory) to handle large volumes of messages and complex configurations.
#### Performance Overhead
- The flexibility and pluggability can introduce performance overhead compared to more lightweight brokers.

### Conclusion
Apache ActiveMQ is a versatile and powerful message broker that supports a wide range of messaging protocols and patterns. Its flexibility, robustness, and extensive feature set make it suitable for a variety of use cases, from simple message queuing to complex enterprise-level messaging solutions. While it introduces some complexity and resource overhead, its benefits in terms of reliability, scalability, and integration capabilities make it a valuable tool for modern software architectures.

## Installation
Please make sure to initialize a Go module before installing core-go/activemq:

```shell
go get -u github.com/core-go/activemq
```

Import:

```go
import "github.com/core-go/activemq"
```
