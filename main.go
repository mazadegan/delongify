package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/* (key = Slug, value = Url). Key-value pair struct. API should redirect any GET
request that it receives to the Url value pair of its Slug key. */
type SlugURLPair struct {
	// The url that
	Url       string    `bson:"Url"`
	Slug      string    `bson:"Slug"`
	CreatedAt time.Time `bson:"CreatedAt"`
	ExpireOn  int64     `bson:"ExpireOn"`
}

/* All SlugURLPairs in the database should expire this many seconds after their
creation. */
const SECONDS_TO_EXPIRATION = 60 * 5

func HelloHandler(c *gin.Context) {
	c.String(200, "Hello, world!")
}

/*	- Receives a url in the request body.
	- Produces a unique 6-character slug for SlugURLPair.
	- Inserts SlugURLPair into MongoDB Atlas database.
	- Returns SlugURLPair in message body.
*/
func ShortenUrlHandler(c *gin.Context) {
	var slugURLPair SlugURLPair
	// Grab url from request body
	if err := c.BindJSON(&slugURLPair); err != nil {
		// TODO: Handle this more gracefully.
		panic(err)
	}
	// Add 'http://' to beginning of url if not included
	if strings.Index(slugURLPair.Url, "http://") != 0 {
		slugURLPair.Url = "http://" + slugURLPair.Url
	}
	// Produce unique 6-character slug for slugURLPair.
	// Ensures uniqueness.
	for {
		slugURLPair.Slug = RandStr(6)
		if SlugIsUnique(slugURLPair.Slug) {
			break
		}
	}
	// Set time values
	slugURLPair.CreatedAt = time.Now().UTC()
	slugURLPair.ExpireOn = time.Now().Add(SECONDS_TO_EXPIRATION).Unix()

	// Insert slugURLPair into MongoDB Atlas database.
	InsertSlugURLPairToAtlasCollection(slugURLPair)

	c.JSON(200, slugURLPair)
}

func RedirectToSlugURLPairHandler(c *gin.Context) {
	slug := c.Param("slug")
	slugURLPair, err := GetSlugURLPair(slug)
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusMovedPermanently, slugURLPair.Url)
}

func main() {
	// load dotenv values
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// create default gin engine instance
	r := gin.Default()

	r.GET("/hello", HelloHandler)
	r.GET("/:slug", RedirectToSlugURLPairHandler)

	api := r.Group("/api")
	api.POST("/shorten", ShortenUrlHandler)

	// Serve built Vue app as application homepage.
	r.Use(static.Serve("/", static.LocalFile("./delongify_frontend/dist", true)))

	r.Run()
}
