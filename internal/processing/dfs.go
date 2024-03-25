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
	switch node.operator.Type.name {
	case "HEADING":
		{
			HTMLLine += prefixHeadings[node.operator.Text]
			for i := range node.operand {
				HTMLLine += LineLayout(*node.operand[i])
			}
			HTMLLine += postfixHeadings[node.operator.Text]
		}
	case "LINE":
		{
			HTMLLine += "<hr stule=\"border: none; background-color: black; color: black; height: 2px;\"></hr>"
		}
	case "SEMICOLON": // It will never happen
		{ // Reserved for future feature additions
			HTMLLine += "\n"
		}
	case "WORD":
		{
			HTMLLine += node.operator.Text
			for i := range node.operand {
				HTMLLine += LineLayout(*node.operand[i])
			}
		}
	case "NUMBEREDLIST":
		{
			HTMLLine += `<li style="list-style-type:'`
			HTMLLine += node.operator.Text
			HTMLLine += `'; margin-left:1vw">`
			for i := range node.operand {
				HTMLLine += LineLayout(*node.operand[i])
			}
			HTMLLine += "</li>"
		}
	case "CODE":
		{
			HTMLLine += "<code>"
			for i := range node.operand {
				HTMLLine += LineLayout(*node.operand[i])
			}
			HTMLLine += "</code>"
		}
	case "SPACE":
		{
			HTMLLine += " "
			for i := range node.operand {
				HTMLLine += LineLayout(*node.operand[i])
			}
		}
	case "LIST":
		{
			HTMLLine += `<li style="margin-left:1vw">`
			for i := range node.operand {
				HTMLLine += LineLayout(*node.operand[i])
			}
			HTMLLine += "</li>"
		}
	case "ITALIC":
		{
			HTMLLine += "<i>"
			HTMLLine += node.operator.Text[1 : len(node.operator.Text)-1]
			HTMLLine += "</i>"
			for i := range node.operand {
				HTMLLine += LineLayout(*node.operand[i])
			}
		}
	case "BOLT":
		{
			HTMLLine += "<b>"
			HTMLLine += node.operator.Text[2 : len(node.operator.Text)-2]
			HTMLLine += "</b>"
			for i := range node.operand {
				HTMLLine += LineLayout(*node.operand[i])
			}
		}
	case "SPECIALCHAR":
		{
			HTMLLine += node.operator.Text
			for i := range node.operand {
				HTMLLine += LineLayout(*node.operand[i])
			}
		}

	}
	return HTMLLine
}
