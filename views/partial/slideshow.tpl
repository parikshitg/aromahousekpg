	<!-- Slider with Jquery bxslider -->
	<section class="md-slide-wrapper md-slide-home">
		<!-- Slide Image-->
		<ul id="md-slide-fade" class="md-slide clearfix">

			{{ range $banner := .Banners }}
			<li>
				<!-- img src="http://placehold.it/1615x647" alt="" -->
				<img src="{{ image $banner.Image.URI 1615 647 }}" alt="" >
				<div class="content-slide">
					<div class="home-content">
						<h2 class="slide-title alter">{{$banner.Title}}</h2>
						<h3 class="slide-title alter">{{$banner.Subtitle}}</h3>
						<div class="dots-line">
							<span></span>
							<span></span>
							<span></span>
						</div>
						<p class="slide-description">{{$banner.Text}}</p>
						<a href="{{$banner.Link}}" class="btn btn-border btn-border-white">see more</a>
					</div>
				</div>
			</li>
			{{ end }}

		</ul>
	</section>
	<!-- Slider End -->