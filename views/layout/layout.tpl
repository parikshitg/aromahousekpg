<!doctype html>
<html lang="en">
<head>
	<title>{{ block "title" . }} {{ end }} Aroma House </title>
	<meta charset="UTF-8">
	<meta name="format-detection" content="telephone=no">
	<!--  Viewport setting -->
	<meta name="viewport" content="width=device-width, initial-scale = 1.0, maximum-scale=1.0, user-scalable=no" />

	{{ range $k, $v  := .Meta }}
	<meta name="{{$k}}" content="{{$v}}">
	{{ end }}


  	<link rel="shortcut icon" type="image/png" href="/images/favicon.png">
	<!--  Font -->
<!--  Font -->
	<!-- link href='http://fonts.googleapis.com/css?family=Open+Sans:400,300,600' rel='stylesheet' type='text/css' -->
	<link rel="stylesheet" type="text/css" href="/css/font.css">
	<link rel="stylesheet" type="text/css" href="/css/font-awesome/css/font-awesome.min.css">
	
	<!-- Jquery UI CSS   -->
	<link rel="stylesheet" type="text/css" href="/css/libs/jquery-ui-1.10.3.custom.css">

	<!-- Library CSS  -->
	<link rel="stylesheet" type="text/css" href="/css/jquery.bxslider.css">

	<!-- Main CSS  -->
	<link rel="stylesheet" type="text/css" href="/css/style.css">


	<!-- Responsive CSS  -->
	<link rel="stylesheet" type="text/css" href="/css/responsive.css">
	<link rel="stylesheet" type="text/css" href="/css/responsive-menu.css">

	{{ block "css" . }} {{ end }}


	<!-- Fix ie9  -->
	<!--[if IE 9]>
	<link rel="stylesheet" type="text/css" href="css/ie9.css">
	<![endif]-->

</head>
<body>
	<!-- PRELOADER -->
	<div id="preloader">
		<span class="loader" data-loading-text="loading">
		</span>
	</div>
	<!-- END / PRELOADER -->
	<div class="md-hotel">
		<div id="mp-pusher" class="mp-pusher">
			{{ template "partial/header.tpl" . }}
			{{ template "partial/menubar.tpl" . }}
			{{ block "contents" . }} {{ end }}
			{{ template "partial/footer.tpl" . }}
		</div><!-- /.mp-pusher -->
	</div>

<!-- Library Javascript  -->
<script type="text/javascript" src="/js/modernizr.custom.js"></script>
<script type="text/javascript" src="/js/jquery-1.9.1.js"></script>
<script type="text/javascript" src="/js/jquery-ui-1.10.3.js"></script>
<script type="text/javascript" src="/js/jquery.bxslider.min.js"></script>


<!-- Responsive Menu Javascript  -->
<script type="text/javascript" src="/js/classie.js"></script>
<script type="text/javascript" src="/js/mlpushmenu.js"></script>

<!-- Main Javascript  -->
<script type="text/javascript" src="/js/script.js"></script>

{{ block "js" . }} {{ end }}

<!-- Separate Javascript for each page -->
<script type="text/javascript" src="/js/home.js"></script>
<script type="text/javascript">
	$(function() {
		"use strict";
		$("#arrival-date, #departure-date").datepicker({
		});
	});
</script>

<!-- Matomo -->
<script type="text/javascript">
  var _paq = window._paq || [];
  /* tracker methods like "setCustomDimension" should be called before "trackPageView" */
  _paq.push(['trackPageView']);
  _paq.push(['enableLinkTracking']);
  (function() {
    var u="//matomo.urantiatech.com/";
    _paq.push(['setTrackerUrl', u+'matomo.php']);
    _paq.push(['setSiteId', '4']);
    var d=document, g=d.createElement('script'), s=d.getElementsByTagName('script')[0];
    g.type='text/javascript'; g.async=true; g.defer=true; g.src=u+'matomo.js'; s.parentNode.insertBefore(g,s);
  })();
</script>
<!-- End Matomo Code -->


</body>
</html>