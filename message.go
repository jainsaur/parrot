/*Package parrot - This is a simple chat service. Here are the features for version 1.0:

1. Client has to connect everytime to use the service.
2. A client can directly message any client who is online.

Protocol:
1. CONNECT <username>; return OK or ERR
2. MSG @<username>; broadcasts message to caller and callee
3. GETUSERS; returns a list of online users */
package parrot

// CONST - message types
const (
	CONNECT int = iota
	MSG
	GETUSERS
)
