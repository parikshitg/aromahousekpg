{{ template "layout/layout.tpl" . }}

{{ define "title" }} Rooms - {{ end }}

{{ define "js" }}
<script type="text/javascript" src="/js/room-detail.js"></script>
{{ end }}

{{ define "contents" }}

	<div class="container">
		{{ with .Room }}
		<div class="md-body md-room-detail clearfix">
			<div class="grid_9 md-main">
				<section class="heading-absolute box-sidebar align-center">
					<header class="box-heading heading-large">
						{{ firstChar .Title }}
						<div class="decription-override">
							<h2 class="h1">{{.Title}}</h2>
						</div>
					</header>
				</section>
				<div class="row">
					<!--  Slide Preview Room -->
					<div class="row slide-preview-room">
						<div id="sliderBig">
							<ul>
								<li class="slideshow-photo"><img src="http://placeholder.mac/1900x647.png" alt=""></li>
								<li class="slideshow-photo"><img src="http://placeholder.mac/1900x647.png" alt=""></li>
								<li class="slideshow-photo"><img src="http://placeholder.mac/1900x647.png" alt=""></li>
								<li class="slideshow-photo"><img src="http://placeholder.mac/1900x647.png" alt=""></li>
								<li class="slideshow-photo"><img src="http://placeholder.mac/1900x647.png" alt=""></li>
							</ul>
						</div>
						<div id="sliderThumb">
							<ul>
								<li slideIndex="0"><a href="#"><img src="http://placeholder.mac/190x110.png" alt=""></a></li>
								<li slideIndex="0"><a href="#"><img src="http://placeholder.mac/190x110.png" alt=""></a></li>
								<li slideIndex="0"><a href="#"><img src="http://placeholder.mac/190x110.png" alt=""></a></li>
								<li slideIndex="0"><a href="#"><img src="http://placeholder.mac/190x110.png" alt=""></a></li>
								<li slideIndex="0"><a href="#"><img src="http://placeholder.mac/190x110.png" alt=""></a></li>
							</ul>
						</div>
					</div>
					
					<div class="preview-content">
						{{.Details}}
					</div>

					<!-- Md booking -->
					<section class="md-booking">
						<h3 class="h3">Book Deluxe room only from <span class="number number-big">₹ <span class="one-night">/ Night</span></span></h3>
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
									<a href="#" class="btn btn-large btn-darkbrown">Book now</a>
								</div>
							</form>
						</div><!-- /.box-booking -->
					</section>
				</div>
			</div><!-- /.md-main -->
			<aside class="grid_3 md-sidebar md-sidebar-pt">
				<section class="box-sidebar">
					<h2 class="h3 header-sidebar">Room Facilities</h2>
					<ul class="list list-check">
						{{ range $facility := .Facilities }}
						<li>{{$facility}}</li>
						{{ end }}
					</ul>
				</section>
			</aside><!-- /.md-sidebar -->
		</div><!-- /.md-testimonail -->
		{{ end }}
	</div><!-- /.md-wrapper  -->
{{ end }}