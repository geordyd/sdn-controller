# SDN Controller

This project implements an SDN controller using a Pub/Sub architecture to efficiently handle events.

## Prerequisites

- Go
- `make` utility

## Getting Started

### Running the Program
To start the SDN controller, run:
```sh
make run
```

### Building the Program
To compile the program, use:
```sh
make build
```

### Testing the Program
Run tests with:
```sh
make test
```

## Managing Firewall Rules

### Add a Rule
To add a rule allowing or denying traffic on a specific port:
```sh
curl http://localhost:1337/addrule/{allow|deny}/{port}
```
Example:
```sh
curl http://localhost:1337/addrule/allow/80
```

### Remove a Rule
To remove a rule from the firewall:
```sh
curl http://localhost:1337/removerule/{port}
```
Example:
```sh
curl http://localhost:1337/removerule/80
```

## Retrieving Events
To get a list of all recorded events:
```sh
curl http://localhost:1337/getevents
```