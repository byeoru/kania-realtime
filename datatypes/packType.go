package datatypes

import "time"

type Pack struct {
	Info struct {
		Version     string    `json:"version"`
		Description string    `json:"description"`
		ExportedAt  time.Time `json:"exportedAt"`
		MapName     string    `json:"mapName"`
		Width       int       `json:"width"`
		Height      int       `json:"height"`
		Seed        string    `json:"seed"`
		MapID       int64     `json:"mapId"`
	} `json:"info"`
	Cells struct {
		Vertices struct {
			P [][]int `json:"p"`
			V [][]int `json:"v"`
			C [][]int `json:"c"`
		} `json:"vertices"`
		Cells struct {
			V [][]int          `json:"v"`
			C [][]int          `json:"c"`
			B []int            `json:"b"`
			I map[string]int32 `json:"i"`
			P [][]float64      `json:"p"`
			G map[string]int32 `json:"g"`
			Q struct {
				X0   int           `json:"_x0"`
				Y0   int           `json:"_y0"`
				X1   int           `json:"_x1"`
				Y1   int           `json:"_y1"`
				Root []interface{} `json:"_root"`
			} `json:"q"`
			H        map[string]int32          `json:"h"`
			Area     map[string]int32          `json:"area"`
			F        map[string]int32          `json:"f"`
			T        map[string]int32          `json:"t"`
			Haven    map[string]int32          `json:"haven"`
			Harbor   map[string]int32          `json:"harbor"`
			Biome    map[string]int32          `json:"biome"`
			Burg     map[string]int32          `json:"burg"`
			Conf     map[string]int32          `json:"conf"`
			Culture  map[string]int32          `json:"culture"`
			Fl       map[string]int32          `json:"fl"`
			Pop      map[string]int32          `json:"pop"`
			R        map[string]int32          `json:"r"`
			S        map[string]int32          `json:"s"`
			Religion map[string]int32          `json:"religion"`
			Province map[string]int32          `json:"province"`
			Routes   map[string]map[string]int `json:"routes"`
		} `json:"cells"`
		Features []interface{} `json:"features"`
		Rivers   []struct {
			I           int     `json:"i"`
			Source      int     `json:"source"`
			Mouth       int     `json:"mouth"`
			Discharge   int     `json:"discharge"`
			Length      float64 `json:"length"`
			Width       float64 `json:"width"`
			WidthFactor float64 `json:"widthFactor"`
			SourceWidth int     `json:"sourceWidth"`
			Parent      int     `json:"parent"`
			Cells       []int   `json:"cells"`
			Basin       int     `json:"basin"`
			Name        string  `json:"name"`
			Type        string  `json:"type"`
		} `json:"rivers"`
		Cultures []struct {
			Name         string        `json:"name"`
			I            int           `json:"i"`
			Base         int           `json:"base"`
			Origins      []interface{} `json:"origins"`
			Shield       string        `json:"shield"`
			Center       int           `json:"center,omitempty"`
			Color        string        `json:"color,omitempty"`
			Type         string        `json:"type,omitempty"`
			Expansionism float64       `json:"expansionism,omitempty"`
			Code         string        `json:"code,omitempty"`
		} `json:"cultures"`
		Burgs []struct {
			Cell       int     `json:"cell,omitempty"`
			X          float64 `json:"x,omitempty"`
			Y          float64 `json:"y,omitempty"`
			State      int     `json:"state,omitempty"`
			I          int     `json:"i,omitempty"`
			Culture    int     `json:"culture,omitempty"`
			Name       string  `json:"name,omitempty"`
			Feature    int     `json:"feature,omitempty"`
			Capital    int     `json:"capital,omitempty"`
			Port       int     `json:"port,omitempty"`
			Population float64 `json:"population,omitempty"`
			Type       string  `json:"type,omitempty"`
			Coa        struct {
				T1         string `json:"t1"`
				Ordinaries []struct {
					Ordinary string `json:"ordinary"`
					T        string `json:"t"`
					Line     string `json:"line"`
				} `json:"ordinaries"`
				Charges []struct {
					Charge string  `json:"charge"`
					T      string  `json:"t"`
					P      string  `json:"p"`
					Size   float64 `json:"size"`
				} `json:"charges"`
				Shield string `json:"shield"`
			} `json:"coa,omitempty"`
			Citadel int `json:"citadel,omitempty"`
			Plaza   int `json:"plaza,omitempty"`
			Walls   int `json:"walls,omitempty"`
			Shanty  int `json:"shanty,omitempty"`
			Temple  int `json:"temple,omitempty"`
		} `json:"burgs"`
		Routes []struct {
			I       int         `json:"i"`
			Group   string      `json:"group"`
			Feature int         `json:"feature"`
			Points  [][]float64 `json:"points"`
		} `json:"routes"`
		Religions []struct {
			Name         string      `json:"name"`
			I            int         `json:"i"`
			Origins      interface{} `json:"origins"`
			Type         string      `json:"type,omitempty"`
			Form         string      `json:"form,omitempty"`
			Culture      int         `json:"culture,omitempty"`
			Center       int         `json:"center,omitempty"`
			Deity        interface{} `json:"deity,omitempty"`
			Expansion    string      `json:"expansion,omitempty"`
			Expansionism int         `json:"expansionism,omitempty"`
			Color        string      `json:"color,omitempty"`
			Code         string      `json:"code,omitempty"`
		} `json:"religions"`
		Provinces []interface{} `json:"provinces"`
		Markers   []struct {
			Icon string  `json:"icon"`
			Type string  `json:"type"`
			Dx   int     `json:"dx,omitempty"`
			Px   int     `json:"px,omitempty"`
			X    float64 `json:"x"`
			Y    float64 `json:"y"`
			Cell int     `json:"cell"`
			I    int     `json:"i"`
			Dy   int     `json:"dy,omitempty"`
		} `json:"markers"`
	} `json:"cells"`
}
