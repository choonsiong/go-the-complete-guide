# Notes

By the way, you can find the Docker image for section 11 here: 

https://hub.docker.com/r/choonsiong/go-rest-api/tags

To run the container:
```
$ docker pull choonsiong/go-rest-api
$
$ docker run -d -p 8080:8080 --name go-rest-api choonsiong/go-rest-api
$
$ curl -s http://localhost:8080/events | json_pp
{
   "events" : [
      {
         "datetime" : "2025-02-28T14:48:38.166189507Z",
         "description" : "Sed posuere consectetur est at lobortis. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Nullam id dolor id nibh ultricies vehicula ut id elit. Nulla vitae elit libero, a pharetra augue. Nullam quis risus eget urna mollis ornare vel eu leo. Donec sed odio dui.",
         "id" : 1,
         "location" : "Tokyo, Japan",
         "name" : "Lorem Fermentum Inceptos",
         "user_id" : 1
      }
   ],
   "message" : "success"
}
$
```