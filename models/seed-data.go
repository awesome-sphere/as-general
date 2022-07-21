package models

import "log"

func SeedData() {
	log.Println("================ Start Seeding Data ================ ")
	DB.Create(&Movie{
		Title:       "Thor: Love and Thunder",
		Description: "nisi scelerisque eu ultrices vitae auctor eu augue ut lectus arcu bibendum at varius vel pharetra vel turpis nunc eget lorem dolor sed viverra ipsum nunc aliquet bibendum enim facilisis gravida neque convallis a cras semper auctor neque vitae tempus quam pellentesque nec nam aliquam sem et tortor consequat id",
		Duration:    120,
		Poster:      "thor.jpeg",
		Trailer:     "thor.mp4",
	})
	DB.Create(&Movie{
		Title:       "Elvis",
		Description: "nisi scelerisque eu ultrices vitae auctor eu augue ut lectus arcu bibendum at varius vel pharetra vel turpis nunc eget lorem dolor sed viverra ipsum nunc aliquet bibendum enim facilisis gravida neque convallis a cras semper auctor neque vitae tempus quam pellentesque nec nam aliquam sem et tortor consequat id",
		Duration:    145,
		Poster:      "elvis.jpeg",
		Trailer:     "elvis.mp4",
	})
	DB.Create(&Movie{
		Title:       "Jurassic World Dominion",
		Description: "nisi scelerisque eu ultrices vitae auctor eu augue ut lectus arcu bibendum at varius vel pharetra vel turpis nunc eget lorem dolor sed viverra ipsum nunc aliquet bibendum enim facilisis gravida neque convallis a cras semper auctor neque vitae tempus quam pellentesque nec nam aliquam sem et tortor consequat id",
		Duration:    138,
		Poster:      "jurassic.jpeg",
		Trailer:     "jurassic.mp4",
	})
	DB.Create(&Movie{
		Title:       "Lightyear",
		Description: "nisi scelerisque eu ultrices vitae auctor eu augue ut lectus arcu bibendum at varius vel pharetra vel turpis nunc eget lorem dolor sed viverra ipsum nunc aliquet bibendum enim facilisis gravida neque convallis a cras semper auctor neque vitae tempus quam pellentesque nec nam aliquam sem et tortor consequat id",
		Duration:    110,
		Poster:      "lightyear.jpeg",
		Trailer:     "lightyear.mp4",
	})
	DB.Create(&Movie{
		Title:       "Minions: The Rise of Gru",
		Description: "nisi scelerisque eu ultrices vitae auctor eu augue ut lectus arcu bibendum at varius vel pharetra vel turpis nunc eget lorem dolor sed viverra ipsum nunc aliquet bibendum enim facilisis gravida neque convallis a cras semper auctor neque vitae tempus quam pellentesque nec nam aliquam sem et tortor consequat id",
		Duration:    118,
		Poster:      "minions.jpeg",
		Trailer:     "minions.mp4",
	})
	log.Println("================ Done Seeding Data ================ ")

}
