{{ template "layout/layout.tpl" . }}

{{ define "title" }} Restaurant - {{ end }}

{{ define "description" }} 
About Aroma House 
{{ end }}

{{ define "keywords" }}
Restaurant, Hotel, Homestay, Lodge, Guest House.
{{ end }}

{{ define "contents" }}
	<div class="container">
		<div class="md-body md-standar-post clearfix">
			<div class="grid_12 md-main">
				<!--<section class="heading-absolute box-sidebar align-center">
					<header class="box-heading heading-large">
						R
						<div class="decription-override">
							<h2 class="h1">Restaurant</h2>
							<p>Flavours of Authentic Cuisine</p>
						</div>
					</header>
				</section>-->
				<div class="row">
					<section class="post-body">
						<div class="media">
							<header class="post-header">
								<h2 class="h2">Multi Cuisine Restaurant & Banquet Hall</h2>
							</header>
							{{ with .Restaurant }}
							<figure>
								<img src="http://placeholder.mac/1900x647.png" alt="" class="media-object">
							</figure>			
							<div class="media-body">
								{{ htmlString .Richtext }}
							</div>
							{{ end }}
						</div>
					</section>
				</div>
			</div><!-- /.md-main -->
		</div><!-- /.md-testimonail -->

		<div class="row row-alter-110 clearfix">
			{{ range $cuisine, $submenu := .FoodMenu }}
			<article class="box box-bullet-list">
				<header class="grid_12 box-heading">
					<h3 class="head headline">{{$cuisine}}</h3>
				</header>

				<br><br><br>

				<div class="box-body clearfix">
					{{ range $food := $submenu}}
					<div class="grid_3">
						<div class="col media align-center">
							<br>
							<br>
							<figure>
								<img src="http://placeholder.mac/202x126.png" alt="" class="media-object medium">
							</figure>
							<br>
							<header class="box-heading">
								<h5 class="h5">{{$food.Title}}</h5>
							</header>
							
							<span class="number number-big">₹{{$food.Price}}<span class="one-night"></span></span>
						</div><!-- /.media -->
					</div>
					{{ end }}

				</div>
				<br><br><br>
			</article>
			{{ end }}

		</div><!-- /.row -->


		<div class="row row-alter-110 clearfix">
			{{ range $category, $submenu := .DrinkMenu }}
			<article class="box box-bullet-list">
				<header class="grid_12 box-heading">
					<h3 class="head headline">{{$category}}</h3>
				</header>

				<br><br><br>

				<div class="box-body clearfix">
					{{ range $drink := $submenu}}
					<div class="grid_3">
						<div class="col media align-center">
							<br>
							<br>
							<figure>
								<img src="http://placeholder.mac/202x126.png" alt="" class="media-object medium">
							</figure>
							<br>
							<header class="box-heading">
								<h5 class="h5">{{$drink.Title}}</h5>
							</header>
							
							<span class="number number-big">₹{{$drink.Price}}<span class="one-night"></span></span>
						</div><!-- /.media -->
					</div>
					{{ end }}

				</div>
				<br><br><br>
			</article>
			{{ end }}

		</div><!-- /.row -->

	</div><!-- /.md-wrapper  -->
{{ end }}