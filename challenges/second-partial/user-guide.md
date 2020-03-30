User Guide ========== To start running the server you need to posisionate in the folder of the file. Once your there you run the program whith this command. go run Challenge2.go
  
Before start asking any request to the program you need to enter to your git bash terminal, and posisionate in the folder of the file. Once the program is up and running you need to 
log in, in your git bash terminal, you enter this command.
  curl -u username:password http://localhost:8080/login
  
  In username, you need to introduce one of the default users with his default password: user1 --> password: love user2 --> password: god user3 --> password: save Once your log, you 
  will recive a TOKEN you're gonna need this to access the other functions of the program. If you´re already loged and run this command again, the program wil send you a message, 
  and it will not give you your token.
  
To get the status you need to run this command. curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/status
  
  You will need to put your given token in <ACCESS_TOKEN>, after running the command you will recive information about the user.
  
To upload a file to need to run this command. curl -F "data=@path/to/local/image.png" -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/upload
  
  To upload a file you need to introduce the complete path of the file, you will need to change path/to/local/image.png to the rigth path, than you will need to introduce the 
  corresponding user token in <ACCESS_TOKEN>.
  
To logout you will need to run this command. curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/logout
  
  You will need to introduce your token, to logout, once you're loged out the token that was given to you will be eliminated, and to get   another, yout will had to log in again.
