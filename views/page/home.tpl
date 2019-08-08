{{ template "layout/layout.tpl" . }}

{{ define "contents" }}

	{{ template "partial/slideshow.tpl" . }}
	
	<!-- Body -->
	<div class="md-body-wrapper">
		<div class="md-home">
			<section class="row check-availability">
				<div class="container clearfix">
					<div class="grid_3">
						<h2 class="title-checkroom">Check availability</h2>
					</div>
					<div class="grid_9">
						<section class="md-booking">
							<div class="box-booking booking-inline clearfix">
								<form>
									<div class="form-group">
										<label class="label-control">Arrival Date</label>
										<div class="booking-form select-white">
											<label class="collapse input">
												<input type="text" id="arrival-date" class="input-control border-white"/>
											</label>
										</div>
									</div>
									<div class="form-group">
										<label class="label-control">Departure Date</label>
										<div class="booking-form select-white">
											<label class="collapse input">
												<input type="text" id="departure-date" class="input-control border-white"/>
											</label>
										</div>
									</div>
									<div class="form-group select">
										<label class="label-control">Adults</label>
										<div class="input-group select-white">
											<label class="collapse">
												<select id="form-select" class="form-select">
													<option class="option-test">1</option>
													<option>2</option>
													<option>3</option>
													<option>4</option>
													<option>5</option>
												</select>
											</label>
										</div>
									</div>
									<div class="form-group select">
										<label class="label-control">Children</label>
										<div class="input-group select-white">
											<label class="collapse">
												<select class="form-select">
													<option>1</option>
													<option>2</option>
													<option>3</option>
													<option>4</option>
													<option>5</option>
												</select>
											</label>
										</div>
									</div>
									<div class="form-group last">
										<label class="label-control"></label>
										<a href="#" class="btn btn-large btn-darkbrown">Check</a>
									</div>
								</form>
							</div><!-- /.box-booking -->
						</section><!-- /.md-booking -->
					</div>
				</div><!-- /.container -->
			</section>
			<section class="row md-home-body">
				<div class="container clearfix">

					{{ with .Sidebar }}
					<aside class="grid_3 md-sidebar">
						<section class="widget-home-info">
							<h3 class="h3 header-sidebar">{{.Title}}</h3>
							<div class="home-content">
								{{htmlString .Richtext}}
								<ul class="list list-check">
									{{ range $facility :=  .List }}
									<li>{{ $facility }}</li>
									{{ end }}
								</ul>
							</div>
						</section>
					</aside><!-- /.md-sidebar -->
					{{ end }}

					<div class="grid_9 md-primary">

						{{ range $room := .Rooms }}
						<div class="row row-home">
							<article class="media">
								<figure class="media-image">
									<img src="http://placeholder.mac/459x280.png" alt="" class="media-object">
								</figure>
								<section class="media-body">
									<h3 class="media-header media-header-big">
										<a href="/room/{{.Slug}}">{{.Title}}</a>
									</h3>
									<p class="media-content">
										{{ printf "%.500s" $room.Details }} ...
										<a href="/room/{{.Slug}}" class="text-link link-direct">see more</a>
									</p>
								</section>
							</article>
						</div><!-- /.row-article -->
						{{ end }}

					</div><!-- /.md-sidebar -->

				</div>
			</section><!-- /.md-home-body -->
		</div>
	</div>

{{ end }}