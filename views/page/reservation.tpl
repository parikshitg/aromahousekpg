{{ template "layout/layout.tpl" . }}

{{ define "title" }} Reservation - {{ end }}

{{ define "description" }} 
About Aroma House 
{{ end }}

{{ define "keywords" }}
Restaurant, Hotel, Homestay, Lodge, Guest House.
{{ end }}

{{ define "contents" }}
<div class="container">
				<div class="md-body md-checkout clearfix">

					<aside class="grid_3 md-sidebar ">
						<header class="box-heading heading-large">
							R
							<div class="decription-override">
								<h2 class="h1">Reservation</h2>
								<p>Room Reservation </p> 
							</div>
						</header>
						<section class="box-sidebar ">
								<h3 class="h3 header-sidebar">Packages</h3>
								<ul class="list-cate list list-triangle list-uppercase">
									<li><a href="#">Cras condimentum</a></li>
									<li><a href="#">Nibh et pellentesque</a></li>
									<li><a href="#">Leo nibh gravida velit</a></li>
									<li><a href="#">Tortor augue a eros</a></li>
									<li><a href="#">libero</a></li>
								</ul>
							</section>
					</aside>
					<div class="grid_9">
						<header class="box-heading">
							<h3 class="head headborder">Instant Booking</h3>
						</header>
						<br><br>
						<form  class="label-placeholder" action="/reservation" method="POST">
							<div class="row clearfix">
								<div class="form-group">
									<label>Name:</label>
									<input  type="text" name="name" class="input-control" placeholder="Name">
								</div>
								<div class="col-left">
									<label>Email:</label>
									<div class="form-group">
										<input  type="text" name="email" class="input-control" placeholder="Email">
									</div>
									<label>Arrival Date:</label>
									<div class="form-group">
										<div class="booking-form select-black">
											<label class="collapse input">
												<input type="text" name="arrival-date" id="arrival-date" class="input-control border-black"  />
											</label>
										</div>
									</div>
									<div class="grid_6">
										<label>Adults:</label>
										<div class="form-group">
											<div class="input-group select-brown">
												<label class="collapse">
													<select class="form-select" name="adults" >
														<option></option>
														<option>1</option>
														<option>2</option>
														<option>3</option>
														<option>4</option>
														<option>5</option>
														<option>6</option>
														<option>7</option>
														<option>8</option>
														<option>9</option>
														<option>10</option>
													</select>
												</label>
											</div>
										</div>
									</div>
									<div class="grid_6">
										<label>Kids:</label>
										<div class="form-group">
											<div class="input-group select-brown">
												<label class="collapse">
													<select class="form-select" name="kids">
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
									</div>
								</div>
								<div class="col-right">
									<label>Phone No. :</label>
									<div class="form-group">
										<input  type="text" name="phone" class="input-control" placeholder="Phone no.">
									</div>
									<label>Departure Date : </label>
									<div class="form-group">
										<div class="booking-form select-black">
											<label class="collapse input">
												<input type="text" name="departure-date" id="departure-date" class="input-control border-black"  />
											</label>
										</div>
									</div>
									<label>Room Type : </label>
									<div class="form-group">
										<div class="input-group select-brown">
											<label class="collapse">
												<select class="form-select" name="roomtype">
													<option></option>
													<option>Standard Room</option>
													<option>Attic Room</option>
													<option>Suite Room</option>
													<option>Master Room</option>												
													<option>Junior Master</option>
													<option>Luxury</option>
												</select>
											</label>
										</div>
									</div>
								</div>
								</div>
								<label>Your Message : </label>
								<div class="form-group">
									<textarea    class="input-control" placeholder="Your message here" name="message"></textarea>
								</div>
								
								<div class="form-group">
									<button type="submit" class="btn btn-alter btn-border btn-border-brown" >Submit</button>
							</div>
							{{ if .Reservation }}
									{{ .Reservation }}
								{{ end }}
						</form>
					</div>
				</div>
			</div>


{{ end }}