{{ template "layout/layout.tpl" . }}

{{ define "title" }} Accomodation - {{ end }}

{{ define "contents" }}
	<!-- Body -->
	<div class="md-body-wrapper">
		<section class="container">
			<div class="md-body md-accomodation">
				<section class="row md-accomodation-available clearfix">
					<div class="grid_3 align-center">
						<header class="box-heading heading-large">
							A
							<div class="decription-override">
								<h1 class="h1">Accomodation</h1>
								<p>rates & reservations</p>
							</div>
						</header>
					</div>
					<div class="grid_9">
						{{ template "partial/check.tpl" . }}
					</div>
				</section>

				<section class="md-accomodation-content">
					<div class="row clearfix">
						{{ range $room := .Rooms}}
						<div class="grid_4">
							<article class="media">
								<figure>
									<img src="{{ image $room.Image.URI 370 230 }}" alt="" class="media-object">
								</figure>
								<section class="media-body">
									<h3 class="media-header h4"><a href="/room/{{$room.Slug}}">{{$room.Title}}</a></h3>
									{{ $details := htmlString $room.Details }}
									<p class="media-content">{{ printf "%.500s" $details }} ... </p>
									<h3 class="h3">Price : <span class="number number-big">₹ {{ .EPCharges }}<span class="one-night">/ Night</span></span></h3>
									<br>
									<a class="btn btn-small btn-border btn-border-brown" href="/room/{{$room.Slug}}">See Detail</a>
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