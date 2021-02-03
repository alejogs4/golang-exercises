package chooseyouradventure

const HTMLStoryTemplate string = `
<!DOCTYPE html>
<head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
        {{range .Story}}
            <p>{{.}}</p>
        {{end}}
    <ul>
        {{range .Options}}
            <li> <a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>`

const NotFoundStoryTemplate string = `
<!DOCTYPE html>
<head>
    <meta charset="utf-8">
    <title>Not found story</title>
</head>
<body>
		<h1>Not found story, please go back</h1>
		<a href="/intro">Intro</a>
</body>`
