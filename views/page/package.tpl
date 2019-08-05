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
										<p>Best Deals On Cheap Prices</p>
									</div>
								</header>
							</div>
							<div class="grid_9">
								<section class="md-booking">
									<h2 class="title-checkroom">Check availability</h2>
									<!--<div class="box-booking booking-inline">
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
									</div>--><!-- /.box-booking -->
								</section><!-- /.md-booking -->
							</div>
						</section>

						<section class="md-accomodation-content">
							<div class="row clearfix">
								<div class="grid_4">
									<article class="media">
										<figure>
											<img src="/images/package1.jpg" alt="" class="media-object">
										</figure>
										<section class="media-body">
											<h3 class="media-header h4"><a href="room-detail.html">Package 1</a></h3>
											<p class="media-content">Mauris egestas, tellus sed venenatis tincidunt, odio diam iaculis augue, nec tincidunt enim odio id arcu. Ut pellentesque, quam ut sagittis adipiscing, lectus metus mollis magna... </p>
											<div class="grid_12">
												<ul class="list list-check">
													<li>Maecenas posuere enim tellus, eu tur</li>
													<li>Leo consequat id. Donec eget nunc </li>
													<li>Morbi et mattis ipsum. Vivamus</li>
													<li>Nunc elementum semper mattis</li>
												</ul>
											</div>
											<a class="btn btn-small btn-border btn-border-brown" href="/reservation">Book Now</a>
										</section>
									</article>
								</div>
								<div class="grid_4">
									<article class="media">
										<figure>
											<img src="/images/package2.jpg" alt="" class="media-object">
										</figure>
										<section class="media-body">
											<h3 class="media-header h4"><a href="room-detail.html">Package 2</a></h3>
											<p class="media-content">Nunc at nunc in elit pellentesque dignissim at viverra nisl. Nam ac molestie sapien. Quisque posuere nisi non nisi feugiat mattis. Aenean vitae nisl a risus pellentesque pellentesque vel vulputate sem... </p>	
											<div class="grid_12">
												<ul class="list list-check">
													<li>Maecenas posuere enim tellus, eu tur</li>
													<li>Leo consequat id. Donec eget nunc </li>
													<li>Morbi et mattis ipsum. Vivamus</li>
													<li>Nunc elementum semper mattis</li>
												</ul>
											</div>
											<a class="btn btn-small btn-border btn-border-brown" href="/reservation">Book Now</a>
										</section>
									</article>
								</div>
								<div class="grid_4">
									<article class="media">
										<figure>
											<img src="/images/package3.jpg" alt="" class="media-object">
										</figure>
										<section class="media-body">
											<h3 class="media-header h4"><a href="room-detail.html">Package 3</a></h3>
											<p class="media-content">Mauris egestas, tellus sed venenatis tincidunt, odio diam iaculis augue, nec tincidunt enim odio id arcu. Ut pellentesque, quam ut sagittis adipiscing, lectus metus mollis magna... </p>
											<div class="grid_12">
												<ul class="list list-check">
													<li>Maecenas posuere enim tellus, eu tur</li>
													<li>Leo consequat id. Donec eget nunc </li>
													<li>Morbi et mattis ipsum. Vivamus</li>
													<li>Nunc elementum semper mattis</li>
												</ul>
											</div>
											<a class="btn btn-small btn-border btn-border-brown" href="/reservation">Book Now</a>
										</section>
									</article>
								</div>
							</div>
						</section>
					</div><!-- /.md-accomodation -->
				</section><!-- /.md-wrapper  -->
			</div>


{{ end }}