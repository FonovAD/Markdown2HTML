package MarkdownToHTML

const (
	HTMLsizeMultiplier = 10 // <- Both constants serve to reduce the number of passes
	HTMLsizeDevisor    = 3  // <- required to parse the string. Found experimentally
	HTMLPrefix         = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body style="margin-left: 3vw; margin-top: 2vh;">`
	HTMLPostfix = `
	</body>
	</html>`
)
