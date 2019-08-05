{{ template "layout/layout.tpl" . }}

{{ define "title"  }} Deluxe Room - {{ end }}

{{ define "description" }} 
About Aroma House 
{{ end }}

{{ define "keywords" }}
Restaurant, Hotel, Homestay, Lodge, Guest House.
{{ end }}

{{ define "contents" }}
<div class="container">
				<div class="md-body md-room-detail clearfix">
					<div class="grid_9 md-main">
						<section class="heading-absolute box-sidebar align-center">
							<header class="box-heading heading-large">
								S
								<div class="decription-override">
									<h2 class="h1">Standard room</h2>

								</div>
							</header>
						</section>
						<div class="row">
							<!--  Slide Preview Room -->
							<div class="row slide-preview-room">
								<div id="sliderBig">
									<ul>
										<li class="slideshow-photo"><img src="/images/deluxebig2.jpg" alt="" height="auto" width="100%"></li>
										<li class="slideshow-photo"><img src="/images/deluxebig2.jpg" alt="" height="auto" width="100%"></li>
										<li class="slideshow-photo"><img src="/images/deluxebig3.jpg" alt="" height="auto" width="100%"></li>
										<li class="slideshow-photo"><img src="/images/deluxebig4.jpg" alt="" height="auto" width="100%"></li>
									</ul>
								</div>
								<div id="sliderThumb">
									<ul>
										<li slideIndex="0"><a href="#"><img src="/images/deluxesmall1.jpg" alt=""></a></li>
										<li slideIndex="1"><a href="#"><img src="/images/deluxesmall2.jpg" alt=""></a></li>
										<li slideIndex="2"><a href="#"><img src="/images/deluxesmall3.jpg" alt=""></a></li>
										<li slideIndex="3"><a href="#"><img src="/images/deluxesmall4.jpg" alt=""></a></li>
									</ul>
								</div>
							</div>
							
							<div class="preview-content">
								<p>Well maintained and beautifully furnished bedroom. It is equipped with all modern amenities. The room is appx 12 feet by 12 feet with attached bathroom and western toilet. Peek out of the window for a serene view of the valley. The room has a queen size bed and a flat screen TV. Designed for couples but can accomodate maximumof 3 people with extra bed available. </p>
							</div>


							<!-- Md booking -->
							<section class="md-booking">
								<h3 class="h3">Book Standard room only from <span class="number number-big">₹ 1500<span class="one-night">/ Night</span></span></h3>
								<div class="box-booking booking-inline">
										<div class="form-group last">
											<label class="label-control"></label>
											<a href="/reservation" class="btn btn-large btn-darkbrown">Book now</a>
										</div>
								</div><!-- /.box-booking -->
							</section>
						</div>
					</div><!-- /.md-main -->
					<aside class="grid_3 md-sidebar md-sidebar-pt">
						<section class="box-sidebar">
							<h2 class="h3 header-sidebar">Room Facilities</h2>
							<ul class="list list-check">
								<li>Free WiFi access</li>
								<li>Concierge</li>
								<li>Laundry service/dry cleaning</li>
								<li>Attached Bathroom</li>
								<li>Pets allowed</li>
								<li>Heater</li>
								<li>Geyser</li>
								<li>Telivision</li>	
							</ul>
						</section>
					</aside><!-- /.md-sidebar -->
				</div><!-- /.md-testimonail -->
			</div><!-- /.md-wrapper  -->

{{ end }}

{{ define "js" }}
<script type="text/javascript" src="/js/room-detail.js"></script>
{{ end }}