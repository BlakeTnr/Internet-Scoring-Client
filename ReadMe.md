# Internet Scoring Client
This is a program that runs on the clients during blue-team red-team competitions to enforce realistic firewall policies

## How to use
To see if a team has legitimate LAN access to the internet, like normal workers would, you contact the client via its webserver on port :23021

## How to deploy
First clone this project. Go into the database directory and then replace the publickey.txt with your RSA public key. Then save the private key which will be used for decrypting the API responses. Then run "go build" to compile to an exe. Then you can run this program and then the API server will be hosted on port 23021. You can use the API Paths & parameters for communicating

## API Paths & parameters
- /internetconnection - returns an encrypted value, which when decoded shows true or false on whether the team has passed the test
    - ?attempts - attempts the test this many times - default=2

## How to read encrypted values
This program uses custom RSA encryption for communication. If your program uses go, you can take the encryption package from this code and use it, as it supports decryption. To decrypt, take the string, use base64 to decode the string to bytes then use those bytes to then decrypt using your RSA private key.
