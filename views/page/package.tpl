{{ template "layout/layout.tpl" . }}

{{ define "title" }} Packages - {{ end }}

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
						{{ template "partial/check-package.tpl" . }}
					</div>
				</section>

				<section class="md-accomodation-content">
					<div class="row clearfix">
						{{ range $package := .Packages}}



						<!--div class="grid_4">
							<article class="media">
								<!--figure>
									<img src="{{ image $package.Image.URI 370 230 }}" alt="" class="media-object">
								</figure-->
								<section class="media-body">
									<h2 class="media-header h3">{{$package.Title}}</h2>
									<p class="media-content"> {{htmlString $package.Details}} </p>
									<br>
									<div class="grid_12">
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
								</section>
							</article>
						</div-->

						<br>
						<br>
						<br>

						{{ end }}
					</div>
				</section>
			</div><!-- /.md-accomodation -->
		</section><!-- /.md-wrapper  -->
	</div>








	
{{ end }}