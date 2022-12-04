package renderer

import (
	"log"
	"sync"
	"syscall/js"

	"icy.com/internal/models"
)

type renderer struct {
	doc    js.Value
	data   *lineData
	canvas *canvas
}

type lineData struct {
	mtx   sync.RWMutex
	lines map[int]models.Line
}

type canvas struct {
	canv   js.Value
	width  float64
	height float64
	ctx    js.Value
}

func New() *renderer {
	return &renderer{data: &lineData{lines: make(map[int]models.Line)}}
}

func (r *renderer) Init() {
	r.doc = js.Global().Get("document")
	canvas := canvas{}
	canvas.canv = r.doc.Call("getElementById", "canv")
	canvas.ctx = canvas.canv.Call("getContext", "2d")
	canvas.updateSize(&r.doc)
	r.canvas = &canvas
}

func (r *renderer) Start(drawLine chan models.Line, resize chan struct{}) {
	for {
		select {
		case v := <-drawLine:
			r.canvas.drawLine(v, r.data)

		case <-resize:
			r.canvas.updateSize(&r.doc)
			r.canvas.draw()
			log.Println(r.canvas.height, r.canvas.width)
		}
	}
}

func (c *canvas) updateSize(doc *js.Value) {
	if doc == nil {
		return
	}
	width := doc.Get("body").Get("clientWidth").Float()
	height := doc.Get("body").Get("clientHeight").Float()

	if width != c.width || height != c.height {
		c.canv.Set("width", width)
		c.canv.Set("height", 1080)
		c.width = width
		c.height = height
	}
}

func (c *canvas) drawLine(line models.Line, data *lineData) {
	data.mtx.Lock()
	defer data.mtx.Unlock()
	line.GetChords()
	c.ctx.Set("font", "30pt Arial")
	c.ctx.Call("fillText", line.Text, line.PosX, line.PosY)
	data.lines[line.Number] = line
	log.Println("drew", line)
}

func (c *canvas) draw() {
	c.ctx.Set("fillStyle", "black")
	c.ctx.Call("fillRect", 0, 0, c.width, c.height)
}
