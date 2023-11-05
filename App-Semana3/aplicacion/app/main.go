package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Pelicula struct {
	ID            int    `json:"id"`
	Titulo        string `form:"titulo" binding:"required"`
	Genero        string `form:"genero" binding:"required"`
	Actores       string `form:"actores" binding:"required"`
	Duracion      string `form:"duracion" binding:"required"`
	Clasificacion string `form:"clasificacion" binding:"required"`
}

func initDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Crea la tabla Pelicula si no existe
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS peliculas (
            id SERIAL PRIMARY KEY,
            titulo VARCHAR(255) NOT NULL,
            genero VARCHAR(255) NOT NULL,
            actores VARCHAR(255) NOT NULL,
            duracion VARCHAR(255) NOT NULL,
            clasificacion VARCHAR(255) NOT NULL
        );
    `)

	return db
}

func main() {
	db = initDB()
	defer db.Close()

	r := gin.Default()
	port := os.Getenv("PORT")

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		var peliculas []Pelicula
		rows, err := db.Query("SELECT * FROM peliculas ORDER BY id ASC")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var p Pelicula
			err := rows.Scan(&p.ID, &p.Titulo, &p.Genero, &p.Actores, &p.Duracion, &p.Clasificacion)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			peliculas = append(peliculas, p)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"peliculas": peliculas,
		})
	})

	r.POST("/guardar", func(c *gin.Context) {
		var pelicula Pelicula
		if err := c.ShouldBind(&pelicula); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("INSERT INTO peliculas (titulo, genero, actores, duracion, clasificacion) VALUES ($1, $2, $3, $4, $5)",
			pelicula.Titulo, pelicula.Genero, pelicula.Actores, pelicula.Duracion, pelicula.Clasificacion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Redirigir a la página principal después de guardar el pelicula
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/actualizar", func(c *gin.Context) {
		id := c.PostForm("id")
		var pelicula Pelicula

		if err := c.ShouldBind(&pelicula); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Datos de la pelicula obtenidos:", pelicula)

		_, err := db.Exec("UPDATE peliculas SET titulo=$1, genero=$2, actores=$3, duracion=$4, clasificacion=$5 WHERE id=$6",
			pelicula.Titulo, pelicula.Genero, pelicula.Actores, pelicula.Duracion, pelicula.Clasificacion, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Redirect(http.StatusFound, "/")
	})

	r.GET("/eliminar.html", func(c *gin.Context) {
		id := c.DefaultQuery("id", "0")
		_, err := db.Exec("DELETE FROM peliculas WHERE id=$1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Redirect(http.StatusFound, "/")
	})

	r.Run(fmt.Sprintf(":%s", port))
}
