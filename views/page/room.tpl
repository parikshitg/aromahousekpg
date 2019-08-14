{{ template "layout/layout.tpl" . }}

{{ define "title" }} Rooms - {{ end }}

{{ define "js" }}
<script type="text/javascript" src="/js/room-detail.js"></script>
{{ end }}

{{ define "contents" }}

	<div class="container">
		<div class="md-body md-room-detail clearfix">
			{{ with .Room }}
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
								<li class="slideshow-photo"><img src="{{ image "" 1900 647 }}" alt=""></li>
								<li class="slideshow-photo"><img src="{{ image "" 1900 647 }}" alt=""></li>
								<li class="slideshow-photo"><img src="{{ image "" 1900 647 }}" alt=""></li>
								<li class="slideshow-photo"><img src="{{ image "" 1900 647 }}" alt=""></li>
								<li class="slideshow-photo"><img src="{{ image "" 1900 647 }}" alt=""></li>
							</ul>
						</div>
						<div id="sliderThumb">
							<ul>
								<li slideIndex="0"><a href="#"><img src="{{ image "" 190 110 }}" alt=""></a></li>
								<li slideIndex="0"><a href="#"><img src="{{ image "" 190 110 }}" alt=""></a></li>
								<li slideIndex="0"><a href="#"><img src="{{ image "" 190 110 }}" alt=""></a></li>
								<li slideIndex="0"><a href="#"><img src="{{ image "" 190 110 }}" alt=""></a></li>
								<li slideIndex="0"><a href="#"><img src="{{ image "" 190 110 }}" alt=""></a></li>
							</ul>
						</div>
					</div>
					
					<div class="preview-content">
						<p>{{ htmlString .Details }}</p>
						<ul class="list list-check">
							{{ range $facility :=  .Facilities }}
							<li>{{ $facility }}</li>
							{{ end }}
						</ul>

					</div>

					<!-- Md booking -->
					<section class="md-booking">
						<h3 class="h">Book {{.Title}} from only <span class="number number-big">₹ {{ .EPCharges }}<span class="one-night">/ Night</span></span></h3>
						<br>
						<p>European Plan (Room only): ₹ {{.EPCharges}}</p>
						<p>Continental Plan (Breakfast included): ₹ {{.CPCharges}}</p>
						<p>American Plan (Breakfast + Lunch + Dinner): ₹ {{.APCharges}}</p>
						<p>Modified American Plan (Breakfast + Lunch/Dinner): ₹ {{.MAPCharges}}</p>
						<br>
						
						<strong>{{ if .Discount }}Discount: {{ .Discount }}{{ end }}</strong>

						<div class="box-booking booking-inline">
							<form method="post" action="/reservation">
								<div class="form-group">
									<label class="label-control">Arrival Date</label>
									<div class="booking-form select-black">
										<label class="collapse input">
											<input type="text" id="arrival-date" name="arrival-date" class="input-control border-black"/>
										</label>
									</div>
								</div>
								<div class="form-group">
									<label class="label-control">Departure Date</label>
									<div class="booking-form select-black">
										<label class="collapse input">
											<input type="text" id="departure-date" name="departure-date" class="input-control border-black"/>
										</label>
									</div>
								</div>
								<div class="form-group select">
									<label class="label-control">Adults</label>
									<div class="input-group select-black">
										<label class="collapse">
											<select class="form-select" name="adults">
												<option></option>
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
											<select class="form-select" name="childs">
												<option></option>
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
									<input type="hidden" name="roomtype" value="{{.RoomType}}">
									<input type="hidden" name="form" value="check">
									<button class="btn btn-large btn-darkbrown">Book now</button>
								</div>
							</form>
						</div>

					</section>
				</div>
			</div>
			{{ end }}

			<aside class="grid_3 md-sidebar md-sidebar-pt">
				<section class="box-sidebar">
					<h2 class="h3 header-sidebar">Room Types</h2>
					<ul class="list list-triangle">
						{{ range $room := .Rooms }}
						<li><a href="/room/{{$room.Slug}}"><h2 class="h3 header-sidebar">{{$room.Title}}</h2></a></li>
						{{ end }}
					</ul>
				</section>
			</aside>

		</div>
	</div>
{{ end }}