package career

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	r := strings.NewReader(`companies:
  - name: A
    summary: 〇〇をやった
    projects:
      - period: 2021年05月〜2021年06月
        role: リーダー
        technology: Go, TypeScript
        activities:
          - 実装をした
          - テストをした
      - period: 2021年05月〜2021年06月
        role: リーダー
        technology: Go, TypeScript
        activities:
          - 実装をした
          - テストをした
  - name: B
    summary: xxをやった
    projects:
      - period: 2021年04月〜2021年05月
        role: メンバー
        technology: PHP, MySQL
        activities:
          - マネジメントをした
`)
	_, err := Parse(r)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
}

func TestGenerator(t *testing.T) {
	c := Career{
		Companies: []Company{
			{
				"company name",
				"summary",
				[]Project{
					{
						"period",
						"role",
						"technology",
						[]Activity{
							"activity1",
							"activity2",
						},
					},
				},
			},
		},
	}

	_, err := Generate(c)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}
}
