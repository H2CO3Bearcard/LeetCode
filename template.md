
## Code FileName

```gotemplate
$!{question.frontendQuestionId}_${question.titleSlug}/$!{question.frontendQuestionId}_${question.title}_test
```


## Code Template
```gotemplate
${question.content}
package leetcode

import (
    "testing"
)

func Test$!velocityTool.camelCaseName(${question.titleSlug})(t *testing.T) {
	cases := []struct {
		Case   string
		Input  string
		Output string
	}{
	    {"1", "", ""},
	    {"2", "", ""},
	}

	for _, c := range cases {
		t.Run(c.Case, func(t *testing.T) {
			if output := $!velocityTool.smallCamelCaseName(${question.titleSlug})(c.Input); output != c.Output {
				t.Fatalf("测试用例%v不通过，期望：%v，得到：%v", c.Case, c.Output, output)
			}else{
			    t.Logf("测试用例%v通过了，期望：%v，得到：%v", c.Case, c.Output, output)
			}
		})
	}
}

${question.code}
```