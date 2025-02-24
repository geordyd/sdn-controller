# SDN Controller

This project implements an SDN controller using a Pub/Sub architecture to efficiently handle events.

## DDD

### 1. Policy Management

**Responsibilities:**
- Managing allowed and blocked ports
- Handling rule modifications
- Ensuring firewall policies are enforced

**Bounded Context:** `Policy`  
**Entities:** `Policy`  
**Value Objects:** `Rule`  
**Aggregate:** `Policy`  
**Domain Events:** `RuleAdded`, `RuleRemoved`

### 2. Traffic Processing

**Responsibilities:**
- Inspecting traffic and checking it against policy rules
- Generating traffic events when traffic is processed
- Logging traffic decisions

**Bounded Context:** `Traffic`  
**Entities:** `Traffic`  
**Domain Events:** `TrafficReceived`, `TrafficAllowed`, `TrafficBlocked`, `TrafficDropped`

### Application Layer

The application layer contains the application services and handlers that coordinate the domain logic.

**Handlers:**
- Handles events related to policy changes and traffic processing

**Services:**
- Contains the logic to check if traffic is allowed based on the current policy

### Infrastructure Layer

The infrastructure layer contains the implementation of the Pub/Sub system used for event sourcing.


## Getting Started

### Prerequisites
- Go
- `make` utility

### Running the Program
To start the SDN controller, run:

```bash
make run
```

### Building the Program
To compile the program, use:

```bash
make build
```

### Testing the Program
Run tests with:

```bash
make test
```

### Managing Firewall Rules

#### Add a Rule
To add a rule allowing or denying traffic on a specific port:

```bash
curl http://localhost:1337/addrule/{allow|deny}/{port}
```

Example:

```bash
curl http://localhost:1337/addrule/allow/80
```

#### Remove a Rule
To remove a rule from the firewall:

```bash
curl http://localhost:1337/removerule/{port}
```

Example:

```bash
curl http://localhost:1337/removerule/80
```

### Retrieving Events
To get a list of all recorded events:

```bash
curl http://localhost:1337/getevents
```