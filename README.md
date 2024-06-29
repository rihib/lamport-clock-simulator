# Lamport Clock Simulator

## What is Lamport Clock?

Lamport Clock is a logical clock used to determine the order of events in a distributed system. It was introduced by Leslie Lamport in 1978[1]. It is not a physical clock but a logical clock that helps to determine the order of events in a distributed system.

Lamport Clocks are used to establish a partial ordering of events in a distributed system. Each process in the system maintains a logical clock that is incremented for each event it generates. When a process sends a message, it includes its logical clock value in the message. When a process receives a message, it updates its logical clock value to be greater than the maximum of its current value and the value in the received message.

[1] Leslie Lamport. 1978. Time, clocks, and the ordering of events in a distributed system. Commun. ACM 21, 7 (July 1978), 558–565. https://doi.org/10.1145/359545.359563

## How to use this implementation?

This implementation simulates the operation of the Lamport clock with the CLI. There are two versions of the implementation, one in Go and the other in C. You can use either of them to simulate the Lamport clock.

### Setting up the environment

```bash
git clone https://github.com/rihib/lamport-clock-simulator.git
cd lamport-clock-simulator
go run go/main.go 9000
go run go/main.go 9001 // Open a new terminal
go run go/main.go 9002 // Open a new terminal
```

### Running the simulation

You can send and receive messages and perform computation events and see how the Lamport Clock value changes each time you do so.

```bash
// Terminal 1
% go run go/main.go 9000
Please type event, c or s: c
Calculate Event Success!!
Time: 1

Please type event, c or s: c
Calculate Event Success!!
Time: 2

Please type event, c or s: s
Please type the destination port: 9001
Sending Event Success!!
Time: 3

Please type event, c or s:
```

```bash
// Terminal 2
% go run go/main.go 9001
Please type event, c or s:
Message Received!!
Time: 4

Please type event, c or s:
```
