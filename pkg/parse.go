package parse

import "fmt"

func GetHTTPURLFromOrigin(origin string) string {
	tokens := extractTokensFromOrigin(origin)
	return fmt.Sprintf("http://%s/%s/%s", tokens.Host, tokens.Org, tokens.Repo)
}
func GetHTTPSURLFromOrigin(origin string) string {
	tokens := extractTokensFromOrigin(origin)
	return fmt.Sprintf("https://%s/%s/%s", tokens.Host, tokens.Org, tokens.Repo)
}

func GetSSHURLFromOrigin(origin string) string {
	tokens := extractTokensFromOrigin(origin)
	return fmt.Sprintf("%s%s:%s/%s.git", tokens.GitUserName, tokens.Host, tokens.Org, tokens.Repo)
}

func TransformRawGitToClean(url string) string {
	scanning := true
	items := []string{}
	gitUsername := ""
	lexer := lex("url", url)
	for scanning {
		token := lexer.nextItem()
		switch token.typ {
		case itemGitUsername:
			gitUsername = token.val
		case itemHost, itemOrg, itemRepo:
			items = append(items, token.val)
		case itemError:
			return ""
		case itemEOF:
			scanning = false

		}

	}

	return gitUsername + fmt.Sprintf("%s:%s/%s.git", items[0], items[1], items[2])

}

type Tokens struct {
	Host        string
	Org         string
	Repo        string
	GitUserName string
}

func extractTokensFromOrigin(origin string) Tokens {
	scanning := true
	lexer := lex("url", origin)
	tokens := Tokens{}
	for scanning {
		token := lexer.nextItem()
		switch token.typ {
		case itemGitUsername:
			tokens.GitUserName = token.val
		case itemHost:
			tokens.Host = token.val
		case itemOrg:
			tokens.Org = token.val
		case itemRepo:
			tokens.Repo = token.val
		case itemError:
			return Tokens{}
		case itemEOF:
			scanning = false

		}

	}

	return tokens
}
