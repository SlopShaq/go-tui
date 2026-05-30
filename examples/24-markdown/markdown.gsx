package main

import (
	tui "github.com/grindlemire/go-tui"
)

type viewer struct {
	doc     string
	scrollY *tui.State[int]
	content *tui.Ref
}

func Viewer() *viewer {
	return &viewer{
		doc:     sampleDoc,
		scrollY: tui.NewState(0),
		content: tui.NewRef(),
	}
}

func (v *viewer) scrollBy(delta int) {
	el := v.content.El()
	if el == nil {
		return
	}
	_, maxY := el.MaxScroll()
	newY := v.scrollY.Get() + delta
	if newY < 0 {
		newY = 0
	}
	if newY > maxY {
		newY = maxY
	}
	v.scrollY.Set(newY)
}

func (v *viewer) KeyMap() tui.KeyMap {
	return tui.KeyMap{
		tui.On(tui.KeyEscape, func(ke tui.KeyEvent) { ke.App().Stop() }),
		tui.On(tui.Rune('q'), func(ke tui.KeyEvent) { ke.App().Stop() }),
		tui.On(tui.Rune('j'), func(ke tui.KeyEvent) { v.scrollBy(1) }),
		tui.On(tui.Rune('k'), func(ke tui.KeyEvent) { v.scrollBy(-1) }),
		tui.On(tui.KeyDown, func(ke tui.KeyEvent) { v.scrollBy(1) }),
		tui.On(tui.KeyUp, func(ke tui.KeyEvent) { v.scrollBy(-1) }),
		tui.On(tui.KeyPageDown, func(ke tui.KeyEvent) { v.scrollBy(10) }),
		tui.On(tui.KeyPageUp, func(ke tui.KeyEvent) { v.scrollBy(-10) }),
	}
}

func (v *viewer) HandleMouse(me tui.MouseEvent) bool {
	switch me.Button {
	case tui.MouseWheelUp:
		v.scrollBy(-3)
		return true
	case tui.MouseWheelDown:
		v.scrollBy(3)
		return true
	}
	return false
}

templ (v *viewer) Render() {
	<div class="flex-col p-1 border-rounded border-cyan">
		<span class="text-gradient-cyan-magenta font-bold">Markdown Viewer</span>
		<div
			ref={v.content}
			class="overflow-y-scroll scrollbar-cyan scrollbar-thumb-bright-cyan"
			height={20}
			scrollOffset={0, v.scrollY.Get()}>
			<markdown source={v.doc} width={72} />
		</div>
		<span class="font-dim">j/k or arrows scroll | q/esc quit</span>
	</div>
}
