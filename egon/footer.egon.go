// Generated by egon.
// 🚫Edit at your own risk.

package egon
import (
"github.com/commondream/egon"
"fmt"
"io"
)

func FooterView() *egon.View {
	packageName := "egon"
	name := "Footer"
	templatePath := "footer.egon"
	renderFunc := func(w io.Writer) error {
		return FooterTemplate(w, )
	}
	return &egon.View{PackageName: packageName, Name: name, TemplatePath: templatePath, RenderFunc: renderFunc}
}

func FooterTemplate(w io.Writer) error {//line footer.egon:1
_, _ = fmt.Fprint(w, "<div class=\"footer\">copyright 2016</div>")
return nil
}
