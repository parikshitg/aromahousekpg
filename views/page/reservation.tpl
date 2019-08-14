{{ template "layout/layout.tpl" . }}

{{ define "title" }} Reservation - {{ end }}

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

				<section class="box-sidebar">
				</section>
				
				<section class="box-sidebar">
					<h3 class="h3 header-sidebar">Packages</h3>
					<ul class="list list-triangle list-uppercase">
						{{ range $package := .Packages }}
						<li><a href="/packages">{{$package.Title}}</a></li>
						{{ end }}
					</ul>
				</section>
			</aside>


			<div class="grid_9">

				{{ with .ErrorMsg }}
				<div class="alert alert-error">
					<p class="alert-body">{{.}}</p>
				</div>
				{{ end }}

				{{ with .SuccessMsg }}
				<div class="alert alert-success">
					<p class="alert-body">{{.}}</p>
				</div>
				{{ end }}

				{{ with .GeneralMsg }}
				<div class="alert alert-general">
					<p class="alert-body">{{.}}</p>
				</div>
				{{ end }}

				<header class="box-heading">
					<h3 class="head headborder">Instant Booking</h3>
				</header>
				<br><br>
				<form  class="label-placeholder" action="/reservation" method="POST">
					<div class="row clearfix">
						<div class="form-group">
							<label>Name:</label>
							<input  type="text" name="name" class="input-control" placeholder="Name" 
								value="{{ with .Name }}{{.}}{{ end }}">
						</div>
						<div class="col-left">
							<label>Email:</label>
							<div class="form-group">
								<input  type="text" name="email" class="input-control" placeholder="Email" 
									value="{{ with .Email }}{{.}}{{ end }}">
							</div>
							<label>Arrival Date:</label>
							<div class="form-group">
								<div class="booking-form select-black">
									<label class="collapse input">
										<input type="text" name="arrival-date" id="arrival-date" class="input-control border-black"
											value="{{ with .ArrivalDate }}{{.}}{{ end }}" />
									</label>
								</div>
							</div>
							<div class="grid_6">
								<label>Adults:</label>
								<div class="form-group">
									<div class="input-group select-brown">
										<label class="collapse">
											{{ $adults := .Adults }}
											<select class="form-select" name="adults">
												<option></option>
												{{ range $i := numbers 1 10 }}
												{{ if eq $i $adults }}
												<option selected>{{$i}}</option>
												{{ else }}
												<option>{{$i}}</option>
												{{ end }}
												{{ end }}
											</select>
										</label>
									</div>
								</div>
							</div>
							<div class="grid_6">
								<label>Childs:</label>
								<div class="form-group">
									<div class="input-group select-brown">
										<label class="collapse">
											{{ $childs := .Childs }}
											<select class="form-select" name="childs">
												<option></option>
												{{ range $i := numbers 1 5 }}
												{{ if eq $i $childs }}
												<option selected>{{$i}}</option>
												{{ else }}
												<option>{{$i}}</option>
												{{ end }}
												{{ end }}
											</select>
										</label>
									</div>
								</div>
							</div>
						</div>
						<div class="col-right">
							<label>Phone No. :</label>
							<div class="form-group">
								<input  type="text" name="phone" class="input-control" placeholder="Phone no." 
									value="{{ with .Phone }}{{.}}{{ end }}">
							</div>
							<label>Departure Date : </label>
							<div class="form-group">
								<div class="booking-form select-black">
									<label class="collapse input">
										<input type="text" name="departure-date" id="departure-date" class="input-control border-black"
											value="{{ with .DepartureDate }}{{.}}{{ end }}"   />
									</label>
								</div>
							</div>
							<label>Room Type : </label>
							<div class="form-group">
								<div class="input-group select-brown">
									<label class="collapse">
										<select class="form-select" name="roomtype">
											<option></option>
											{{ if .RoomType }}
											<option {{if eq .RoomType "Standard Room"}}selected{{end}}>Standard Room</option>
											<option {{if eq .RoomType "Attic Room"}}selected{{end}}>Attic Room</option>
											<option {{if eq .RoomType "Suite Room"}}selected{{end}}>Suite Room</option>
											<option {{if eq .RoomType "Master Room"}}selected{{end}}>Master Room</option>
											<option {{if eq .RoomType "Junior Room"}}selected{{end}}>Junior Master</option>
											<option {{if eq .RoomType "Luxury Room"}}selected{{end}}>Luxury Room</option>
											{{ else }}
											<option>Standard Room</option>
											<option>Attic Room</option>
											<option>Suite Room</option>
											<option>Master Room</option>
											<option>Junior Master</option>
											<option>Luxury Room</option>
											{{ end }}
										</select>
									</label>
								</div>
							</div>
						</div>
					</div>
					<label>Your Message : </label>
					<div class="form-group">
						<textarea class="input-control" placeholder="Your message here" 
							name="message">{{ with .Message }}{{.}}{{ end }}</textarea>
					</div>
					
					<div class="form-group">
						<input type="hidden" name="form" value="reserve">
						<button type="submit" class="btn btn-alter btn-border btn-border-brown" >Submit</button>
					</div>
				</form>
			</div>
		</div>
	</div>

{{ end }}