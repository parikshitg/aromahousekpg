{{ template "layout/layout.tpl" . }}

{{ define "title" }} Packages - {{ end }}

{{ define "contents" }}

	<div class="md-body-wrapper">
		<div class="container">
			<div class="md-body md-blog clearfix">
				<div class="grid_9 md-main">
					<section class="heading-absolute box-sidebar align-center">
						<header class="box-heading heading-large">
							P
							<div class="decription-override">
								<h2 class="h1">Packages</h2>	
								<p> All inclusive packages</p>
							</div>
						</header>
					</section>

					<div class="">
						{{ range $package := .Packages }}
						<article class="">
							<div class="media-body">
									<h2 class="h2">{{$package.Title}}</h2>
									<br>
									<p class="media-content"> {{htmlString $package.Details}} </p>
									<br>
									<div>
										<ul class="list list-check">
											{{ range $l := .List }}
											<li>{{ $l }}</li>
											{{ end }}
										</ul>
									</div>
									<form method="post" action="/reservation">
										<input type="hidden" name="package" value="{{.Title}}">
										<input type="hidden" name="form" value="package">
										<button class="btn btn-small btn-border btn-border-brown">Book now</button>
									</form>
							</div>
						</article>

						<br>
						<br>
						<br>
						
						{{ end }}
					</div>

				</div>

				<aside class="grid_3 md-sidebar md-sidebar-pt">

					<section class="box-sidebar">
						<h3 class="h3 header-sidebar">Attractions </h3>
						<ul class="list list-triangle list-uppercase">
							{{ range $attraction := .Attractions }}
							<li><a href="/attractions">{{$attraction.Title}}</a></li>
							{{ end }}
						</ul>
					</section>

					<section class="box-sidebar">
						<h3 class="h3 header-sidebar">Activities </h3>
						<ul class="list list-triangle list-uppercase">
							{{ range $activity := .Activities }}
							<li><a href="/activities">{{$activity.Title}}</a></li>
							{{ end }}
						</ul>
					</section>

				</aside><!-- /.md-sidebar -->
			</div><!-- /.md-testimonail -->
		</div>
	</div>

{{ end }}