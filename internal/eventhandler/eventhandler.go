package eventhandler

import "icy.com/internal/models"

type eventHandler struct {
	resize   chan struct{}
	drawLine chan models.Line
}

func New(resize chan struct{}, drawLine chan models.Line) *eventHandler {
	return &eventHandler{resize: resize, drawLine: drawLine}
}

func (h *eventHandler) Start() {
}

func (h *eventHandler) NewLine(line models.Line) {
	h.drawLine <- line
}
