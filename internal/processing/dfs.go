package processing

func Run(node StatmentsNode) string {
	HTML := ""
	for i := range node.CodeString {
		HTML += LineLayout(node.CodeString[i])
	}
	return HTML
}

func LineLayout(node Node) string {
	HTMLLine := ""
	if node.operator.Type.name == "HEADING" {
		HTMLLine += prefixHeadings[node.operator.Text]
		for i := range node.operand {
			HTMLLine += LineLayout(*node.operand[i])
		}
		HTMLLine += postfixHeadings[node.operator.Text]
		return HTMLLine
	}
	if node.operator.Type.name == "LINE" {
		HTMLLine += "<hr stule=\"border: none; background-color: black; color: black; height: 2px;\"></hr>"
		return HTMLLine
	}
	if node.operator.Type.name == "SEMICOLON" {
		HTMLLine += "\n"
		return HTMLLine
	}
	if node.operator.Type.name == "WORD" {
		HTMLLine += node.operator.Text
		for i := range node.operand {
			HTMLLine += LineLayout(*node.operand[i])
		}
		return HTMLLine
	}
	if node.operator.Type.name == "GROUPNUMBEREDLIST" {
		HTMLLine += "<ol>"
		for i := range node.operand {
			HTMLLine += LineLayout(*node.operand[i])
		}
		HTMLLine += "</ol>"
		return HTMLLine
	}
	if node.operator.Type.name == "CODE" {
		HTMLLine += "<code>"
		for i := range node.operand {
			HTMLLine += LineLayout(*node.operand[i])
		}
		HTMLLine += "</code>"
		return HTMLLine
	}
	if node.operator.Type.name == "SPACE" {
		HTMLLine += " "
		for i := range node.operand {
			HTMLLine += LineLayout(*node.operand[i])
		}
		return HTMLLine
	}
	return HTMLLine
}
