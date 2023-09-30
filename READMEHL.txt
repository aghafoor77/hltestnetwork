Foe MS Services, we have to do following:
Copy connection file 
Change IP address in connection file
Add hosts in the sudo nano /etc/hosts file and point to the ledger like 

192.168.132.135 orderer.example.com
192.168.132.135 peer0.org1.example.com
192.168.132.135 peer0.org2.example.com

Change or enable HL connectivity in the user service's properties file


