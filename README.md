# CS17B010_Distrubuted-Systems-Assignment_2
The main file is wiki.go

To run the code the commands are
first (go build wiki.go)
second (./wiki) excutable

And the server will be running on the port localhost:8080

to get all empreqs the route is (http://localhost:8080/empreqs)
to get any specific empoly details the route is (http://localhost:8080/empreqs/{Name})
here Name can be any employ name (Examples like :f1,f2,teja,ram)
to get any specific employ details with corresponding to respective date the route is (http://localhost:8080/empreqs/{Name}/{Date}) 
here Name can be any employ name (Examples like :f1,f2,teja,ram)
here Date can be of the form (Examples like: 22-03-2021,24-05-2021)
to get any specific employ details with corresponding to respective date and time the route is (http://localhost:8080/empreqs/{Name}/{Date}/{Time})
here Name can be any employ name (Examples like :f1,f2,teja,ram)
here Date can be of the form (Examples like: 22-03-2021,24-05-2021)
here Time can be of the form (Examples like: 02:00,04:00,05:00)
To get the meetings conducted in a month for the hod the route is (http://localhost:8080/f10)
To add a new request of any employ the route is (http://localhost:8080/edit)
To Update any request corresponding to any employee the route is (http://localhost:8080/empreqs/update/{Name}/{Date}/{Time})
To Delete any request corresponding to any employee the route is (http://localhost:8080/empreqs/delete/{Name}/{Date}/{Time})
