package model

import (
	"bytes"
	"html/template"
	"log"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// Notice struct defines how a notice should be structured
type Notice struct {
	Name     string
	Regard   string
	Body     template.HTML
	Document *gofpdf.Fpdf
}

func (n *Notice) headerFunc() {
	n.Document.ImageOptions("./assets/images/logo.png", 145, 4, 51.054, 0, false, gofpdf.ImageOptions{}, 0, "")
	n.Document.SetFont("Kameron", "B", 15)
	n.Document.Cell(5, 0, "")
	n.Document.CellFormat(120.65, 12.7, "Massachusetts Health Connector Appeals Unit", "", 0, "L", false, 0, "")
	n.Document.Ln(13)
}

func (n *Notice) footerFunc() {
	n.Document.SetY(-27)
	n.Document.SetFont("OpenSans", "I", 10)
	n.Document.MultiCell(0, 5, "Health Connector Appeals Unit\nP.O. Box 960189, Boston, MA 02196\n617-933-3030 Fax 617-933-3099\nBusiness Hours Monday-Friday 8am-5pm", "0", "C", false)
}

func (n *Notice) appealHeaderFunc(a Appeal) {
	t := time.Now()
	n.Document.Ln(0.1)
	n.Document.Cellf(0, 10, "%v %v, %v", t.Month(), t.Day(), t.Year())
	n.Document.Ln(13)
	n.Document.Cellf(0, 10, "%v %v", a.Appellant.FirstName, a.Appellant.LastName)
	n.Document.Ln(4)
	n.Document.Cellf(0, 10, "%v", a.Appellant.Address.Street)
	n.Document.Ln(4)
	n.Document.Cellf(0, 10, "%v, %v %v", a.Appellant.Address.City, a.Appellant.Address.State, a.Appellant.Address.Zip)
	n.Document.Ln(10)
	n.Document.Line(15, 56, 190.5, 56)
	n.Document.Line(15, 65.5, 190.5, 65.5)
	n.Document.Cell(12, 10, "Re:")
	n.Document.Cell(15, 10, "Appeal - Acknowledgement of Appeal")
	n.Document.Ln(4)
	n.Document.Cell(12, 10, "")
	n.Document.Cellf(0, 10, "Appeal Number: %v", a.AppealID)
	n.Document.Ln(12)
}

func (n *Notice) bodyFunc(a Appeal) {
	n.Document.SetCellMargin(7)
	var templ *template.Template
	templ, err := templ.ParseFiles("./assets/templates/acknowledgementNoAidPending.html")
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Fatalln(err)
	}
	var tpl bytes.Buffer
	if err := templ.Execute(&tpl, a); err != nil {
		log.Fatalln(err)
	}
	htmlStr := tpl.String()
	html := n.Document.HTMLBasicNew()
	html.Write(6.5, htmlStr)
}

func setFonts(pdf *gofpdf.Fpdf) {
	pdf.AddFont("Franklin", "", "LibreFranklin-Regular.json")
	pdf.AddFont("Franklin", "B", "LibreFranklin-Bold.json")
	pdf.AddFont("Franklin", "I", "LibreFranklin-Italic.json")
	pdf.AddFont("Kameron", "", "Kameron-Regular.json")
	pdf.AddFont("Kameron", "B", "Kameron-Bold.json")
	pdf.AddFont("OpenSans", "", "OpenSans-Regular.json")
	pdf.AddFont("OpenSans", "B", "OpenSans-Bold.json")
	pdf.AddFont("OpenSans", "I", "OpenSans-Italic.json")
}

// NewNotice Generates a new notice.
func NewNotice(appeal Appeal) *Notice {
	notice := &Notice{}
	notice.Document = gofpdf.New("P", "mm", "A4", "./assets/fonts/")

	setFonts(notice.Document)

	notice.Document.SetHeaderFunc(notice.headerFunc)
	notice.Document.SetFooterFunc(notice.footerFunc)

	notice.Document.AddPage()
	notice.Document.AliasNbPages("")

	notice.Document.SetMargins(15, 15, 15)
	notice.Document.SetFont("Franklin", "", 10)
	notice.appealHeaderFunc(appeal)
	notice.bodyFunc(appeal)

	return notice
}
