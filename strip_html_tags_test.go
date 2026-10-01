package str_test

import (
	"testing"

	"github.com/dracory/str"
)

func TestStripHTMLTagsBasicHTML(t *testing.T) {
	input := "<p>Hello <b>World</b></p>"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsScriptTags(t *testing.T) {
	input := "Hello <script type=\"text/javascript\">alert('xss')</script>World"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsStyleTags(t *testing.T) {
	input := "Hello <style media=\"screen\">body { color: red; }</style>World"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsScriptAndStyleCaseInsensitive(t *testing.T) {
	input := "Hello <SCRIPT>alert(1)</SCRIPT> World <Style>body{}</Style>"
	expected := "Hello alert(1) World body{}"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsHTMLComments(t *testing.T) {
	input := "Hello <!-- comment --> World"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsSelfClosingTags(t *testing.T) {
	input := "Hello<br/>World <img src=\"test.jpg\" alt=\"image\" />"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsHTMLEntities(t *testing.T) {
	input := "Foo &amp; Bar &lt;Baz&gt; &quot;Qux&quot; &#39;Quux&#39; &nbsp; Space"
	expected := "Foo & Bar <Baz> \"Qux\" 'Quux' \u00a0 Space"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsWhitespaceAndNewlines(t *testing.T) {
	input := "  \n\t <p>  Multiple   \n spaces  </p>  \t "
	expected := "Multiple spaces"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsPlainText(t *testing.T) {
	input := "Plain text without HTML"
	expected := "Plain text without HTML"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsEmptyString(t *testing.T) {
	input := ""
	expected := ""
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsOnlyTags(t *testing.T) {
	input := "<div><p><span></span></p></div>"
	expected := ""
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsNestedTags(t *testing.T) {
	input := "<div>Outer <span>Inner</span> Text</div>"
	expected := "Outer Inner Text"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsUnclosedTags(t *testing.T) {
	input := "Hello <a href='link'> World"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestStripHTMLTagsWithoutClosingAngleBracket(t *testing.T) {
	input := "Hello <a href='link' World"
	expected := "Hello <a href='link' World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
