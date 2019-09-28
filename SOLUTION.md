# Major design choice
## Language
For REST service, almost all the typically language should work and have backing
framework to do it. I choose GO for 2 reasons. First it is suitable for lightweight
service and especially good for IO intensive application due to concurrency model of Go routine,
where this REST service is mostly IO based. Second, I am still in the learning mode
of GoLang, so this project would be a great hands on experience for me.
## REST framework
For the REST framework, because the purpose of the exercise is only having two sets
of APIs that operates on really lightweight Get and Create operation, and we are
not considering complex query as well as some middleware for authentication, so I am really looking for something lightweight to dispatch the routing. Gorilla/mux seems to be a great choice here, which is lightweight library that implements a request router and dispatcher for matching incoming requests to their respective handler, which fits the purpose of this project pretty well.
## ORM Mapping
For ORM mapping, we need to find some library to map the Model defined in the project to the SQL schema defined in the database. GORM is a well documented library that is easy to use and have all the features required for this project. In addition, it also have more features that can be used to extend the project in future requirement such as handler hooks and database relations.

## DataBase
Because the data provided to bootstrap is in the form of sql and the access pattern for the APIs mostly involves with searching based on attributes, so I think SQL would be a good fit. If we choose to do in NOSQL database, there would need to be some extra global index for the attributes need to be set, also we will have to do transactional read for the magic search, which makes the project more complex. For the simplicity of the project, I use the Sqlite3 for the SQL database implementation, however, using the GROM library in this project, we can easily switch to other SQL implementation such as PostgreSQL if necessary.  


# Example usage
## Install && bootstrap
`go install`  to install all the dependencies in the project, there is only one main package.
cd to the working directory, then run `go run *go` to start the server which is listened in the localhost at port 8080. For the simplicity, the project is hardcoded some of it's values, for production, should be stored in environment variables or some distributed config service.

## Customer APIs
* Note: All search attributes are case sensitive.
* Note: Current search implementation supports multiple attributes, but only one attributes at each query.
### Actor API
Actor API has three searchable attributes: "id", "imdb_id" and "name"
1. Search for actor with id equals 1

  `curl "http://localhost:8080/actsrch?id=1" --header "X-Identifier: C"`
2. Search for actor with imdb_id equals "nm0000842"

  `curl "http://localhost:8080/actsrch?imdb_id=nm0000842" --header "X-Identifier: C"`
3. Search for actor with name equals "John Savoca"

  `curl "http://localhost:8080/actsrch?name=John%20Savoca" --header "X-Identifier: C"`
### Movie API
Movie API has three searchable attributes: "title", "genres", "imdb_score", this has some semantic difference with actor API where all search are equality based, for title it is equal search, for genres because it is a string contain all genres split by comma, so we search if genre provided is one of the genres as contains search, for imdb_score it is a greater and equal search for all movies has higher score than input value.
1. Search for movie with title equals "Pulp Fiction"

  `curl "http://localhost:8080/movsrch?title=Pulp%20Fiction" --header "X-Identifier: C"`
2. Search for movie with one of the genres as Drama

  `curl "http://localhost:8080/movsrch?genres=drama" --header "X-Identifier: C"`
3. Search for movie has imdb_score greater or equal 8

  `curl "http://localhost:8080/movsrch?imdb_score=8" --header "X-Identifier: C"`
### Magic API
Magic search does not take any search attributes, it will always search
the movie that has Uma Truman as actor and Quentin Tarantino as director.

  `curl "http://localhost:8080/magicsrch?" --header "X-Identifier: C"`

## Admin APIs
### Create Actor
Create actor based on the value provides in the request, all fields besides id are require for the operation.

Example:
  `curl "http://localhost:8080/NEWACT"  -d '{"MovieId": 10, "Name": "here", "ImdbId": "ddd"}' --header "X-Identifier: ADMIN"`

### Create Movie
Create movie based on the value provides in the request, all fields besides id are require for the operation.

Example:
  `curl "http://localhost:8080/NEWMOV"  -d '{"ImdbId": "123", "Title": "new movie", "Director": "matt", "Year": 2000, "Rating": "R", "Genres": "Horror", "Runtime": 100, "Country": "US", "Language": "En", "ImdbScore": 9, "ImdbVotes": 10, "MetacriticScore": 10}' --header "X-Identifier: ADMIN"`


# Scaling
The bottle neck of the system in terms of scaling would be the SQL database. Because the server is using Gorilla/mux that is wrapper around GoLang net/http library, so for each request we are issuing the handler with a separate Go Routine, so we will scale pretty well in terms of IO throughput. Also it is easy to add more nodes of server and put a load balance in front of those nodes to further scaling out the system. However, it is because much harder to scaling out the sql database horizontally. For current use case, we do not have much relationship between two tables actors and movies, so it is still easier to scale our by partitioning the data. There are some managed solution such as AWS aurora provides promise of scalability and latency, however, if we are really end up with high throughput system, we need to migrate to a NOSQL database solution eventually.

# Source used for development
https://github.com/gorilla/mux
http://gorm.io/
https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b
