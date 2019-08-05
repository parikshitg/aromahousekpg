{{ template "layout/layout.tpl" . }}

{{ define "description" }} 
About Aroma House 
{{ end }}

{{ define "keywords" }}
Restaurant, Hotel, Homestay, Lodge, Guest House.
{{ end }}

{{ define "contents" }}
<!-- Slider with Jquery bxslider -->
			<section class="md-slide-wrapper md-slide-home">
				<!-- Slide Image-->
				<ul id="md-slide-fade" class="md-slide clearfix">
					<li>
						<img src="/images/slide1.jpg" alt="" >
						<div class="content-slide">
							<div class="home-content">
								<h2 class="slide-title">Aroma House</h2>
								<h3 class="slide-title">A welcome treat in kalimpong</h3>
								<div class="dots-line">
									<span></span>
									<span></span>
									<span></span>
								</div>
								<p class="slide-description">Vestibulum facilisis eros at mattis aliquam. Aliquam lacinia nulla eget eros volutpat dictum. Ut venenatis ante mauris, quis vestibulum felis.</p>
								<a href="/about" class="btn btn-border btn-border-white">see more</a>
							</div>
						</div>
					</li>
					<li>
						<img src="/images/g11.jpg" alt="" >
						<div class="content-slide">
							<div class="home-content">
								<h2 class="slide-title">Our Restaurant</h2>
								<h3 class="slide-title">Taste the difference</h3>
								<div class="dots-line">
									<span></span>
									<span></span>
									<span></span>
								</div>
								<p class="slide-description">Vestibulum facilisis eros at mattis aliquam. Aliquam lacinia nulla eget eros volutpat dictum. Ut venenatis ante mauris, quis vestibulum felis.</p>
								<a href="/restaurant" class="btn btn-border btn-border-white">see more</a>
							</div>
						</div>
					</li>
					<li>
						<img src="/images/kanchenjunga.jpg" alt="">
						<div class="content-slide">
							<div class="home-content">
								<h2 class="slide-title alter">Aenean Turpis </h2>
								<h3 class="slide-title alter">Malesuada Quis Ante</h3>
								<p class="slide-description">Morbi consectetur, sem sit amet mollis fermentum, lectus risus euismod diam, vitae sollicitudin purus eros vel dolor. .</p>
								<a href="#" class="btn btn-border btn-border-white">see more</a>
							</div>
						</div>
					</li>
					<li>
						<img src="/images/g12.jpg" alt="" >
						<div class="content-slide ">
							<div class="home-content">
								<h3 class="slide-title">Enjoy A luxury Experience</h3>
								<p class="slide-description">Morbi consectetur, sem sit amet mollis fermentum, lectus risus euismod diam, vitae sollicitudin purus eros vel dolor. .</p>
								<a href="/rooms" class="btn btn-border btn-border-white">see more</a>
							</div>
						</div>
					</li>
				</ul>
			</section>
			<!-- Slider End -->

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
							<aside class="grid_3 md-sidebar">
								<section class="widget-home-info">
									<h3 class="h3 header-sidebar">why Aroma House?</h3>
									<div class="home-content">
										<p>Aroma House is a brand new homestay in the scenic eco village located 5 kms from the main town of Kalimpong.The Aroma House is owned & run by master chef Bikas Lama. We strive to give you clean, comfortable, and affordable accommodation with spectacular views in the most idyllic and tranquil area of Kalimpong. </p>
										<p>Having worked within the hotel industry for over 25 years running and organizing restuarants Chef Lama have built Aroma House to satisfy your every need, with a friendly, cosy and relaxing atmosphere creating a home from home experience.</p>
										<ul class="list list-check">
											<li>Clean and Tidy Rooms</li>
											<li>Multicuisine Restuarant</li>
											<li>24 Hour Room Service</li>											
											<li>Free Parking & Security</li>
											<li>Free WiFi</li>
											<li>Laundry Service</li>
											<li>Wedding Venue</li>
											<li>Hot Showers</li>
											<li>Air Conditioning</li>
											<li>Intercom Facility</li>
											<li>Special Packages for devoted travellers</li>
										</ul>
									</div>
								</section>
							</aside><!-- /.md-sidebar -->

							<div class="grid_9 md-primary">
								<div class="row row-home">
									<article class="media">
										<figure class="media-image">
											<img src="/images/deluxroom.jpeg" alt="" class="media-object">
										</figure>
										<section class="media-body">
											<h3 class="media-header media-header-big">
												<a href="/deluxe">Standard Room</a>
											</h3>
											<p class="media-content">Mauris egestas, tellus sed venenatis tincidunt, odio diam iaculis augue, nec tincidunt enim odio id arcu. Ut pellentesque, quam ut sagittis adipiscing, lectus metus mollis magna...<a href="/deluxe" class="text-link link-direct">see more</a></p>
										</section>
									</article>
								</div><!-- /.row-article -->
								<div class="row row-home">
									<article class="media">
										<figure class="media-image">
											<img src="/images/elegantroom.jpg" alt="" class="media-object">
										</figure>
										<section class="media-body">
											<h3 class="media-header media-header-big"><a href="/elegant">Master Room</a></h3>
											<p class="media-content">Nunc at nunc in elit pellentesque dignissim at viverra nisl. Nam ac molestie sapien. Quisque posuere nisi non nisi feugiat mattis. Aenean vitae nisl a risus pellentesque pellentesque vel vulputate sem... <a href="/elegant" class="text-link link-direct">see more</a></p>
										</section>
									</article>
								</div><!-- /.row-article -->
								<div class="row row-home">
									<article class="media">
										<figure class="media-image">
											<img src="/images/woodenroom.jpg" alt="" class="media-object">
										</figure>
										<section class="media-body">
											<h3 class="media-header media-header-big"><a href="wooden">Junior Master Room</a></h3>
											<p class="media-content">Mauris egestas, tellus sed venenatis tincidunt, odio diam iaculis augue, nec tincidunt enim odio id arcu. Ut pellentesque, quam ut sagittis adipiscing, lectus metus mollis magna...<a href="/wooden" class="text-link link-direct">see more</a></p>
										</section>
									</article>
								</div><!-- /.row-article -->
							</div><!-- /.md-sidebar -->
						</div>
					</section><!-- /.md-home-body -->
				</div>
			</div>


{{ end }}