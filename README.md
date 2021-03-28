# CS17B010_Distrubuted-Systems-Assignment_2
The main file is wiki.go

To run the code the commands are
 (go build wiki.go)
 (./wiki) 

And the server will be running on the port localhost:8080

To get all empreqs the route is (http://localhost:8080/empreqs)

To get any specific empoly details the route is (http://localhost:8080/empreqs/{Name})
here Name can be any employ name (Examples like :f1,f2,teja,ram)

To get any specific employ details with corresponding to respective date the route is (http://localhost:8080/empreqs/{Name}/{Date}) 
here Name can be any employ name (Examples like :f1,f2,teja,ram)
here Date can be of the form (Examples like: 22-03-2021,24-05-2021)

To get any specific employ details with corresponding to respective date and time the route is (http://localhost:8080/empreqs/{Name}/{Date}/{Time})
here Name can be any employ name (Examples like :f1,f2,teja,ram)
here Date can be of the form (Examples like: 22-03-2021,24-05-2021)
here Time can be of the form (Examples like: 02:00,04:00,05:00)

To get the meetings conducted in a month for the hod the route is (http://localhost:8080/f10)

To add a new request of any employ the route is (http://localhost:8080/edit)

To Update any request corresponding to any employee the route is (http://localhost:8080/empreqs/update/{Name}/{Date}/{Time})
here Name can be any employ name (Examples like :f1,f2,teja,ram)
here Date can be of the form (Examples like: 22-03-2021,24-05-2021)
here Time can be of the form (Examples like: 02:00,04:00,05:00)

To Delete any request corresponding to any employee the route is (http://localhost:8080/empreqs/delete/{Name}/{Date}/{Time})
here Name can be any employ name (Examples like :f1,f2,teja,ram)
here Date can be of the form (Examples like: 22-03-2021,24-05-2021)
here Time can be of the form (Examples like: 02:00,04:00,05:00)


And some entries are stored in data structure for testing purposes and when server is oned remaning enteries will be appended to the data structures

And concurrency is handled by mutexs by keeping in apporiate functions such that when their is a override the code of locks will serialize the actions send by 
multiple users

It is non-presistent storage so once server is down every data enteries stored is gone and only statical stored 5-6 are maintained in the data structue for
testing purpose

#And should use a chorme Borwser for viewing the details of editing,updating,creating the employee requests 
