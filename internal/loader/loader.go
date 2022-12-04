package loader

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"

	"github.com/pelletier/go-toml/v2"
	"icy.com/internal/models"
)

type loader struct {
	symbols      map[string]string
	data         map[string][]string
	output       chan models.Line
	eventHandler EventHandler
	Options      *Options
}

type EventHandler interface {
	NewLine(line models.Line)
}

type Options struct {
	directory string
}

type prefix byte

const (
	heading prefix = '#'
	dash prefix = '-'
)

func (p prefix) String() string{
switch p {
case '#':
	//log.Println("wha", string(p))
	return "heading"
case '-':
	return "dash"
}
	return "default"
}

//go:embed TEST.md
var main []byte

//go:embed styles.toml
var styles []byte

func New(eventHandler EventHandler) *loader {
	return &loader{data: make(map[string][]string), symbols: make(map[string]string), eventHandler: eventHandler}
}

func (ld *loader) Scan() {
	styles, err := ld.updateStyles(styles)
	if err != nil {
		log.Println(err.Error())
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(main))
	for i := 0; scanner.Scan(); i++ {
		log.Println(scanner.Text())
		line := ld.styleToLine(styles, scanner.Bytes())
		line.Text = scanner.Text() // exclude the first car
		line.Number = i
		ld.eventHandler.NewLine(line)
	}

}

func (ld *loader) updateStyles(data []byte) (map[string]models.Style, error){
	var styles = make(map[string]models.Style)
	err := toml.Unmarshal(data, &styles)
	log.Println(styles)
	return styles, err
}

func (ld *loader) styleToLine(styles map[string]models.Style, data []byte) models.Line {
	var pref prefix = ' '
	if len(data) > 0 {
	pref = prefix(data[0])
	}
	style, fd := styles[pref.String()]
	line := models.Line{}
	line.Style = &style
	if !fd {
		style = styles["default"]
		log.Println("couldnt find style for prefix falling down to def style", pref.String(), pref)
	}
	style.AsString = fmt.Sprintf("%dpx %s", style.Size, style.Font)
	line.Prefix = string(pref)
	return line
}
