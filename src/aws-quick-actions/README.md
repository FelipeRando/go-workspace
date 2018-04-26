#Ideas for actions
#ALWAYS ASK THE USER IF EVERYTHING IS OK
#EX: (I will clone this machine (show everythin it needs to know) and then if the anwser is yes, do it!)
1- Clone machine:
    Clone a machine whith on the same VPC,Subnet, AMI, and everything. But you can specify that too
    (ex: clone machine i-0203kd3928j but on a different VPC)

2- reservation status:
    The state of all reservations, and throw me an alert in case of near expiration
    (ex: how much instances, when they are going to expirate)
    show the actual status: How much instances do I have and wich ones are using my reservations, is there 
    any reservation not being used? Is there any waste?

3- Manage testing accounts:
    I want to terminate all instances with age more than 3 months (in a POC account)