package str_test

import (
	"testing"

	"github.com/dracory/str"
)

// TestStripHTMLTagsBasicHTML verifies that standard HTML tags like <p> and <b> are removed from the input string.
func TestStripHTMLTagsBasicHTML(t *testing.T) {
	input := "<p>Hello <b>World</b></p>"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsScriptTags verifies that <script> tags and their inner content are removed.
func TestStripHTMLTagsScriptTags(t *testing.T) {
	input := "Hello <script type=\"text/javascript\">alert('xss')</script>World"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsStyleTags verifies that <style> tags and their inner content are removed.
func TestStripHTMLTagsStyleTags(t *testing.T) {
	input := "Hello <style media=\"screen\">body { color: red; }</style>World"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsScriptAndStyleCaseInsensitive verifies tag matching behavior when script or style tags have mixed/uppercase casing.
func TestStripHTMLTagsScriptAndStyleCaseInsensitive(t *testing.T) {
	input := "Hello <SCRIPT>alert(1)</SCRIPT> World <Style>body{}</Style>"
	expected := "Hello alert(1) World body{}"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsHTMLComments verifies that HTML comments like <!-- ... --> are stripped out.
func TestStripHTMLTagsHTMLComments(t *testing.T) {
	input := "Hello <!-- comment --> World"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsSelfClosingTags verifies that self-closing tags like <br/> and <img> are properly removed.
func TestStripHTMLTagsSelfClosingTags(t *testing.T) {
	input := "Hello<br/>World <img src=\"test.jpg\" alt=\"image\" />"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsHTMLEntities verifies that HTML entities such as &amp;, &lt;, &gt;, &quot;, and &nbsp; are decoded.
func TestStripHTMLTagsHTMLEntities(t *testing.T) {
	input := "Foo &amp; Bar &lt;Baz&gt; &quot;Qux&quot; &#39;Quux&#39; &nbsp; Space"
	expected := "Foo & Bar <Baz> \"Qux\" 'Quux' \u00a0 Space"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsWhitespaceAndNewlines verifies that multiple spaces, tabs, and newlines are squished into single spaces and trimmed.
func TestStripHTMLTagsWhitespaceAndNewlines(t *testing.T) {
	input := "  \n\t <p>  Multiple   \n spaces  </p>  \t "
	expected := "Multiple spaces"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsPlainText verifies that input strings containing no HTML tags remain unchanged (other than trimming whitespace).
func TestStripHTMLTagsPlainText(t *testing.T) {
	input := "Plain text without HTML"
	expected := "Plain text without HTML"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsEmptyString verifies that passing an empty string yields an empty string output.
func TestStripHTMLTagsEmptyString(t *testing.T) {
	input := ""
	expected := ""
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsOnlyTags verifies that an input string consisting entirely of HTML tags results in an empty string.
func TestStripHTMLTagsOnlyTags(t *testing.T) {
	input := "<div><p><span></span></p></div>"
	expected := ""
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsNestedTags verifies that text inside nested HTML tags is preserved while all tags are removed.
func TestStripHTMLTagsNestedTags(t *testing.T) {
	input := "<div>Outer <span>Inner</span> Text</div>"
	expected := "Outer Inner Text"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsUnclosedTags verifies the behavior when an unclosed HTML tag with closing bracket is processed.
func TestStripHTMLTagsUnclosedTags(t *testing.T) {
	input := "Hello <a href='link'> World"
	expected := "Hello World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestStripHTMLTagsWithoutClosingAngleBracket verifies the behavior when text contains an open angle bracket with no closing bracket.
func TestStripHTMLTagsWithoutClosingAngleBracket(t *testing.T) {
	input := "Hello <a href='link' World"
	expected := "Hello <a href='link' World"
	result := str.StripHTMLTags(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
