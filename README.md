# MiniDHT
A simplified implementation of a Distributed Hash Table which uses concurrency mechanisms to mimic a P2P network allowing execution in absence of any external connection or communication.
<br><br>

<h2>Implementation</h2>
In a normal DHT, decentralized nodes distributed over a P2P network cooperate to store values collectively in the fashion of a classical hash table. This program uses Golang's Goroutines running in parallel to represent nodes which mimic the behaviour of P2P nodes. They will store values corresponding to their keyspace and retrieve values. They will communicate with each other to find values if a keyspace mismatch occurs.
<br>

<h2>Further Changes</h2>
Currently, DHT.go only allows for setting and retrieving of values. Further additions will include permitting a node to "fail" or "join" the "network" alongside more information gathering commands. 
