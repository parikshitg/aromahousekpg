<section class="md-booking">
	<h2 class="title-checkroom">Check Package Availability</h2>
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
				<input type="hidden" name="form" value="package">
				<input type="hidden" name="package" value="{{ with .Package}}{{.Title}}{{end}}">
				<button type="submit" class="btn btn-large btn-darkbrown">Check</button>
			</div>
		</form>
	</div><!-- /.box-booking -->
</section><!-- /.md-booking -->
