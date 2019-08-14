{{ template "layout/layout.tpl" . }}

{{ define "title" }} Testimonials - {{ end }}

{{ define "contents" }}

	<div class="container">
		<div class="md-body md-testimonail clearfix">
			<div class="grid_9 md-main">
				<div class="heading-absolute box-sidebar align-center">
					<header class="box-heading heading-large">
						T
						<div class="decription-override">
							<h2 class="h1">Testimonials</h2>
							<p>Guest Reviews</p>
						</div>
					</header>
				</div>

				{{ range $t := .Testimonials }}
				<div class="media">
					<div class="pull-left">
						<img class="media-object" src="{{ image $t.Image.URI 80 80 }}" alt="avatar">
					</div>
					<div class="media-body">
						<div class="box-quote box-quote-alter">
							<p>{{$t.Text}}</p>
						    <div class="text-link link-direct">&mdash; {{$t.Title}}</div>
						</div>
					</div>
				</div><!-- /.media -->
				{{ end }}

			</div><!-- /.md-main -->
			<aside class="grid_3 md-sidebar md-sidebar-pt">
				{{ template "partial/check-sidebar.tpl" . }}
			</aside><!-- /.md-sidebar -->
		</div><!-- /.md-testimonail -->
	</div><!-- /.md-wrapper  -->

{{ end }}