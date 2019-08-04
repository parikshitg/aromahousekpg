<!DOCTYPE html>
<html>
<head>
  <title>Aroma House Kalimpong</title>
</head>
<body>
<h1>Blogs</h1>
<ul>
  {{ range $b := .Blogs }}
  <li>{{ $b.Title }}</li>
  {{ end }}
</ul>
</body>
</html>