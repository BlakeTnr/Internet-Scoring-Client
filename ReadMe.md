# Internet Scoring Client
This is a program that runs on the clients during blue-team red-team competitions to enforce realistic firewall policies

## How to use
To see if a team has legitimate LAN access to the internet, like normal workers would, you contact the client via its webserver on port :23021

## Paths & parameters
- /internetconnection - returns a true or false variable representing if the client passes the internet test
    - ?attempts - attempts the test this many times - default=2
    