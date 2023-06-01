package _2_example

func NewDocumentFormat(format string) DocumentFormat {
	switch format {
	case "md":
		return &Markdown{}

	default:
		return &HTML{}
	}
}

type DocumentFormat interface {
	Header(string) string
	Bold(string) string
}

type Markdown struct{}

func (m *Markdown) Header(text string) string {
	return "# " + text
}

func (m *Markdown) Bold(text string) string {
	return "**" + text + "**"
}

type HTML struct{}

func (h *HTML) Header(text string) string {
	return "<h1>" + text + "</h1>"
}

func (h *HTML) Bold(text string) string {
	return "<strong>" + text + "</strong>"
}
