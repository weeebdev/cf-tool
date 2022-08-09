package client

import "testing"

func TestSimpleFindSample(t *testing.T) {
	body := "<div class=\"sample-tests\"><div class=\"section-title\">Example</div><div class=\"sample-test\"><div class=\"input\"><div class=\"title\">Input<div title=\"Copy\" data-clipboard-target=\"#id004679888550762963\" id=\"id0081034027387341\" class=\"input-output-copier\">Copy</div></div><pre id=\"id004679888550762963\">4\n3876\n387\n4489\n3\n</pre></div><div class=\"output\"><div class=\"title\">Output<div title=\"Copy\" data-clipboard-target=\"#id008440845462789317\" id=\"id002626928483810227\" class=\"input-output-copier\">Copy</div></div><pre id=\"id008440845462789317\">0\n2\n1\n-1\n</pre></div></div></div>"
	input, output, _ := findSample([]byte(body))
	expectInput := "4\n3876\n387\n4489\n3\n"
	expectOutput := "0\n2\n1\n-1\n"
	realInput := ""
	realOutput := ""
	for i := 0; i < len(input); i += 1 {
		realInput += string(input[i])
	}
	for i := 0; i < len(output); i += 1 {
		realOutput += string(output[i])
	}
	if realInput != expectInput {
		t.Errorf("Expect %s, but found %s.", expectInput, realInput)
	}
	if realOutput != expectOutput {
		t.Errorf("Expect %s, but found %s.", expectOutput, realOutput)
	}
}

func TestHighlightFindSample(t *testing.T) {
	body := "<div class=\"sample-tests\"><div class=\"section-title\">Example</div><div class=\"sample-test\"><div class=\"input\"><div class=\"title\">Input<div title=\"Copy\" data-clipboard-target=\"#id004159111468340382\" id=\"id0022207187427049613\" class=\"input-output-copier\">Copy</div></div><pre id=\"id004159111468340382\"><div class=\"test-example-line test-example-line-even test-example-line-0\">3</div><div class=\"test-example-line test-example-line-odd test-example-line-1\">4</div><div class=\"test-example-line test-example-line-odd test-example-line-1\">0 -2</div><div class=\"test-example-line test-example-line-odd test-example-line-1\">1 0</div><div class=\"test-example-line test-example-line-odd test-example-line-1\">-1 0</div><div class=\"test-example-line test-example-line-odd test-example-line-1\">0 2</div><div class=\"test-example-line test-example-line-even test-example-line-2\">3</div><div class=\"test-example-line test-example-line-even test-example-line-2\">0 2</div><div class=\"test-example-line test-example-line-even test-example-line-2\">-3 0</div><div class=\"test-example-line test-example-line-even test-example-line-2\">0 -1</div><div class=\"test-example-line test-example-line-odd test-example-line-3\">1</div><div class=\"test-example-line test-example-line-odd test-example-line-3\">0 0</div><div class=\"test-example-line test-example-line-odd test-example-line-3\"></div></pre></div><div class=\"output\"><div class=\"title\">Output<div title=\"Copy\" data-clipboard-target=\"#id007798350551237874\" id=\"id007534479199011845\" class=\"input-output-copier\">Copy</div></div><pre id=\"id007798350551237874\">12\n12\n0\n</pre></div></div></div>"
	input, output, _ := findSample([]byte(body))
	expectInput := "3\n4\n0 -2\n1 0\n-1 0\n0 2\n3\n0 2\n-3 0\n0 -1\n1\n0 0\n"
	expectOutput := "12\n12\n0\n"
	realInput := ""
	realOutput := ""
	for i := 0; i < len(input); i += 1 {
		realInput += string(input[i])
	}
	for i := 0; i < len(output); i += 1 {
		realOutput += string(output[i])
	}
	if realInput != expectInput {
		t.Errorf("Expect %s, but found %s.", expectInput, realInput)
	}
	if realOutput != expectOutput {
		t.Errorf("Expect %s, but found %s.", expectOutput, realOutput)
	}
}
