package utils_test
import (
	_"beego_study/test/initials"
	_"beego_study/initials"
	"testing"
	"fmt"
	"strings"
	"github.com/goquery"
	"html"
)

func TestHtml(t *testing.T) {
	s := "<p>dlfjafjadfjl;ajfd<pre><code>&lt;script src=&quot;http://code.jquery.com/jquery-1.9.1.js&quot;&gt;&lt;/script&gt;&lt;script src=&quot;../static/js/bootstrap.min.js&quot;&gt;&lt;/script&gt;&lt;script src=&quot;</p>"
	s = html.UnescapeString(s)
	reader := strings.NewReader(s)
	doc, _ := goquery.NewDocumentFromReader(reader)
	fmt.Println(doc.Text())
}
