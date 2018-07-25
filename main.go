package main

import ( "fmt"
	"encoding/json"
	elastic "gopkg.in/olivere/elastic.v3"
	"net/http"
	"log"
	"strconv"
      "reflect"
      "github.com/pborman/uuid"
)

const(
	DISTANCE = "200km"
	INDEX = "around"
	TYPE = "post"
	ES_URL = "http://35.199.31.117:9200"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	User string `json: "user"`
	Message string `json:"message"`
	Location Location `json: "location"`
}


func main() {
      	// Create a client
	client, err := elastic.NewClient(elastic.SetURL(ES_URL), elastic.SetSniff(false))
	if err != nil {
		panic(err)
		return
	}

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(INDEX).Do()
	if err != nil {
		panic(err)
	}
	if !exists {
		// Create a new index.
		mapping := `{
			"mappings":{
				"post":{
					"properties":{
						"location":{
							"type":"geo_point"
						}
					}
				}
			}
		}`
		_, err := client.CreateIndex(INDEX).Body(mapping).Do()
		if err != nil {
			// Handle error
			panic(err)
		}
	}

	fmt.Println("Started-service")
	http.HandleFunc("/search", handlerSearch)
	http.HandleFunc("/post", handlerPost)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one post request.")

	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}
      fmt.Fprintf(w, "Post received: %s\n", p.Message)
      
      id := uuid.New()
      saveToES(&p, id)
}

func saveToES(p *Post, id string) {
      es_client, err := elastic.NewClient(elastic.SetURL(ES_URL), 
            elastic.SetSniff(false))
      if err != nil {
            panic(err)
      }

	_, err = es_client.Index().
		Index(INDEX).
		Type(TYPE).
		Id(id).
		BodyJson(p).
		Refresh(true).
		Do()
	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("Post is saved to Index: %s\n", p.Message)

}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")

	lat,_ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon,_ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

	ran := DISTANCE
	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}

     fmt.Printf( "Search received: %f %f %s\n", lat, lon, ran)

      // Create a client
      client, err := elastic.NewClient(elastic.SetURL(ES_URL), elastic.SetSniff(false))
      if err != nil {
             panic(err)
             return
      }

      // Define geo distance query as specified in
      // https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-geo-distance-query.html
      q := elastic.NewGeoDistanceQuery("location")
      q = q.Distance(ran).Lat(lat).Lon(lon)

      // Some delay may range from seconds to minutes. So if you don't get enough results. Try it later.
      searchResult, err := client.Search().
             Index(INDEX).
             Query(q).
             Pretty(true).
             Do()
      if err != nil {
             // Handle error
             panic(err)
      }

      // searchResult is of type SearchResult and returns hits, suggestions,
      // and all kinds of other information from Elasticsearch.
      fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
      // TotalHits is another convenience function that works even when something goes wrong.
      fmt.Printf("Found a total of %d post\n", searchResult.TotalHits())

      // Each is a convenience function that iterates over hits in a search result.
      // It makes sure you don't need to check for nil values in the response.
      // However, it ignores errors in serialization.
      var typ Post
      var ps []Post
      for _, item := range searchResult.Each(reflect.TypeOf(typ)) { // instance of
             p := item.(Post) // p = (Post) item
             fmt.Printf("Post by %s: %s at lat %v and lon %v\n", p.User, p.Message, p.Location.Lat, p.Location.Lon)
             // TODO(student homework): Perform filtering based on keywords such as web spam etc.
             ps = append(ps, p)

      }
      js, err := json.Marshal(ps)
      if err != nil {
             panic(err)
             return
      }

      w.Header().Set("Content-Type", "application/json")
      w.Header().Set("Access-Control-Allow-Origin", "*")
      w.Write(js)
}



