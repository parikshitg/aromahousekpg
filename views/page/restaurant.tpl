{{ template "layout/layout.tpl" . }}

{{ define "title" }} Restaurant - {{ end }}



{{ define "contents" }}

	<div class="container">
		<div class="md-body md-room-detail clearfix">
			<div class="grid_9 md-main">
				<section class="heading-absolute box-sidebar align-center">
					<header class="box-heading heading-large">
						F
						<div class="decription-override">
							<h2 class="h1">Food &amp; Drinks</h2>
						</div>
					</header>
				</section>
				<div class="row">
					<section class="post-body">
						<div class="media">
							<header class="post-header">
								<h2 class="h2">Multi Cuisine Restaurant & Banquet Hall</h2>
							</header>
							{{ with .Restaurant }}
							<figure>
								<img src="{{ image .Image.URI 1900 647 }}" alt="" class="media-object">
							</figure>			
							<div class="media-body">
								<p>{{ htmlString .Richtext }}</p>
							</div>
							{{ end }}
						</div>
					</section>
				</div>

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
							<figure>
								<img src="{{ image $food.Image.URI 202 126 }}" alt="" class="media-object medium">
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
								<img src="{{ image $drink.Image.URI 202 126 }}" alt="" class="media-object medium">
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


			</div><!-- /.md-main -->
			<aside class="grid_3 md-sidebar md-sidebar-pt">
				<section class="box-sidebar">
					<h2 class="h3 header-sidebar"> Menu</h2>
					<ul class="list list-check">
						<li><a href="/restaurant">All Items</a></li>
						{{ range $k, $v := .Categories }}
						<li><a href="/restaurant/{{$k}}">{{$v}}</a></li>
						{{ end }}
					</ul>
				</section>
			</aside><!-- /.md-sidebar -->
		</div><!-- /.md-testimonail -->
	</div><!-- /.md-wrapper  -->
{{ end }}
