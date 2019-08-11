{{ template "layout/layout.tpl" . }}

{{ define "title" }} Activities - {{ end }}

{{ define "contents" }}

	<div class="md-body-wrapper">
		<div class="container">
			<div class="md-body md-blog clearfix">
				<div class="grid_9 md-main">
					<section class="heading-absolute box-sidebar align-center">
						<header class="box-heading heading-large">
							A
							<div class="decription-override">
								<h2 class="h1">Activities</h2>	
								<p> Adventures Activities</p>
							</div>
						</header>
					</section>

					<div class="list-article">
						{{ range $a := .Activities }}
						<article class="media">
							<figure class="pull-left">
								<img src="http://placeholder.mac/349x194.png" alt="" class="media-object">
							</figure>
							<div class="media-body">
								<header class="media-header">
									
									<h3 class="h4"><a href="blog-detail.html">{{$a.Title}}</a></h3>
								</header>
								<p class="media-content">{{$a.Body}} </p>
								
							</div>
						</article>
						{{ end }}

					</div>

				</div><!-- /.md-main -->
				<aside class="grid_3 md-sidebar md-sidebar-pt">

					<section class="box-sidebar">
						<h3 class="h3 header-sidebar">Rooms </h3>
						<ul class="list list-triangle list-uppercase">
							{{ range $room := .Rooms }}
							<li><a href="/room/{{$room.Slug}}">{{$room.Title}}</a></li>
							{{ end }}
						</ul>
					</section>

					<section class="box-sidebar">
						<h3 class="h3 header-sidebar">Packages</h3>
						<ul class="list list-triangle list-uppercase">
							{{ range $package := .Packages }}
							<li><a href="/packages">{{$package.Title}}</a></li>
							{{ end }}
						</ul>
					</section>

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
		</div>
	</div>

{{ end }}