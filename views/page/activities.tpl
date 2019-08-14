{{ template "layout/layout.tpl" . }}

{{ define "title" }} Activities - {{ end }}

{{ define "contents" }}

	<div class="md-body-wrapper">
		<div class="container">
			<div class="md-body md-blog clearfix">
				<div class="grid_9 md-main">
					<section class="heading-absolute box-sidebar align-center">
						<header class="box-heading heading-large">
							A
							<div class="decription-override">
								<h2 class="h1">Activities</h2>	
								<p> Adventures Activities</p>
							</div>
						</header>
					</section>

					<div class="list-article">
						{{ range $a := .Activities }}
						<article class="media">
							<figure class="pull-left">
								<img src="{{ image $a.Image.URI 349 194 }}" alt="" class="media-object">
							</figure>
							<div class="media-body">
								<header class="media-header">
									
									<h3 class="h4"><a href="blog-detail.html">{{$a.Title}}</a></h3>
								</header>
								<p class="media-content">{{ htmlString $a.Body }} </p>
								
							</div>
						</article>
						{{ end }}

					</div>

				</div><!-- /.md-main -->
				<aside class="grid_3 md-sidebar md-sidebar-pt">

					<section class="box-sidebar">
						<h3 class="h3 header-sidebar">Rooms </h3>
						<ul class="list list-triangle list-uppercase">
							{{ range $room := .Rooms }}
							<li><a href="/room/{{$room.Slug}}">{{$room.Title}}</a></li>
							{{ end }}
						</ul>
					</section>

					<section class="box-sidebar">
						<h3 class="h3 header-sidebar">Packages</h3>
						<ul class="list list-triangle list-uppercase">
							{{ range $package := .Packages }}
							<li><a href="/packages">{{$package.Title}}</a></li>
							{{ end }}
						</ul>
					</section>

					{{ template "partial/check-sidebar.tpl" . }}

				</aside><!-- /.md-sidebar -->
			</div><!-- /.md-testimonail -->
		</div>
	</div>

{{ end }}