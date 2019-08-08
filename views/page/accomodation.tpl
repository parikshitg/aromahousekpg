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
						<section class="md-booking">
							<h2 class="title-checkroom">Check availability</h2>
							<div class="box-booking booking-inline">
								<form>
									<div class="form-group">
										<label class="label-control">Arrival Date</label>
										<div class="booking-form select-black">
											<label class="collapse input">
												<input type="text" id="arrival-date" class="input-control border-black"/>
											</label>
										</div>
									</div>
									<div class="form-group">
										<label class="label-control">Departure Date</label>
										<div class="booking-form select-black">
											<label class="collapse input">
												<input type="text" id="departure-date" class="input-control border-black"/>
											</label>
										</div>
									</div>
									<div class="form-group select">
										<label class="label-control">Adults</label>
										<div class="input-group select-black">
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
									<div class="form-group select">
										<label class="label-control">Children</label>
										<div class="input-group select-black">
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
				</section>

				<section class="md-accomodation-content">
					<div class="row clearfix">
						{{ range $room := .Rooms}}
						<div class="grid_4">
							<article class="media">
								<figure>
									<img src="http://placeholder.mac/370x230.png" alt="" class="media-object">
								</figure>
								<section class="media-body">
									<h3 class="media-header h4"><a href="/room/{{$room.Slug}}">{{$room.Title}}</a></h3>
									<p class="media-content">{{ printf "%.300s" $room.Details }} ... </p>
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