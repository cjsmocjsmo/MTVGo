package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type headComp struct {
	app.Compo
	title string
}
func (h *headComp) Render() app.UI {
	return app.Head().Body(
		app.Title().Text(h.title),
		app.Link().Rel("stylesheet").Href("/web/css/app.css"),
	)
}

type navBarComp struct {
	app.Compo
}
func (h *navBarComp) Render() app.UI {
	return app.Nav().Class("navBar").Body(
		app.A().Href("/").Style("color", "White").Text("Movies"),
		app.A().Href("/tvshows").Style("color", "White").Text("TVShows"),
		app.A().Href("/search").Style("color", "White").Text("Search"),
	)
}

type action struct {
	app.Compo
}
func (h *action) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Action"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Action"),
		),
	)
}

type arnold struct {
	app.Compo
}
func (h *arnold) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Arnold"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Arnold"),
		),
	)
}

type brucelee struct {
	app.Compo
}
func (h *brucelee) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Bruce Lee"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Bruce Lee"),
		),
	)
}

type brucewillis struct {
	app.Compo
}
func (h *brucewillis) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Bruce Willis"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Bruce Willis"),
		),
	)
}

type buzz struct {
	app.Compo
}
func (h *buzz) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Buzz"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Buzz"),
		),
	)
}

type cartoons struct {
	app.Compo
}
func (h *cartoons) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Cartoons"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Cartoons"),
		),
	)
}

type charliebrown struct {
	app.Compo
}
func (h *charliebrown) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Charlie Brown"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Charlie Brown"),
		),
	)
}

type chucknorris struct {
	app.Compo
}
func (h *chucknorris) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Chuck Norris"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Chuck Norris"),
		),
	)
}

type comedy struct {
	app.Compo
}
func (h *comedy) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Comedy"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Comedy"),
		),
	)
}

type drama struct {
	app.Compo
}
func (h *drama) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Drama"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Drama"),
		),
	)
}

type documentary struct {
	app.Compo
}
func (h *documentary) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Documentary"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Documentary"),
		),
	)
}

type fantasy struct {
	app.Compo
}
func (h *fantasy) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Fantasy"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Fantasy"),
		),
	)
}

type ghostbusters struct {
	app.Compo
}
func (h *ghostbusters) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Ghost Busters"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Ghost Busters"),
		),
	)
}

type godzilla struct {
	app.Compo
}
func (h *godzilla) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Godzilla"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Godzilla"),
		),
	)
}

type harrisonford struct {
	app.Compo
}
func (h *harrisonford) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Harrison Ford"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Harrison Ford"),
		),
	)
}

type harrypotter struct {
	app.Compo
}
func (h *harrypotter) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Harry Potter"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Harry Potter"),
		),
	)
}

type hellboy struct {
	app.Compo
}
func (h *hellboy) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "HellBoy"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("HellBoy"),
		),
	)
}

type indianajones struct {
	app.Compo
}
func (h *indianajones) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Indiana Jones"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Indiana Jones"),
		),
	)
}

type jamesbond struct {
	app.Compo
}
func (h *jamesbond) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "James Bond"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("James Bond"),
		),
	)
}

type johnwayne struct {
	app.Compo
}
func (h *johnwayne) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "John Wayne"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("John Wayne"),
		),
	)
}

type johnwick struct {
	app.Compo
}
func (h *johnwick) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "John Wick"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("John Wick"),
		),
	)
}

type jurassicpark struct {
	app.Compo
}
func (h *jurassicpark) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Jurassic Park"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Jurassic Park"),
		),
	)
}

type kevincostner struct {
	app.Compo
}
func (h *kevincostner) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Kevin Costner"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Kevin Costner"),
		),
	)
}

type kingsmen struct {
	app.Compo
}
func (h *kingsmen) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Kingsmen"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Kingsmen"),
		),
	)
}

type lego struct {
	app.Compo
}
func (h *lego) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Lego"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Lego"),
		),
	)
}

type meninblack struct {
	app.Compo
}
func (h *meninblack) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Men In Black"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Men In Black"),
		),
	)
}

type minions struct {
	app.Compo
}
func (h *minions) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Minions"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Minions"),
		),
	)
}

type misc struct {
	app.Compo
}
func (h *misc) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Misc"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Misc"),
		),
	)
}

type nicolascage struct {
	app.Compo
}
func (h *nicolascage) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Nicolas Cage"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Nicolas Cage"),
		),
	)
}

type oldies struct {
	app.Compo
}
func (h *oldies) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Oldies"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Oldies"),
		),
	)
}

type pandas struct {
	app.Compo
}
func (h *pandas) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Pandas"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Pandas"),
		),
	)
}

type pirates struct {
	app.Compo
}
func (h *pirates) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Pirates"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Pirates"),
		),
	)
}

type riddick struct {
	app.Compo
}
func (h *riddick) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Riddick"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Riddick"),
		),
	)
}

type scifi struct {
	app.Compo
}
func (h *scifi) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "SciFi"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("SciFi"),
		),
	)
}

type stalone struct {
	app.Compo
}
func (h *stalone) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Stalone"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Stalone"),
		),
	)
}

type startrek struct {
	app.Compo
}
func (h *startrek) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Star Trek"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Star Trek"),
		),
	)
}

type starwars struct {
	app.Compo
}
func (h *starwars) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Star Wars"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Star Wars"),
		),
	)
}

type superheros struct {
	app.Compo
}
func (h *superheros) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Super Heros"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Super Heros"),
		),
	)
}

type tinkerbell struct {
	app.Compo
}
func (h *tinkerbell) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Tinker Bell"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Tinker Bell"),
		),
	)
}

type tomcruise struct {
	app.Compo
}
func (h *tomcruise) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Tom Cruise"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Tom Cruise"),
		),
	)
}

type transformers struct {
	app.Compo
}
func (h *transformers) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Transformers"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Transformers"),
		),
	)
}

type tremors struct {
	app.Compo
}
func (h *tremors) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Tremors"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Tremors"),
		),
	)
}

type therock struct {
	app.Compo
}
func (h *therock) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "The Rock"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("The Rock"),
		),
	)
}

type vandam struct {
	app.Compo
}
func (h *vandam) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Van Dam"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Van Dam"),
		),
	)
}

type xmen struct {
	app.Compo
}
func (h *xmen) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "XMen"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("XMen"),
		),
	)
}

type movies struct {
	app.Compo
}
func (h *movies) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Movies"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Movies"),
			app.Div().Class("wrapper").Body(
				app.Div().Class("wrapItem").Body(app.A().Href("/action").Text("Action")),
				app.Div().Class("wrapItem").Body(app.A().Href("/arnold").Text("Arnold")),
				app.Div().Class("wrapItem").Body(app.A().Href("/brucelee").Text("Bruce Lee")),
				app.Div().Class("wrapItem").Body(app.A().Href("/brucewillis").Text("Bruce Willis")),
				app.Div().Class("wrapItem").Body(app.A().Href("/buzz").Text("Buzz")),
				app.Div().Class("wrapItem").Body(app.A().Href("/cartoons").Text("Cartoons")),
				app.Div().Class("wrapItem").Body(app.A().Href("/charliebrown").Text("Charlie Brown")),
				app.Div().Class("wrapItem").Body(app.A().Href("/chucknorris").Text("Chuck Norris")),
				app.Div().Class("wrapItem").Body(app.A().Href("/comedy").Text("Comedy")),
				app.Div().Class("wrapItem").Body(app.A().Href("/drama").Text("Drama")),
				app.Div().Class("wrapItem").Body(app.A().Href("/documentary").Text("Documentary")),
				app.Div().Class("wrapItem").Body(app.A().Href("/fantasy").Text("Fantasy")),
				app.Div().Class("wrapItem").Body(app.A().Href("/ghostbusters").Text("Ghost Busters")),
				app.Div().Class("wrapItem").Body(app.A().Href("/godzilla").Text("Godzilla")),
				app.Div().Class("wrapItem").Body(app.A().Href("/harrisonford").Text("Harrison Ford")),
				app.Div().Class("wrapItem").Body(app.A().Href("/harrypotter").Text("Harry Potter")),
				app.Div().Class("wrapItem").Body(app.A().Href("/hellboy").Text("Hell Boy")),
				app.Div().Class("wrapItem").Body(app.A().Href("/indianajones").Text("Indiana Jones")),
				app.Div().Class("wrapItem").Body(app.A().Href("/jamesbond").Text("James Bond")),
				app.Div().Class("wrapItem").Body(app.A().Href("/johnwayne").Text("John Wayne")),
				app.Div().Class("wrapItem").Body(app.A().Href("/johnwick").Text("John Wick")),
				app.Div().Class("wrapItem").Body(app.A().Href("/jurassicpark").Text("Jurassic Park")),
				app.Div().Class("wrapItem").Body(app.A().Href("/kevincostner").Text("Kevin Costner")),
				app.Div().Class("wrapItem").Body(app.A().Href("/kingsmen").Text("Kings Men")),
				app.Div().Class("wrapItem").Body(app.A().Href("/lego").Text("Lego")),
				app.Div().Class("wrapItem").Body(app.A().Href("/meninblack").Text("Men In Black")),
				app.Div().Class("wrapItem").Body(app.A().Href("/minions").Text("Minions")),
				app.Div().Class("wrapItem").Body(app.A().Href("/misc").Text("Misc")),
				app.Div().Class("wrapItem").Body(app.A().Href("/nicolascage").Text("Nicolas Cage")),
				app.Div().Class("wrapItem").Body(app.A().Href("/oldies").Text("Oldies")),
				app.Div().Class("wrapItem").Body(app.A().Href("/pandas").Text("Pandas")),
				app.Div().Class("wrapItem").Body(app.A().Href("/pirates").Text("Pirates")),
				app.Div().Class("wrapItem").Body(app.A().Href("/riddick").Text("Riddick")),
				app.Div().Class("wrapItem").Body(app.A().Href("/scifi").Text("SciFi")),
				app.Div().Class("wrapItem").Body(app.A().Href("/stalone").Text("Stalone")),
				app.Div().Class("wrapItem").Body(app.A().Href("/startrek").Text("Star Trek")),
				app.Div().Class("wrapItem").Body(app.A().Href("/starwars").Text("Star Wars")),
				app.Div().Class("wrapItem").Body(app.A().Href("/superheros").Text("Super Heros")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tinkerbell").Text("Tinker Bell")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tomcruise").Text("Tom Cruise")),
				app.Div().Class("wrapItem").Body(app.A().Href("/transformers").Text("Trasnformers")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tremors").Text("Tremors")),
				app.Div().Class("wrapItem").Body(app.A().Href("/therock").Text("The Rock")),
				app.Div().Class("wrapItem").Body(app.A().Href("/vandam").Text("Van Dam")),
				app.Div().Class("wrapItem").Body(app.A().Href("/xmen").Text("X Men")),

			),
		),
	)
}

type tvaction struct {
	app.Compo
}
func (h *tvaction) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVAction"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVAction"),
		),
	)
}

type tvcomedy struct {
	app.Compo
}
func (h *tvcomedy) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVComedy"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVComedy"),
		),
	)
}

type tvfantasy struct {
	app.Compo
}
func (h *tvfantasy) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVFantasy"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVFantasy"),
		),
	)
}

type tvmcu struct {
	app.Compo
}
func (h *tvmcu) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVMCU"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVMCU"),
		),
	)
}

type tvscience struct {
	app.Compo
}
func (h *tvscience) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVScience"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVScience"),
		),
	)
}

type tvscifi struct {
	app.Compo
}
func (h *tvscifi) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVSciFi"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVSciFi"),
		),
	)
}

type tvstartrek struct {
	app.Compo
}
func (h *tvstartrek) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVStar Trek"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVStar Trek"),
		),
	)
}

type tvstarwars struct {
	app.Compo
}
func (h *tvstarwars) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVStar Wars"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVStar Wars"),
		),
	)
}

type tvwesterns struct {
	app.Compo
}
func (h *tvwesterns) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVWesterns"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVWesterns"),
		),
	)
}

type tvshows struct {
	app.Compo
}

func (h *tvshows) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "TVShows"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("TVShows"),
			app.Div().Class("wrapper").Body(
				app.Div().Class("wrapItem").Body(app.A().Href("/tvaction").Text("Action")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tvcomedy").Text("Comedy")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tvfantasy").Text("Fantasy")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tvmcu").Text("MCU")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tvscience").Text("Science")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tvscifi").Text("SciFi")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tvstartrek").Text("Star Trek")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tvstarwars").Text("Star Wars")),
				app.Div().Class("wrapItem").Body(app.A().Href("/tvwesterns").Text("Westerns")),
		),
	),
	)
}

type search struct {
	app.Compo
}

func (h *search) Render() app.UI {
	return app.Div().Body(
		&headComp{title: "Search"},
		&navBarComp{},
		app.Main().Body(
			app.H1().Class("test1").Text("Search"),
			app.H2().Class("test1").Text("Fuck"),
		),
	)
}

func main() {

	app.Route("/", func() app.Composer { return &movies{} })
	app.Route("/tvshows", func() app.Composer { return &tvshows{} })
	app.Route("/search", func() app.Composer { return &search{} })
	app.Route("/action", func() app.Composer { return &action{} })
	app.Route("/arnold", func() app.Composer { return &arnold{} })
	app.Route("/brucelee", func() app.Composer { return &brucelee{} })
	app.Route("/brucewillis", func() app.Composer{ return &brucewillis{} })
	app.Route("/buzz", func() app.Composer{ return &buzz{} })
	app.Route("/cartoons", func() app.Composer{ return &cartoons{} })
	app.Route("/charliebrown", func() app.Composer{ return &charliebrown{} })
	app.Route("/chucknorris", func() app.Composer{ return &chucknorris{} })
	app.Route("/comedy", func() app.Composer{ return &comedy{} })
	app.Route("/drama", func() app.Composer{ return &drama{} })
	app.Route("/documentary", func() app.Composer{ return &documentary{} })
	app.Route("/fantasy", func() app.Composer{ return &fantasy{} })
	app.Route("/ghostbusters", func() app.Composer{ return &ghostbusters{} })
	app.Route("/godzilla", func() app.Composer{ return &godzilla{} })
	app.Route("/harrisonford", func() app.Composer{ return &harrisonford{} })
	app.Route("/harrypotter", func() app.Composer{ return &harrypotter{} })
	app.Route("/hellboy", func() app.Composer{ return &hellboy{} })
	app.Route("/indianajones", func() app.Composer{ return &indianajones{} })
	app.Route("/jamesbond", func() app.Composer{ return &jamesbond{} })
	app.Route("/johnwayne", func() app.Composer{ return &johnwayne{} })
	app.Route("/johnwick", func() app.Composer{ return &johnwick{} })
	app.Route("/jurrassicpark", func() app.Composer{ return &jurassicpark{} })
	app.Route("/kevincostner", func() app.Composer{ return &kevincostner{} })
	app.Route("/kingsman", func() app.Composer{ return &kingsmen{} })
	app.Route("/lego", func() app.Composer{ return &lego{} })
	app.Route("/meninblack", func() app.Composer{ return &meninblack{} })
	app.Route("/minions", func() app.Composer{ return &minions{} })
	app.Route("/misc", func() app.Composer{ return &misc{} })
	app.Route("/nicolascage", func() app.Composer{ return &nicolascage{} })
	app.Route("/oldies", func() app.Composer{ return &oldies{} })
	app.Route("/pandas", func() app.Composer{ return &pandas{} })
	app.Route("/pirates", func() app.Composer{ return &pirates{} })
	app.Route("/riddick", func() app.Composer{ return &riddick{} })
	app.Route("/scifi", func() app.Composer{ return &scifi{} })
	app.Route("/stalone", func() app.Composer{ return &stalone{} })
	app.Route("/startrek", func() app.Composer{ return &startrek{} })
	app.Route("/starwars", func() app.Composer{ return &starwars{} })
	app.Route("/superheros", func() app.Composer{ return &superheros{} })
	app.Route("/tinkerbell", func() app.Composer{ return &tinkerbell{} })
	app.Route("/tomcruise", func() app.Composer{ return &tomcruise{} })
	app.Route("/transformers", func() app.Composer{ return &transformers{} })
	app.Route("/tremors", func() app.Composer{ return &tremors{} })
	app.Route("/therock", func() app.Composer{ return &therock{} })
	app.Route("/vandam", func() app.Composer{ return &vandam{} })
	app.Route("/xmen", func() app.Composer{ return &xmen{} })

	app.Route("/tvaction", func() app.Composer { return &tvaction{} })
	app.Route("/tvcomedy", func() app.Composer { return &tvcomedy{} })
	app.Route("/tvfantasy", func() app.Composer { return &tvfantasy{} })
	app.Route("/tvmcu", func() app.Composer { return &tvmcu{} })
	app.Route("/tvscience", func() app.Composer { return &tvscience{} })
	app.Route("/tvscifi", func() app.Composer { return &tvscifi{} })
	app.Route("/tvstartrek", func() app.Composer { return &tvstartrek{} })
	app.Route("/tvstarwars", func() app.Composer { return &tvstarwars{} })
	app.Route("/tvwesterns", func() app.Composer { return &tvwesterns{} })

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Movies",
		Description: "Movies",
	})

	http.Handle("/tvshows", &app.Handler{
		Name:        "Tvshows",
		Description: "Tvshows",
	})

	http.Handle("/search", &app.Handler{
		Name:        "Search",
		Description: "Search",
	})

	http.Handle("/action", &app.Handler{
		Name:	"Action",
		Description: "Action page",
	})

	http.Handle("/arnold", &app.Handler{
		Name:	"Arnold",
		Description: "Arnold page",
	})

	http.Handle("/brucelee", &app.Handler{
		Name:	"Bruce Lee",
		Description: "Bruce Lee page",
	})

	http.Handle("/brucewillis", &app.Handler{
		Name:	"Bruce Willis",
		Description: "Bruce Willis page",
	})

	http.Handle("/buzz", &app.Handler{
		Name:	"Buzz",
		Description: "Buzz page",
	})

	http.Handle("/cartoons", &app.Handler{
		Name:	"Cartoons",
		Description: "Cartoons page",
	})

	http.Handle("/charliebrown", &app.Handler{
		Name:	"Charlie Brown",
		Description: "Charlie Brown page",
	})

	http.Handle("/chucknorris", &app.Handler{
		Name:	"Chuck Norris",
		Description: "Chuck Norris page",
	})

	http.Handle("/comedy", &app.Handler{
		Name:	"Comedy",
		Description: "Comedy page",
	})

	http.Handle("/drama", &app.Handler{
		Name:	"Drama",
		Description: "Drama page",
	})

	http.Handle("/documentary", &app.Handler{
		Name:	"Documentary",
		Description: "Documentary page",
	})

	http.Handle("/fantasy", &app.Handler{
		Name:	"Fantasy",
		Description: "Fantasy page",
	})

	http.Handle("/ghostbusters", &app.Handler{
		Name:	"Ghost Busters",
		Description: "Ghost Busters page",
	})

	http.Handle("/godzilla", &app.Handler{
		Name:	"Godzilla",
		Description: "Godzilla page",
	})

	http.Handle("/harrisonford", &app.Handler{
		Name:	"Harrison Ford",
		Description: "Harrison Ford page",
	})

	http.Handle("/harrypotter", &app.Handler{
		Name:	"Harry Potter",
		Description: "Harry Potter page",
	})

	http.Handle("/hellboy", &app.Handler{
		Name:	"HellBoy",
		Description: "HellBoy page",
	})

	http.Handle("/indianajones", &app.Handler{
		Name:	"Indiana Jones",
		Description: "Indiana Jones page",
	})

	http.Handle("/jamesbond", &app.Handler{
		Name:	"James Bond",
		Description: "James Bond page",
	})

	http.Handle("/johnwayne", &app.Handler{
		Name:	"John Wayne",
		Description: "John Wayne page",
	})

	http.Handle("/johnwick", &app.Handler{
		Name:	"John Wick",
		Description: "John Wick page",
	})

	http.Handle("/jurassicpark", &app.Handler{
		Name:	"Jurassic Park",
		Description: "Jurassic Park page",
	})

	http.Handle("/kevincostner", &app.Handler{
		Name:	"Kevin Costner",
		Description: "Kevin Costner page",
	})

	http.Handle("/kingsmen", &app.Handler{
		Name:	"Kingsmen",
		Description: "Kingsmen page",
	})

	http.Handle("/lego", &app.Handler{
		Name:	"Lego",
		Description: "Lego page",
	})

	http.Handle("/meninblack", &app.Handler{
		Name: "Men In Black",
		Description: "Men In Black page",
	})

	http.Handle("/minions", &app.Handler{
		Name: "Minions",
		Description: "Minions page",
	})

	http.Handle("/misc", &app.Handler{
		Name: "Misc",
		Description: "Misc page",
	})

	http.Handle("/nicolascage", &app.Handler{
		Name: "Nicolas Cage",
		Description: "Nicolas Cage page",
	})

	http.Handle("/oldies", &app.Handler{
		Name: "Oldies",
		Description: "Oldies page",
	})

	http.Handle("/pandas", &app.Handler{
		Name: "Pandas",
		Description: "Pandas page",
	})

	http.Handle("/pirates", &app.Handler{
		Name: "Pirates",
		Description: "Pirates page",
	})

	http.Handle("/riddick", &app.Handler{
		Name: "Riddick",
		Description: "Riddick page",
	})

	http.Handle("/scifi", &app.Handler{
		Name: "SciFi",
		Description: "SciFi page",
	})

	http.Handle("/stalone", &app.Handler{
		Name: "Stalone",
		Description: "Stalone page",
	})

	http.Handle("/startrek", &app.Handler{
		Name: "Star Trek",
		Description: "Star Trek page",
	})

	http.Handle("/starwars", &app.Handler{
		Name: "Star Wars",
		Description: "Star Wars page",
	})

	http.Handle("/superheros", &app.Handler{
		Name: "Super Heros",
		Description: "Super Heros page",
	})

	http.Handle("/tinkerbell", &app.Handler{
		Name: "Tinker Bell",
		Description: "Tinker Bell page",
	})

	http.Handle("/tomcruise", &app.Handler{
		Name: "Tom Cruise",
		Description: "Tom Cruise page",
	})

	http.Handle("/transformers", &app.Handler{
		Name: "Transformers",
		Description: "Transformers page",
	})

	http.Handle("/tremors", &app.Handler{
		Name: "Tremors",
		Description: "Tremors page",
	})

	http.Handle("/therock", &app.Handler{
		Name: "The Rock",
		Description: "The Rock page",
	})

	http.Handle("/vandam", &app.Handler{
		Name: "Van Dam",
		Description: "Van Dam page",
	})

	http.Handle("/xmen", &app.Handler{
		Name: "XMen",
		Description: "XMen page",
	})

	http.Handle("/tvaction", &app.Handler{
		Name: "TVAction",
		Description: "TVAction page",
	})

	http.Handle("/tvcomedy", &app.Handler{
		Name: "TVComedy",
		Description: "TVComedy page",
	})

	http.Handle("/tvfantasy", &app.Handler{
		Name: "TVFantasy",
		Description: "TVFantasy page",
	})

	http.Handle("/tvmcu", &app.Handler{
		Name: "TVMCU",
		Description: "TVMCU page",
	})

	http.Handle("/tvscience", &app.Handler{
		Name: "TVScience",
		Description: "TVScience page",
	})

	http.Handle("/tvscifi", &app.Handler{
		Name: "TVSciFi",
		Description: "TVSciFi page",
	})

	http.Handle("/tvstartrek", &app.Handler{
		Name: "TVStar Trek",
		Description: "TVStar Trek page",
	})

	http.Handle("/tvstarwars", &app.Handler{
		Name: "TVStar Wars",
		Description: "TVStar Wars page",
	})

	http.Handle("/tvwesterns", &app.Handler{
		Name: "TVWesterns",
		Description: "TVWesterns page",
	})

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}
