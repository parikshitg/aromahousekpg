{{ template "layout/layout.tpl" . }}

{{ define "title" }} Packages - {{ end }}

{{ define "description" }} 
About Aroma House 
{{ end }}

{{ define "keywords" }}
Restaurant, Hotel, Homestay, Lodge, Guest House.
{{ end }}

{{ define "contents" }}
	<!-- Body -->
	<div class="md-body-wrapper">
		<section class="container">
			<div class="md-body md-accomodation">
				<section class="row md-accomodation-available clearfix">
					<div class="grid_3 align-center">
						<header class="box-heading heading-large">
							P
							<div class="decription-override">
								<h1 class="h1">Packages</h1>
								<p>All Inclusive Deals</p>
							</div>
						</header>
					</div>
					<div class="grid_9">
						{{ template "partial/check.tpl" . }}
					</div>
				</section>

				<section class="md-accomodation-content">
					<div class="row clearfix">
						{{ range $package := .Packages}}
						<div class="grid_4">
							<article class="media">
								<figure>
									<img src="{{ image $package.Image.URI 370 230 }}" alt="" class="media-object">
								</figure>
								<section class="media-body">
									<h3 class="media-header h3">{{$package.Title}}</h3>
									<p class="media-content"> {{htmlString $package.Details}} </p>
									<br>
									<div class="grid_12">
										<ul class="list list-check">
											{{ range $l := .List }}
											<li>{{ $l }}</li>
											{{ end }}
										</ul>
									</div>
									<a class="btn btn-small btn-border btn-border-brown" href="/reservation">Book Now</a>
								</section>
							</article>
						</div>
						{{ end }}
					</div>
				</section>
			</div><!-- /.md-accomodation -->
		</section><!-- /.md-wrapper  -->
	</div>
{{ end }}