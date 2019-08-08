{{ template "layout/layout.tpl" . }}

{{ define "title" }} Guestbook - {{ end }}

{{ define "description" }} 
About Aroma House 
{{ end }}

{{ define "keywords" }}
Restaurant, Hotel, Homestay, Lodge, Guest House.
{{ end }}

{{ define "contents" }}

	<div class="container">
		<div class="md-body md-testimonail clearfix">
			<div class="grid_9 md-main">
				<div class="heading-absolute box-sidebar align-center">
					<header class="box-heading heading-large">
						T
						<div class="decription-override">
							<h2 class="h1">Testimonials</h2>
							<p>Etiam acondimentum</p>
						</div>
					</header>
				</div>

				{{ range $t := .Testimonials }}
				<div class="media">
					<a href="#" class="pull-left">
						<img class="media-object" src="http://placeholder.mac/80.png" alt="avatar">
					</a>
					<div class="media-body">
						<div class="box-quote box-quote-alter">
							<p>{{$t.Text}}</p>
						    <div class="text-link link-direct">{{$t.Title}}</div>
						</div>
					</div>
				</div><!-- /.media -->
				{{ end }}

			</div><!-- /.md-main -->
			<aside class="grid_3 md-sidebar md-sidebar-pt">
				<div class="box-sidebar">
					<h3 class="h3 header-sidebar">Check availability</h3>
					<div class="box-booking booking-stack">
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
							<div class="form-group row clearfix">
								<div class="col-left">
									<div class="form-group">
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
								</div>
								<div class="col-right">
									<div class="form-group">
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
								</div>
							</div>
							<div class="form-group">
								<label class="label-control"></label>
								<a href="#" class="btn btn-large btn-darkbrown">check</a>
							</div>
						</form>
					</div><!-- /.box-booking -->
				</div>
			</aside><!-- /.md-sidebar -->
		</div><!-- /.md-testimonail -->
	</div><!-- /.md-wrapper  -->

{{ end }}